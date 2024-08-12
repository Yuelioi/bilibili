package client

import (
	"github.com/go-resty/resty/v2"
)

type Client struct {
	HTTPClient *resty.Client
	DedeUserID int
	SESSDATA   string
	CSRF       string
	AccessKey  string
	Buvid3     string
}

func New() *Client {
	return &Client{HTTPClient: resty.New()}
}
