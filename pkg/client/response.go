package client

type BaseResponse struct {
	Code    int    `json:"code"`    // 返回值，0 表示成功，其他值表示错误
	Message string `json:"message"` // 错误信息，默认为0
	TTL     int    `json:"ttl"`     // 1
}
