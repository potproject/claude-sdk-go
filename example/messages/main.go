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
		MaxTokens: 1024,
		Messages: []claude.RequestBodyMessagesMessages{
			{
				Role:    claude.MessagesRoleUser,
				Content: "Hello, world!",
				// Alternatively, you can use ContentTypeText
				//
				// ContentTypeText: []claude.RequestBodyMessagesMessagesContentTypeText{
				// 	{
				// 		Text: "Hello, world!",
				// 	},
				// },
			},
		},
	}
	ctx := context.Background()
	res, err := c.CreateMessages(ctx, m)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Content[0].Text)
	// Output:
	// Hello! How can I assist you today?
}
