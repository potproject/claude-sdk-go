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
		SystemTypeText: []claude.RequestBodySystemTypeText{
			claude.UseSystemCacheEphemeral("Please speak in Japanese."),
		},
		Messages: []claude.RequestBodyMessagesMessages{
			{
				Role: claude.MessagesRoleUser,
				ContentTypeText: []claude.RequestBodyMessagesMessagesContentTypeText{
					{
						Text:         "Hello!",
						CacheControl: claude.UseCacheEphemeral(),
					},
				},
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
	// こんにちは！日本語でお話しましょう。
}
