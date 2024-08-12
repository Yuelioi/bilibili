package video

import "bilibili/pkg/client"

type Video struct {
	client *client.Client
}

func New(client *client.Client) *Video {
	return &Video{client}
}
