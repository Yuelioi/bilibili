package article

import (
	"fmt"
	"net/http"

	"github.com/Yuelioi/bilibili/pkg/misc"
)

// 点赞文章
//
// 参数：
//   - id (int): 文章cvid
//   - type (int): 1:点赞 2:取消点赞
func (a *Article) Like(id int, likeType int) (*misc.BaseResponse, error) {

	formData := map[string]string{
		"id":   fmt.Sprintf("%d", id),
		"type": fmt.Sprintf("%d", likeType),
		"csrf": a.client.CSRF,
	}

	resp, err := a.client.HTTPClient.R().
		SetFormData(formData).
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetCookie(&http.Cookie{
			Name:  "SESSDATA",
			Value: a.client.SESSDATA,
		}).
		SetResult(&misc.BaseResponse{}).
		Post("https://api.bilibili.com/x/article/like")

	if err != nil {
		return nil, err
	}

	return resp.Result().(*misc.BaseResponse), nil
}

// 投币文章
//
// 参数：
//   - id (int): 文章cvid
//   - upid (int): 文章作者mid
//   - multiply (int): 投币数量（上限为2）
//
// 备注：
//   - 必须有 csrf
//   - 认证方式：Cookie（SESSDATA）
func (a *Article) Coin(aid, upid, multiply int) (*CoinResponse, error) {

	formData := map[string]string{
		"aid":      fmt.Sprintf("%d", aid),
		"upid":     fmt.Sprintf("%d", upid),
		"multiply": fmt.Sprintf("%d", multiply),
		"avtype":   "2",
		"csrf":     a.client.CSRF,
	}

	resp, err := a.client.HTTPClient.R().
		SetFormData(formData).
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetCookie(&http.Cookie{
			Name:  "SESSDATA",
			Value: a.client.SESSDATA,
		}).
		SetResult(&CoinResponse{}).
		Post("https://api.bilibili.com/x/web-interface/coin/add")

	if err != nil {
		return nil, err
	}

	return resp.Result().(*CoinResponse), nil

}

// 收藏文章
//
// 参数：
//   - id (int): 文章cvid
//
// 备注：
//   - 必须有 csrf
//   - 认证方式：Cookie（SESSDATA）
func (a *Article) Favorite(id int) (*misc.BaseResponse, error) {
	formData := map[string]string{
		"id":   fmt.Sprintf("%d", id),
		"csrf": a.client.CSRF,
	}

	resp, err := a.client.HTTPClient.R().
		SetFormData(formData).
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetCookie(&http.Cookie{
			Name:  "SESSDATA",
			Value: a.client.SESSDATA,
		}).
		SetResult(&misc.BaseResponse{}).
		Post("https://api.bilibili.com/x/article/favorites/add")

	if err != nil {
		return nil, err
	}

	return resp.Result().(*misc.BaseResponse), nil
}

// 取消收藏文章
//
// 参数：
//   - id (int): 文章cvid
//
// 备注：
//   - 必须有 csrf
//   - 认证方式：Cookie（SESSDATA）
func (a *Article) UnFavorite(id int) (*misc.BaseResponse, error) {
	formData := map[string]string{
		"id":   fmt.Sprintf("%d", id),
		"csrf": a.client.CSRF,
	}

	resp, err := a.client.HTTPClient.R().
		SetFormData(formData).
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetCookie(&http.Cookie{
			Name:  "SESSDATA",
			Value: a.client.SESSDATA,
		}).
		SetResult(&misc.BaseResponse{}).
		Post("https://api.bilibili.com/x/article/favorites/del")

	if err != nil {
		return nil, err
	}

	return resp.Result().(*misc.BaseResponse), nil
}

type CoinResponse struct {
	misc.BaseResponse
	Data struct {
		Like bool `json:"like"`
	} `json:"data"`
}
