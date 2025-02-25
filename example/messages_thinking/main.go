package main

import (
	"context"
	"fmt"
	"os"

	claude "github.com/potproject/claude-sdk-go"
)

func main() {
	apiKey := os.Getenv("API_KEY")
	c := claude.NewClient(apiKey)
	m := claude.RequestBodyMessages{
		Model:     "claude-3-7-sonnet-20250219",
		MaxTokens: 8192,
		Thinking:  claude.UseThinking(4096),
		Messages: []claude.RequestBodyMessagesMessages{
			{
				Role:    claude.MessagesRoleUser,
				Content: "Hello, world!",
			},
		},
	}
	ctx := context.Background()
	res, err := c.CreateMessages(ctx, m)
	if err != nil {
		panic(err)
	}

	// Output:
	// [thinking] This is a simple "Hello, world!" greeting from the user. It's a common first phrase in programming and also a standard greeting in conversations with AI assistants. I should respond in a friendly and welcoming manner.
	// [text] Hi there! It's nice to meet you. "Hello, world!" is such a classic greeting - it brings back memories of first programming lessons for many! How are you doing today? Is there something specific I can help you with?
	for _, v := range res.Content {
		if v.Type == claude.ResponseBodyMessagesContentTypeThinking {
			fmt.Println("[thinking]", v.Thinking)
		}
		if v.Type == claude.ResponseBodyMessagesContentTypeText {
			fmt.Println("[text]", v.Text)
		}
	}
}
