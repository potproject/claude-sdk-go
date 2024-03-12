package v1

type RequestBodyMessages struct {
	Model         string                        `json:"model"`
	Messages      []RequestBodyMessagesMessages `json:"messages"`
	System        string                        `json:"system"` // optional
	MaxTokens     int                           `json:"max_tokens"`
	MetaData      map[string]interface{}        `json:"metadata"`       // optional
	StopSequences []string                      `json:"stop_sequences"` // optional
	Stream        bool                          `json:"stream"`         // optional
	Temperature   float64                       `json:"temperature"`    // optional
	TopP          float64                       `json:"top_p"`          // optional
	TopK          float64                       `json:"top_k"`          // optional
}

type RequestBodyMessagesMessages struct {
	Role             string                                        `json:"role"`
	ContentRaw       interface{}                                   `json:"content"`
	Content          string                                        `json:"-"`
	ContentTypeText  []RequestBodyMessagesMessagesContentTypeText  `json:"-"`
	ContentTypeImage []RequestBodyMessagesMessagesContentTypeImage `json:"-"`
}

const (
	RequestBodyMessagesMessagesContentTypeTextType  = "text"
	RequestBodyMessagesMessagesContentTypeImageType = "image"
)

type RequestBodyMessagesMessagesContentTypeText struct {
	Type string `json:"type"` // always "text"
	Text string `json:"text"`
}

type RequestBodyMessagesMessagesContentTypeImage struct {
	Type   string                                            `json:"type"` // always "image"
	Source RequestBodyMessagesMessagesContentTypeImageSource `json:"source"`
}

type RequestBodyMessagesMessagesContentTypeImageSource struct {
	Type      string `json:"type"`
	MediaType string `json:"media_type"`
	Data      string `json:"data"`
}

const (
	MessagesRoleUser      = "user"
	MessagesRoleAssistant = "assistant"
)
