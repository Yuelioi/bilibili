package article

import (
	"fmt"
	"net/http"

	"github.com/Yuelioi/bilibili/pkg/misc"
)

// 获取文集基本信息
//
// 参数：
//   - id (int): 文集rlid
//
// 备注：
//   - 认证方式：Cookie（SESSDATA）
func (a *Article) Articles(id int) (*ArticlesResponseData, error) {
	baseURL := "https://api.bilibili.com/x/article/list/web/articles"

	formData := map[string]string{
		"id": fmt.Sprintf("%d", id),
	}

	resp, err := a.client.HTTPClient.R().
		SetHeader("Content-Type", "application/json").
		SetFormData(formData).
		SetCookie(&http.Cookie{
			Name:  "SESSDATA",
			Value: a.client.SESSDATA,
		}).
		SetResult(&ArticlesResponseData{}).
		Get(baseURL)

	if err != nil {
		return nil, err
	}

	return resp.Result().(*ArticlesResponseData), nil
}

type ArticlesResponseData struct {
	misc.BaseResponse
	Data ArticlesData `json:"data"` // 信息本体
}
type ArticlesData struct {
	List      List         `json:"list"`      // 文集概览
	Articles  []ArticleOne `json:"articles"`  // 文集内的文章列表
	Author    Author       `json:"author"`    // 文集作者信息
	Last      ArticleOne   `json:"last"`      // 作用尚不明确，结构与 data.articles[] 中相似
	Attention bool         `json:"attention"` // 是否关注文集作者
}

type List struct {
	ID          int    `json:"id"`             // 文集rlid
	MID         int    `json:"mid"`            // 文集作者mid
	Name        string `json:"name"`           // 文集名称
	ImageURL    string `json:"image_url"`      // 文集封面图片url
	UpdateTime  int64  `json:"update_time"`    // 文集更新时间，时间戳
	CTime       int64  `json:"ctime"`          // 文集创建时间，时间戳
	PublishTime int64  `json:"publish_time"`   // 文集发布时间，时间戳
	Summary     string `json:"summary"`        // 文集简介
	Words       int    `json:"words"`          // 文集字数
	Read        int    `json:"read"`           // 文集阅读量
	ArticlesCnt int    `json:"articles_count"` // 文集内文章数量
	State       int    `json:"state"`          // 1 或 3，作用尚不明确
	Reason      string `json:"reason"`         // 空，作用尚不明确
	ApplyTime   string `json:"apply_time"`     // 空，作用尚不明确
	CheckTime   string `json:"check_time"`     // 空，作用尚不明确
}

type ArticleOne struct {
	ID          int        `json:"id"`           // 专栏cvid
	Title       string     `json:"title"`        // 文章标题
	State       int        `json:"state"`        // 作用尚不明确
	PublishTime int64      `json:"publish_time"` // 发布时间，秒时间戳
	Words       int        `json:"words"`        // 文章字数
	ImageURLs   []string   `json:"image_urls"`   // 文章封面
	Category    Category   `json:"category"`     // 文章标签
	Categories  []Category `json:"categories"`   // 文章标签列表
	Summary     string     `json:"summary"`      // 文章摘要
	Stats       Stats      `json:"stats"`        // 文章状态数信息
	LikeState   int        `json:"like_state"`   // 是否点赞
}

type Category struct {
	ID   int    `json:"id"`   // 标签ID
	Name string `json:"name"` // 标签名称
}

type Stats struct {
	View     int `json:"view"`     // 阅读数
	Favorite int `json:"favorite"` // 收藏数
	Like     int `json:"like"`     // 点赞数
	Dislike  int `json:"dislike"`  // 点踩数
	Reply    int `json:"reply"`    // 评论数
	Share    int `json:"share"`    // 分享数
	Coin     int `json:"coin"`     // 投币数
	Dynamic  int `json:"dynamic"`  // 动态转发数
}

type Author struct {
	MID            int            `json:"mid"`             // 作者mid
	Name           string         `json:"name"`            // 作者昵称
	Face           string         `json:"face"`            // 作者头像url
	OfficialVerify OfficialVerify `json:"official_verify"` // 作者认证信息
	Nameplate      Nameplate      `json:"nameplate"`       // 作者勋章
	VIP            VIP            `json:"vip"`             // 作者大会员状态
}

type OfficialVerify struct {
	Type int    `json:"type"` // 认证类型
	Desc string `json:"desc"` // 认证描述
}

type Nameplate struct {
	NID        int    `json:"nid"`   // 勋章ID
	Name       string `json:"name"`  // 勋章名称
	Image      string `json:"image"` // 勋章图片url
	ImageSmall string `json:"image_small"`
	Level      string `json:"level"`
	Condition  string `json:"condition"`
}
type VIP struct {
	Type       int   `json:"type"`        // 大会员类型
	Status     int   `json:"status"`      // 大会员状态
	ExpireTime int64 `json:"expire_time"` // 大会员到期时间
}
