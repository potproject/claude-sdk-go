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
	source, err := claude.TypeImageSourceLoadFile("image.png")
	if err != nil {
		panic(err)
	}
	m := claude.RequestBodyMessages{
		Model:     "claude-3-opus-20240229",
		MaxTokens: 1024,
		Messages: []claude.RequestBodyMessagesMessages{
			{
				Role: claude.MessagesRoleUser,
				ContentTypeImage: []claude.RequestBodyMessagesMessagesContentTypeImage{
					{
						Source: source,
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
	// The image shows a smooth gradient transitioning from a bright, vibrant orange at the bottom to a light blue at the top. The colors blend seamlessly into each other, creating a visually striking and aesthetically pleasing effect. The simplicity of the gradient allows the colors to be the main focus, showcasing their luminosity and the way they interact with one another. This type of gradient is often used as a background or design element to add depth, warmth, and visual interest to various digital or print media projects.
}
