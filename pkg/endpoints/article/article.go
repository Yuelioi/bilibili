package article

import "github.com/Yuelioi/bilibili/pkg/client"

type Article struct {
	client *client.Client
}

func New(client *client.Client) *Article {
	return &Article{client}
}
