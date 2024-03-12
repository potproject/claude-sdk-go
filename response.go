package main

type ResponseBodyMessages struct {
	Id           string                        `json:"id"`
	Type         string                        `json:"type"` // always "message"
	Role         string                        `json:"role"` // always "assistant"
	Content      []ResponseBodyMessagesContent `json:"content"`
	Model        string                        `json:"model"`
	StopReason   string                        `json:"stop_reason"` // "end_turn" or "max_tokens", "stop_sequence", null
	StopSequence string                        `json:"stop_sequence"`
	Usage        struct {
		InputTokens  int64 `json:"input_tokens"`
		OutputTokens int64 `json:"output_tokens"`
	} `json:"usage"`
}

type ResponseBodyMessagesContent struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type ResponseError struct {
	Error struct {
		Message string `json:"message"`
	} `json:"error"`
}
