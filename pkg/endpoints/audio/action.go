package audio

import (
	"fmt"
	"net/http"
)

// 查询音频收藏状态
//
// 参数：
//   - sid (int):音频auid
//
// 备注：
//   - 认证方式：Cookie（SESSDATA）
func (a *Audio) Collect(sid int) (*CollectResponse, error) {
	baseURL := "https://www.bilibili.com/audio/music-service-c/web/collections/songs-coll"

	formData := map[string]string{
		"sid": fmt.Sprintf("%d", sid),
	}

	resp, err := a.client.HTTPClient.R().
		SetFormData(formData).
		SetCookies([]*http.Cookie{
			{
				Name:  "SESSDATA",
				Value: a.client.SESSDATA},
			{
				Name:  "DedeUserID",
				Value: fmt.Sprint(a.client.DedeUserID)}}).
		SetResult(&CollectResponse{}).
		Get(baseURL)

	if err != nil {
		return nil, err
	}

	return resp.Result().(*CollectResponse), nil
}

// 查询音频投币数
//
// 参数：
//   - sid (int): 音频auid
//
// 备注：
//   - 认证方式：Cookie（SESSDATA）
//   - 鉴权方式：Cookie中DedeUserID存在且不为0
func (a *Audio) Coin(sid int) (*CoinResponse, error) {
	baseURL := "https://www.bilibili.com/audio/music-service-c/web/coin/audio"

	formData := map[string]string{
		"sid": fmt.Sprintf("%d", sid),
	}

	resp, err := a.client.HTTPClient.R().
		SetFormData(formData).
		SetCookies([]*http.Cookie{
			{
				Name:  "SESSDATA",
				Value: a.client.SESSDATA},
			{
				Name:  "DedeUserID",
				Value: fmt.Sprint(a.client.DedeUserID)}}).
		SetResult(&CoinResponse{}).
		Get(baseURL)

	if err != nil {
		return nil, err
	}

	return resp.Result().(*CoinResponse), nil
}

// 投币到指定音频
//
// 参数：
//   - sid (int): 音频 auid
//   - multiply (int): 投币数量，最大为2
//   - csrf (string): CSRF Token（从Cookie中获取）
//
// 备注：
//   - 认证方式：Cookie（SESSDATA）
func (a *Audio) AddCoin(sid, multiply int) (*AddCoinResponse, error) {
	baseURL := "https://www.bilibili.com/audio/music-service-c/web/coin/add"

	formData := map[string]string{
		"sid":      fmt.Sprintf("%d", sid),
		"multiply": fmt.Sprintf("%d", multiply),
		"csrf":     a.client.CSRF,
	}

	resp, err := a.client.HTTPClient.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(formData).
		SetCookies([]*http.Cookie{
			{
				Name:  "SESSDATA",
				Value: a.client.SESSDATA},
			{
				Name:  "DedeUserID",
				Value: fmt.Sprint(a.client.DedeUserID)}}).
		SetResult(&AddCoinResponse{}).
		Post(baseURL)

	if err != nil {
		return nil, err
	}

	return resp.Result().(*AddCoinResponse), nil
}

type CollectResponse struct {
	Code int    `json:"code"` // 返回值: 0表示成功, 72000000表示参数错误, 72010002表示账号未登录, 7201006表示该音频不存在或已被下架
	Msg  string `json:"msg"`  // 错误信息, 默认为"success"
	Data bool   `json:"data"` // 是否收藏: false表示未收藏, true表示已收藏
}

type CoinResponse struct {
	Code int    `json:"code"` // 返回值: 0表示成功, 72000000表示参数错误, 72010002表示账号未登录, 7201006表示该音频不存在或已被下架
	Msg  string `json:"msg"`  // 错误信息, 默认为"success"
	Data int    `json:"data"` // 投币数量: 0为未投币，上限为2
}

type AddCoinResponse struct {
	Code int    `json:"code"` // 返回值: 0表示成功, 72000000表示参数错误, 72010002表示账号未登录, 7201006表示该音频不存在或已被下架
	Msg  string `json:"msg"`  // 错误信息, 默认为 "success"
	Data string `json:"data"` // 当前投币数量: 0为未投币，上限为2
}
