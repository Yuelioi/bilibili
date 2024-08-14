package audio

import "github.com/Yuelioi/bilibili/pkg/client"

type Audio struct {
	client *client.Client
}

func New(client *client.Client) *Audio {
	return &Audio{client}
}
