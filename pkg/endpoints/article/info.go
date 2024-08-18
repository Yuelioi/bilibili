package article

import (
	"fmt"
	"net/http"

	"github.com/Yuelioi/bilibili/pkg/misc"
)

// 获取专栏文章基本信息
//
// 参数：
//   - id (int): 专栏cvid
//
// 备注：
//   - 认证方式：Cookie（SESSDATA）
//   - 必须有 User-Agent

func (a *Article) Article(id int) (*ArticleResponseData, error) {
	baseURL := "https://api.bilibili.com/x/article/viewinfo"

	formData := map[string]string{
		"id": fmt.Sprintf("%d", id),
	}

	resp, err := a.client.HTTPClient.R().
		SetHeader("User-Agent", a.client.UserAgent).
		SetFormData(formData).
		SetCookie(&http.Cookie{
			Name:  "SESSDATA",
			Value: a.client.SESSDATA,
		}).
		SetResult(&ArticleResponseData{}).
		Get(baseURL)

	if err != nil {
		return nil, err
	}

	return resp.Result().(*ArticleResponseData), nil
}

type ArticleResponseData struct {
	misc.BaseResponse
	Data ArticleInfoData `json:"data"` // 信息本体
}

type ArticleInfoData struct {
	Like            int            `json:"like"`              // 是否点赞
	Attention       bool           `json:"attention"`         // 是否关注文章作者
	Favorite        bool           `json:"favorite"`          // 是否收藏
	Coin            int            `json:"coin"`              // 为文章投币数
	Stats           Stats          `json:"stats"`             // 状态数信息(与articles.Stats一样)
	Title           string         `json:"title"`             // 文章标题
	BannerURL       string         `json:"banner_url"`        // 文章头图url
	MID             int            `json:"mid"`               // 文章作者mid
	AuthorName      string         `json:"author_name"`       // 文章作者昵称
	IsAuthor        bool           `json:"is_author"`         // 作用尚不明确
	ImageURLs       []string       `json:"image_urls"`        // 动态封面
	OriginImageURLs []string       `json:"origin_image_urls"` // 封面图片
	Shareable       bool           `json:"shareable"`         // 作用尚不明确
	ShowLaterWatch  bool           `json:"show_later_watch"`  // 作用尚不明确
	ShowSmallWindow bool           `json:"show_small_window"` // 作用尚不明确
	InList          bool           `json:"in_list"`           // 是否收于文集
	Pre             int            `json:"pre"`               // 上一篇文章cvid
	Next            int            `json:"next"`              // 下一篇文章cvid
	ShareChannels   []ShareChannel `json:"share_channels"`    // 分享方式列表
	Type            int            `json:"type"`              // 文章类别
}

type ShareChannel struct {
	Name         string `json:"name"`          // 分享名称
	Picture      string `json:"picture"`       // 分享图片url
	ShareChannel string `json:"share_channel"` // 分享代号
}
