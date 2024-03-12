# claude-sdk-go
This is the unofficial Go SDK for the Anthropic Claude API.

https://docs.anthropic.com/claude/

## Support
* /v1/messages
  * Type=Text, Type=Image

## Getting Started
```bash
go get github.com/potproject/claude-sdk-go
```

## Example
### [Create a Message](./example/message/main.go)
```go
package main

import (
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
			},
		},
	}
	res, err := c.CreateMessages(m)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Content[0].Text)
	// Hello! How can I assist you today?
}

```

### [Create a Message with Image](./example/message_type_image/main.go)
```go
package main

import (
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
	res, err := c.CreateMessages(m)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Content[0].Text)
}

```
## LICENSE
MIT