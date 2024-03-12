package v1

import (
	"net/http"
)

type Client struct {
	config ClientConfig
}

const (
	defaultBaseURL  = "https://api.anthropic.com/"
	defaultEndpoint = "v1/messages"
	defaultVersion  = "2023-06-01"
	defaultBeta     = ""
)

const contentType = "application/json"

type ClientConfig struct {
	ApiKey string

	Version string
	Beta    string // Using Beta API : anthropic-beta Header

	BaseURL    string
	Endpoint   string
	HTTPClient *http.Client
}

func defaultConfig(apiKey string) ClientConfig {
	return ClientConfig{
		ApiKey: apiKey,

		Version: defaultVersion,
		Beta:    "",

		BaseURL:    defaultBaseURL,
		Endpoint:   defaultEndpoint,
		HTTPClient: &http.Client{},
	}
}

func NewClient(apiKey string) *Client {
	return &Client{
		config: defaultConfig(apiKey),
	}
}

func NewClientWithConfig(config ClientConfig) *Client {
	return &Client{
		config: config,
	}
}

func (c *Client) SetVersion(version string) {
	c.config.Version = version
}
