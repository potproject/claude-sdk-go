package v1

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/tmaxmax/go-sse"
)

// Constants for SSE event types
const (
	MessagesStreamResponseTypeMessageStart      = "message_start"
	MessagesStreamResponseTypeContentBlockStart = "content_block_start"
	MessagesStreamResponseTypePing              = "ping"
	MessagesStreamResponseTypeContentBlockDelta = "content_block_delta"
	MessagesStreamResponseTypeContentBlockStop  = "content_block_stop"
	MessagesStreamResponseTypeMessageDelta      = "message_delta"
	MessagesStreamResponseTypeMessageStop       = "message_stop"
	MessagesStreamResponseTypeError             = "error"
)

// CreateMessagesStream struct holds the connection and event channels for streaming
type CreateMessagesStream struct {
	Connection                 *sse.Connection
	Unsubscribe                func()
	Event                      chan sse.Event
	Error                      chan error
	ResponseBodyMessagesStream ResponseBodyMessagesStream
}

// ResponseBodyMessagesStream struct holds the response data for streaming messages
type ResponseBodyMessagesStream struct {
	Id           string                              `json:"id"`
	Type         string                              `json:"type"` // always "message"
	Role         string                              `json:"role"` // always "assistant"
	Content      []ResponseBodyMessagesContentStream `json:"content"`
	Model        string                              `json:"model"`
	StopReason   string                              `json:"stop_reason"` // "end_turn" or "max_tokens", "stop_sequence", null
	StopSequence string                              `json:"stop_sequence"`
	Usage        struct {
		InputTokens  int64 `json:"input_tokens"`
		OutputTokens int64 `json:"output_tokens"`
	} `json:"usage"`
}

// ResponseBodyMessagesContentStream struct holds the content of the streaming messages
type ResponseBodyMessagesContentStream struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

// ResponseContentMessageStartStream struct holds the initial message start event data
type ResponseContentMessageStartStream struct {
	Type    string                     `json:"type"`
	Message ResponseBodyMessagesStream `json:"message"`
}

// ResponseContentBlockDeltaStream struct holds the content block delta event data
type ResponseContentBlockDeltaStream struct {
	Type  string `json:"type"`
	Index int64  `json:"index"`
	Delta struct {
		Type string `json:"type"`
		Text string `json:"text"`
	} `json:"delta"`
}

// ResponseMessageDeltaStream struct holds the message delta event data
type ResponseMessageDeltaStream struct {
	Type  string `json:"type"`
	Delta struct {
		StopReason   string `json:"stop_reason"`
		StopSequence string `json:"stop_sequence"`
	} `json:"delta"`
	Usage struct {
		OutputTokens int64 `json:"output_tokens"`
	} `json:"usage"`
}

// CreateMessagesStream initializes a new streaming connection for messages
func (c *Client) CreateMessagesStream(ctx context.Context, body RequestBodyMessages) (*CreateMessagesStream, error) {
	reqURL := c.config.BaseURL + c.config.Endpoint
	body.Stream = true

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	reqHeaders := c.defaultHeaders()
	req, err := http.NewRequestWithContext(ctx, "POST", reqURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	for k, v := range reqHeaders {
		req.Header.Set(k, v)
	}

	client := sse.Client{
		HTTPClient: c.config.HTTPClient,
		Backoff: sse.Backoff{
			MaxRetries: -1,
		},
	}

	conn := client.NewConnection(req)
	chanEvent := make(chan sse.Event)
	connectionError := make(chan error)

	unsubscribe := conn.SubscribeToAll(func(e sse.Event) {
		if e.Type == MessagesStreamResponseTypePing ||
			e.Type == MessagesStreamResponseTypeContentBlockStart ||
			e.Type == MessagesStreamResponseTypeContentBlockStop {
			return
		}
		chanEvent <- e
	})

	go func() {
		err := conn.Connect()
		if !errors.Is(err, io.EOF) && err != nil {
			connectionError <- err
		}
	}()

	return &CreateMessagesStream{
		Connection:                 conn,
		Unsubscribe:                unsubscribe,
		Event:                      chanEvent,
		Error:                      connectionError,
		ResponseBodyMessagesStream: ResponseBodyMessagesStream{},
	}, nil
}

// Close closes the streaming connection and cleans up resources
func (c *CreateMessagesStream) Close() {
	close(c.Event)
	close(c.Error)
	c.Unsubscribe()
}

// Recv receives events from the streaming connection and updates usage data
func (c *CreateMessagesStream) Recv() (ResponseBodyMessagesStream, error) {
	select {
	case e := <-c.Event:
		switch e.Type {
		case MessagesStreamResponseTypeMessageStart:
			d := []byte(e.Data)
			var r ResponseContentMessageStartStream
			err := json.Unmarshal(d, &r)
			if err != nil {
				return ResponseBodyMessagesStream{}, err
			}
			c.ResponseBodyMessagesStream = r.Message
			c.ResponseBodyMessagesStream.Content = []ResponseBodyMessagesContentStream{
				{
					Type: "text",
					Text: "",
				},
			}
			return c.ResponseBodyMessagesStream, nil
		case MessagesStreamResponseTypeContentBlockDelta:
			d := []byte(e.Data)
			var r ResponseContentBlockDeltaStream
			err := json.Unmarshal(d, &r)
			if err != nil {
				return ResponseBodyMessagesStream{}, err
			}
			c.ResponseBodyMessagesStream.Content = append(c.ResponseBodyMessagesStream.Content, ResponseBodyMessagesContentStream{
				Type: "text",
				Text: r.Delta.Text,
			})
			// Update output tokens usage
			c.ResponseBodyMessagesStream.Usage.OutputTokens += int64(len(r.Delta.Text)) // Simplified token count, adjust as needed
			return c.ResponseBodyMessagesStream, nil
		case MessagesStreamResponseTypeMessageDelta:
			d := []byte(e.Data)
			var r ResponseMessageDeltaStream
			err := json.Unmarshal(d, &r)
			if err != nil {
				return ResponseBodyMessagesStream{}, err
			}
			c.ResponseBodyMessagesStream.StopReason = r.Delta.StopReason
			c.ResponseBodyMessagesStream.StopSequence = r.Delta.StopSequence
			c.ResponseBodyMessagesStream.Usage.OutputTokens += r.Usage.OutputTokens
			return c.ResponseBodyMessagesStream, nil
		case MessagesStreamResponseTypeMessageStop:
			return c.ResponseBodyMessagesStream, io.EOF
		case MessagesStreamResponseTypeError:
			d := []byte(e.Data)
			var r ResponseError
			err := json.Unmarshal(d, &r)
			if err != nil {
				return ResponseBodyMessagesStream{}, err
			}
			return c.ResponseBodyMessagesStream, errors.New(r.Error.Message)
		}
	case err := <-c.Error:
		return ResponseBodyMessagesStream{}, err
	}
	return c.ResponseBodyMessagesStream, nil
}

// Helper function to set default headers
func (c *Client) defaultHeaders() map[string]string {
	headers := map[string]string{
		"X-Api-Key":         c.config.ApiKey,
		"Anthropic-Version": c.config.Version,
		"Content-Type":      "application/json",
	}
	if c.config.Beta != "" {
		headers["Anthropic-Beta"] = c.config.Beta
	}
	return headers
}
