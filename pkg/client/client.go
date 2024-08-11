package client

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	HTTPClient *http.Client
	Request    *http.Request
	Body       io.Reader
	SESSDATA   string
	CSRF       string
	Resp       *http.Response
	Error      error
}

func New() *Client {
	return &Client{
		HTTPClient: http.DefaultClient,
		Request:    &http.Request{},
	}
}

func (c *Client) FormData(data map[string]string) *Client {
	formData := url.Values{}
	for key, value := range data {
		formData.Set(key, value)
	}
	c.Body = bytes.NewBufferString(formData.Encode())
	return c
}

func (c *Client) NewRequest(method, url string) *Client {
	req, _ := http.NewRequest(method, url, c.Body)
	c.Request = req
	return c
}

// 添加 SESSDATA Cookie
func (c *Client) WithSESSDATA() *Client {
	if c.Request.Header == nil {
		c.Request.Header = http.Header{}
	}
	c.Request.Header.Set("Cookie", "SESSDATA="+c.SESSDATA)

	return c
}

// 添加 Content-Type
func (c *Client) WithContentType(contentType string) *Client {
	if c.Request != nil {
		c.Request.Header.Set("Content-Type", contentType)
	}
	return c
}

func (c *Client) Do() *Client {
	resp, err := c.HTTPClient.Do(c.Request)
	if err != nil {
		c.Error = err
	} else {
		c.Resp = resp
	}

	return c
}

func (c *Client) Json(v interface{}) error {
	if c.Error != nil {
		return c.Error
	}

	defer c.Resp.Body.Close()

	data, err := io.ReadAll(c.Resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, v)
	if err != nil {
		return err
	}

	return nil
}
