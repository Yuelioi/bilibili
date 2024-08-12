package tools

import (
	"fmt"
	"net/url"
)

func URLWithParams(baseURL string, params map[string]string) string {
	urlParams := url.Values{}
	for key, value := range params {
		urlParams.Add(key, value)
	}
	fullURL := fmt.Sprintf("%s?%s", baseURL, urlParams.Encode())
	return fullURL
}
