package login

import "github.com/Yuelioi/bilibili/pkg/client"

type Login struct {
	client *client.Client
}

func New(client *client.Client) *Login {
	return &Login{client}
}
