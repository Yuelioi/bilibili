package article

import "bilibili/pkg/client"

type CoinResponse struct {
	client.BaseResponse
	Data struct {
		Like bool `json:"like"`
	} `json:"data"`
}
