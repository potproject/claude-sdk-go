package v1

func UseNoCache() *RequestCacheControl {
	return nil
}

func UseCacheEphemeral() *RequestCacheControl {
	t := RequestCacheControl{
		Type: "ephemeral",
	}
	return &t
}

func UseSystemNoCache(text string) RequestBodySystemTypeText {
	return RequestBodySystemTypeText{
		Type:         "text",
		Text:         text,
		CacheControl: nil,
	}
}

func UseSystemCacheEphemeral(text string) RequestBodySystemTypeText {
	return RequestBodySystemTypeText{
		Type:         "text",
		Text:         text,
		CacheControl: UseCacheEphemeral(),
	}
}
