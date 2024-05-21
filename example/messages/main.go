package main

import (
	"context"
	"fmt"
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
		Stream:    false, // Set to false for non-streaming
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	response, err := client.CreateMessages(ctx, body)
	if err != nil {
		log.Fatalf("Failed to create message: %v", err)
	}

	fmt.Printf("Received message: %+v\n", response)
	fmt.Printf("Usage - Input Tokens: %d, Output Tokens: %d\n", response.Usage.InputTokens, response.Usage.OutputTokens)
}
