package v1

type RequestBodyMessages struct {
	Model          string                        `json:"model"`
	Messages       []RequestBodyMessagesMessages `json:"messages"`
	System         string                        `json:"-"`
	SystemTypeText []RequestBodySystemTypeText   `json:"-"`
	SystemRaw      interface{}                   `json:"system,omitempty"` // optional
	MaxTokens      int                           `json:"max_tokens"`
	Thinking       *RequestBodyMessagesThinking  `json:"thinking,omitempty"`    // optional
	MetaData       map[string]interface{}        `json:"metadata"`              // optional
	StopSequences  []string                      `json:"stop_sequences"`        // optional
	Stream         bool                          `json:"stream"`                // optional
	Temperature    float64                       `json:"temperature,omitempty"` // optional
	TopP           float64                       `json:"top_p,omitempty"`       // optional
	TopK           float64                       `json:"top_k,omitempty"`       // optional
}

type RequestBodyMessagesThinking struct {
	Type         string `json:"type"`
	BudgetTokens int    `json:"budget_tokens"`
}

type RequestBodyMessagesMessages struct {
	Role             string                                        `json:"role"`
	ContentRaw       interface{}                                   `json:"content"`
	Content          string                                        `json:"-"`
	ContentTypeText  []RequestBodyMessagesMessagesContentTypeText  `json:"-"`
	ContentTypeImage []RequestBodyMessagesMessagesContentTypeImage `json:"-"`
}

type RequestBodySystemTypeText struct {
	Type         string               `json:"type"` // always "text"
	Text         string               `json:"text"`
	CacheControl *RequestCacheControl `json:"cache_control"`
}

type RequestCacheControl struct {
	Type string `json:"type"` // always "ephemeral"
}

const (
	RequestBodyMessagesMessagesContentTypeTextType        = "text"
	RequestBodyMessagesMessagesContentTypeImageType       = "image"
	RequestBodyMessagesMessagesContentTypeImageSourceData = "data"
	RequestBodyMessagesMessagesContentTypeImageSourceUrl  = "url"
)

type RequestBodyMessagesMessagesContentTypeText struct {
	Type         string               `json:"type"` // always "text"
	Text         string               `json:"text"`
	CacheControl *RequestCacheControl `json:"cache_control"` // optional
}

type RequestBodyMessagesMessagesContentTypeImage struct {
	Type         string                                            `json:"type"` // always "image"
	Source       RequestBodyMessagesMessagesContentTypeImageSource `json:"source"`
	CacheControl *RequestCacheControl                              `json:"cache_control"` // optional
}

type RequestBodyMessagesMessagesContentTypeImageSource struct {
	Type      string `json:"type"`                 // "base64" or "url"
	MediaType string `json:"media_type,omitempty"` // base64 type required
	Data      string `json:"data,omitempty"`       // base64 type required
	Url       string `json:"url,omitempty"`        // url type required
}

const (
	MessagesRoleUser      = "user"
	MessagesRoleAssistant = "assistant"
)
