package client

import (
	"github.com/go-resty/resty/v2"
)

type Client struct {
	HTTPClient *resty.Client
	DedeUserID int
	SESSDATA   string
	CSRF       string
}

func New() *Client {
	return &Client{HTTPClient: resty.New()}
}
