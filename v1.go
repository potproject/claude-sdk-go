package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	config ClientConfig
}

const (
	defaultBaseURL  = "https://api.anthropic.com/"
	defaultEndpoint = "v1/messages"
	defaultVersion  = "2023-06-01"
	defaultBeta     = ""
)

const contentType = "application/json"

type ClientConfig struct {
	ApiKey string

	Version string
	Beta    string // Using Beta API : anthropic-beta Header

	BaseURL    string
	Endpoint   string
	HTTPClient *http.Client
}

func defaultConfig(apiKey string) ClientConfig {
	return ClientConfig{
		ApiKey: apiKey,

		Version: defaultVersion,
		Beta:    "",

		BaseURL:    defaultBaseURL,
		Endpoint:   defaultEndpoint,
		HTTPClient: &http.Client{},
	}
}

func NewClient(apiKey string) *Client {
	return &Client{
		config: defaultConfig(apiKey),
	}
}

func NewClientWithConfig(config ClientConfig) *Client {
	return &Client{
		config: config,
	}
}

func (c *Client) SetVersion(version string) {
	c.config.Version = version
}

func (c *Client) CreateMessages(body RequestBodyMessages) (*ResponseBodyMessages, error) {
	reqURL := c.config.BaseURL + c.config.Endpoint
	reqHeaders := map[string]string{
		"X-Api-Key":         c.config.ApiKey,
		"Anthropic-Version": c.config.Version,
		"Content-Type":      contentType,
	}
	if c.config.Beta != "" {
		reqHeaders["anthropic-beta"] = c.config.Beta
	}

	jsonBody, err := parseBodyJSON(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", reqURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	for k, v := range reqHeaders {
		req.Header.Set(k, v)
	}

	resp, err := c.config.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var result ResponseBodyMessages
		err = json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return &result, nil
	}
	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		var result ResponseError
		err = json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf(result.Error.Message)
	}
	return nil, fmt.Errorf("unexpected error: %d", resp.StatusCode)
}

func parseBodyJSON(req RequestBodyMessages) ([]byte, error) {
	for i, m := range req.Messages {
		if m.Content != "" {
			req.Messages[i].ContentRaw = m.Content
		}

		if len(m.ContentTypeText) > 0 {
			for j := range m.ContentTypeText {
				m.ContentTypeText[j].Type = "text"
			}
			raw, err := json.Marshal(m.ContentTypeText)
			if err != nil {
				return nil, err
			}
			req.Messages[i].ContentRaw = json.RawMessage(raw)
		}

		if len(m.ContentTypeImage) > 0 {
			for j := range m.ContentTypeImage {
				m.ContentTypeImage[j].Type = "image"
			}
			raw, err := json.Marshal(m.ContentTypeImage)
			if err != nil {
				return nil, err
			}
			req.Messages[i].ContentRaw = json.RawMessage(raw)
		}
	}
	return json.Marshal(req)
}
