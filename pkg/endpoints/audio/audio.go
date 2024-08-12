package audio

import "bilibili/pkg/client"

type Audio struct {
	client *client.Client
}

func New(client *client.Client) *Audio {
	return &Audio{client}
}
