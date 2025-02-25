# claude-sdk-go
[![Go Reference](https://pkg.go.dev/badge/github.com/potproject/claude-sdk-go.svg)](https://pkg.go.dev/github.com/potproject/claude-sdk-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/potproject/claude-sdk-go)](https://goreportcard.com/report/github.com/potproject/claude-sdk-go)

This is the unofficial Go SDK for the Anthropic Claude API.

It is designed with reference to the [sashabaranov/go-openai](https://github.com/sashabaranov/go-openai).

Official Docs: https://docs.anthropic.com/claude/

## Supported
* /v1/messages
  * Text Message
  * Image Message
  * Streaming Messages

## Getting Started
```bash
go get github.com/potproject/claude-sdk-go
```

## Example
### Create a Message
```go
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
		Model:     "claude-3-opus-20240229",
		MaxTokens: 64,
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

```

<details>
<summary>Create a Streaming Message</summary>

### Create a Streaming Message
```go
package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	claude "github.com/potproject/claude-sdk-go"
)

func main() {
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

```

</details>

<details>
<summary>Create a Message with Image</summary>

### Create a Message with Image
```go
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
							Data:      "iVBORw0KG...",
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
}

```

</details>


<details>
<summary>Create a Message with Image(Load File)</summary>

### Create a Message with Image(Load File)
```go
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
}

```

</details>

<details>
<summary>Create a Message (Use Thinking)</summary>

### Create a Message (Use Thinking)
```go
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

```

</details>

<details>
<summary>Create a Streaming Message (Use Thinking)</summary>

### Create a Streaming Message (Use Thinking)
```go
package main

import (
	"context"
	"errors"
	"fmt"
	"io"
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
				Content: "Guess the Earth's population in 2100",
			},
		},
	}
	ctx := context.Background()
	stream, err := c.CreateMessagesStream(ctx, m)
	if err != nil {
		panic(err)
	}
	defer stream.Close()
	streamType := ""
	for {
		res, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			panic(err)
		}
		if res.Content[0].Type == claude.ResponseBodyMessagesContentTypeThinking && streamType != claude.ResponseBodyMessagesContentTypeThinking {
			fmt.Println("[thinking]")
			streamType = claude.ResponseBodyMessagesContentTypeThinking
		}
		if res.Content[0].Type == claude.ResponseBodyMessagesContentTypeText && streamType != claude.ResponseBodyMessagesContentTypeText {
			fmt.Println()
			fmt.Println("[text]")
			streamType = claude.ResponseBodyMessagesContentTypeText
		}

		fmt.Printf("%s", res.Content[0].Thinking)
		fmt.Printf("%s", res.Content[0].Text)
	}
	fmt.Println()
}


```

</details>

## LICENSE
MIT