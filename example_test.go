package v1_test

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	claude "github.com/potproject/claude-sdk-go"
)

func Example() {
	apiKey := os.Getenv("API_KEY")
	c := claude.NewClient(apiKey)
	m := claude.RequestBodyMessages{
		Model:     "claude-3-opus-20240229",
		MaxTokens: 64,
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
	fmt.Println(res.Content[0].Text)
	// Output:
	// Hello! How can I assist you today?
}

func Example_typeImage() {
	apiKey := os.Getenv("API_KEY")
	c := claude.NewClient(apiKey)
	m := claude.RequestBodyMessages{
		Model:     "claude-3-opus-20240229",
		MaxTokens: 1024,
		Messages: []claude.RequestBodyMessagesMessages{
			{
				Role: claude.MessagesRoleUser,
				ContentTypeImage: []claude.RequestBodyMessagesMessagesContentTypeImage{
					{
						Source: claude.RequestBodyMessagesMessagesContentTypeImageSource{
							Type:      "base64",
							MediaType: "image/png",
							Data:      "iVBORw0KGgoAAAANSUhEUgAAAQAAAAEACAIAAADTED8xAAADMElEQVR4nOzVwQnAIBQFQYXff81RUkQCOyDj1YOPnbXWPmeTRef+/3O/OyBjzh3CD95BfqICMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMK0CMO0TAAD//2Anhf4QtqobAAAAAElFTkSuQmCC",
						},
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

func Example_stream() {
	apiKey := os.Getenv("API_KEY")
	c := claude.NewClient(apiKey)
	m := claude.RequestBodyMessages{
		Model:     "claude-3-opus-20240229",
		MaxTokens: 64,
		Messages: []claude.RequestBodyMessagesMessages{
			{
				Role:    claude.MessagesRoleUser,
				Content: "Hello, world!",
			},
		},
	}
	ctx := context.Background()
	stream, err := c.CreateMessagesStream(ctx, m)
	if err != nil {
		panic(err)
	}
	defer stream.Close()
	for {
		res, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s", res.Content[0].Text)
	}
	fmt.Println()
	// Output:
	// Hello! How can I assist you today?
	//
}
