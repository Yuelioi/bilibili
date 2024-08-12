package bpi

import (
	"bilibili/pkg/client"
	"bilibili/pkg/endpoints/article"
)

type BPI struct {
	Article *article.Article
	Client  *client.Client
}

// New initializes the App struct and all its modules
func New(cli *client.Client) *BPI {

	return &BPI{
		Client:  cli,
		Article: article.New(cli),
	}
}
