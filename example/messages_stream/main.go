package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	claude "github.com/potproject/claude-sdk-go"
)

func main() {

	const CLAUDE_DEFAULT = "claude-3-haiku-20240307"

	apiKey, exists := os.LookupEnv("ANTHROPIC_API_KEY")
	if !exists {
		log.Fatal("ANTHROPIC_API_KEY environment variable is not set")
	}
	client := claude.NewClient(apiKey)

	body := claude.RequestBodyMessages{
		Model: CLAUDE_DEFAULT,
		Messages: []claude.RequestBodyMessagesMessages{
			{
				Role:    claude.MessagesRoleUser,
				Content: "Hello, Claude!",
			},
		},
		MaxTokens: 100,
		Stream:    true,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	stream, err := client.CreateMessagesStream(ctx, body)
	if err != nil {
		log.Fatalf("Failed to create message stream: %v", err)
	}
	defer stream.Close()

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("Stream ended.")
			break
		}
		if err != nil {
			log.Fatalf("Error receiving message: %v", err)
		}
		fmt.Printf("Received message: %+v\n", msg)
		fmt.Printf("Usage - Input Tokens: %d, Output Tokens: %d\n", msg.Usage.InputTokens, msg.Usage.OutputTokens)
	}
}
