package v1

import (
	"encoding/base64"
	"os"
	"path/filepath"
)

func TypeImageSourceLoadUrl(url string) RequestBodyMessagesMessagesContentTypeImageSource {
	return RequestBodyMessagesMessagesContentTypeImageSource{
		Type: "url",
		Url:  url,
	}
}

func TypeImageSourceLoadFile(filePath string) (RequestBodyMessagesMessagesContentTypeImageSource, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return RequestBodyMessagesMessagesContentTypeImageSource{}, err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return RequestBodyMessagesMessagesContentTypeImageSource{}, err
	}

	size := info.Size()

	data := make([]byte, size)
	file.Read(data)

	b64 := base64.StdEncoding.EncodeToString(data)

	ext := filepath.Ext(filePath)
	var contentType string
	if ext == ".jpg" || ext == ".jpeg" {
		contentType = "image/jpeg"
	}
	if ext == ".png" {
		contentType = "image/png"
	}
	if ext == ".gif" {
		contentType = "image/gif"
	}
	if ext == ".webp" {
		contentType = "image/webp"
	}
	return RequestBodyMessagesMessagesContentTypeImageSource{
		Type:      "base64",
		MediaType: contentType,
		Data:      b64,
	}, nil
}
