package video

import "github.com/Yuelioi/bilibili/pkg/client"

type Video struct {
	client *client.Client
}

func New(client *client.Client) *Video {
	return &Video{client}
}
