package article

import (
	"bilibili/pkg/client"
	"fmt"
)

type Article struct {
	client *client.Client
}

func New(client *client.Client) *Article {
	return &Article{client}
}

// 点赞文章
//
// 参数：
//   - id (int): 文章cvid
//   - type (int): 1:点赞 2:取消点赞
func (a *Article) Like(id int, likeType int) (*client.BaseResponse, error) {

	formData := map[string]string{
		"id":   fmt.Sprintf("%d", id),
		"type": fmt.Sprintf("%d", likeType),
		"csrf": a.client.CSRF,
	}

	resp := &client.BaseResponse{}
	err := a.client.FormData(formData).NewRequest("POST", "https://api.bilibili.com/x/article/like").
		WithSESSDATA().
		WithContentType("application/x-www-form-urlencoded").
		Do().Json(resp)

	return resp, err
}

// 投币文章
//
// 参数：
//   - id (int): 文章cvid
//   - upid (int): 文章作者mid
//   - multiply (int): 投币数量（上限为2）
//
// 备注：
//   - avtype必须为2
func (a *Article) Coin(aid, upid, multiply int) (*CoinResponse, error) {

	formData := map[string]string{
		"aid":      fmt.Sprintf("%d", aid),
		"upid":     fmt.Sprintf("%d", upid),
		"multiply": fmt.Sprintf("%d", multiply),
		"avtype":   "2",
		"csrf":     a.client.CSRF,
	}

	resp := &CoinResponse{}
	err := a.client.FormData(formData).NewRequest("POST", "https://api.bilibili.com/x/web-interface/coin/add").
		WithSESSDATA().
		WithContentType("application/x-www-form-urlencoded").
		Do().Json(resp)

	return resp, err
}

// 收藏文章
//
// 参数：
//   - id (int): 文章cvid
func (a *Article) Favorite(id int) (*client.BaseResponse, error) {

	formData := map[string]string{
		"id":   fmt.Sprintf("%d", id),
		"csrf": a.client.CSRF,
	}

	resp := &client.BaseResponse{}
	err := a.client.FormData(formData).NewRequest("POST", "https://api.bilibili.com/x/article/favorites/add").
		WithSESSDATA().
		WithContentType("application/x-www-form-urlencoded").
		Do().Json(resp)

	return resp, err
}

// 收藏文章
//
// 参数：
//   - id (int): 文章cvid
func (a *Article) UnFavorite(id int) (*client.BaseResponse, error) {

	formData := map[string]string{
		"id":   fmt.Sprintf("%d", id),
		"csrf": a.client.CSRF,
	}

	resp := &client.BaseResponse{}
	err := a.client.FormData(formData).NewRequest("POST", "https://api.bilibili.com/x/article/favorites/del").
		WithSESSDATA().
		WithContentType("application/x-www-form-urlencoded").
		Do().Json(resp)

	return resp, err
}
