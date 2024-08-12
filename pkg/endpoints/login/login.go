package login

import "bilibili/pkg/client"

type Login struct {
	client *client.Client
}

func New(client *client.Client) *Login {
	return &Login{client}
}
