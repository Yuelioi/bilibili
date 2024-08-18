package article

import (
	"fmt"
	"net/http"

	"github.com/Yuelioi/bilibili/pkg/endpoints/login"
)

// 获取用户专栏文章列表
//
// 参数：
//   - mid (int): 用户uid
//   - pn (int): 默认1 页数
//   - ps (int): 默认：30 范围：[1,30]
//   - sort	str		publish_time：最新发布 view：最多阅读 fav：最多收藏 默认：publish_time
//
// 备注：
//   - 认证方式：Cookie（SESSDATA）
//   - 鉴权方式：Wbi 签名
func (a *Article) ArticleList(mid, pn, ps int, sort string) (*ListResponse, error) {
	baseURL := "https://api.bilibili.com/x/space/wbi/article"

	formData := map[string]string{
		"mid":  fmt.Sprintf("%d", mid),
		"pn":   fmt.Sprintf("%d", pn),
		"ps":   fmt.Sprintf("%d", ps),
		"sort": sort,
	}

	newUrl, err := login.New(a.client).SignAndGenerateURL(baseURL)

	if err != nil {
		return nil, err
	}

	resp, err := a.client.HTTPClient.R().
		SetHeader("User-Agent", a.client.UserAgent).
		SetFormData(formData).
		SetCookie(&http.Cookie{
			Name:  "SESSDATA",
			Value: a.client.SESSDATA,
		}).
		SetResult(&ListResponse{}).
		Get(newUrl)

	if err != nil {
		return nil, err
	}

	return resp.Result().(*ListResponse), nil
}

// 获取用户专栏文集列表
//
// 参数：
//   - mid (int): 用户uid
//   - sort	str		0：最近更新 1:最多阅读
//
// 备注：
//   - 认证方式：Cookie（SESSDATA）
func (a *Article) ReadList(mid, sort int) (*ReadListResponse, error) {
	baseURL := "https://api.bilibili.com/x/article/up/lists"

	formData := map[string]string{
		"mid":  fmt.Sprintf("%d", mid),
		"sort": fmt.Sprintf("%d", sort),
	}

	resp, err := a.client.HTTPClient.R().
		SetHeader("User-Agent", a.client.UserAgent).
		SetFormData(formData).
		SetCookie(&http.Cookie{
			Name:  "SESSDATA",
			Value: a.client.SESSDATA,
		}).
		SetResult(&ReadListResponse{}).
		Get(baseURL)

	if err != nil {
		return nil, err
	}

	return resp.Result().(*ReadListResponse), nil
}

// --------------ArticleList-------------------

// ArticleList根对象
type ListResponse struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	TTL     int              `json:"ttl"`
	Data    ListResponseData `json:"data"`
}

// data对象
type ListResponseData struct {
	Articles []ListArticle `json:"articles"`
	PN       int           `json:"pn"`
	PS       int           `json:"ps"`
	Count    int           `json:"count"`
}

// data对象 -> articles -> article
type ListArticle struct {
	ID              int            `json:"id"`
	Category        CategoryItem   `json:"category"`
	Categories      []CategoryItem `json:"categories"`
	Title           string         `json:"title"`
	Summary         string         `json:"summary"`
	BannerURL       string         `json:"banner_url"`
	TemplateID      int            `json:"template_id"`
	State           int            `json:"state"`
	Author          ListAuthor     `json:"author"`
	Reprint         int            `json:"reprint"`
	ImageURLs       []string       `json:"image_urls"`
	PublishTime     int            `json:"publish_time"`
	CTime           int            `json:"ctime"`
	Stats           Stats          `json:"stats"`
	Tags            []Tag          `json:"tags"`
	Words           int            `json:"words"`
	Dynamic         string         `json:"dynamic"`
	OriginImageURLs []string       `json:"origin_image_urls"`
	List            interface{}    `json:"list"`
	IsLike          bool           `json:"is_like"`
	Media           Media          `json:"media"`
	ApplyTime       string         `json:"apply_time"`
	CheckTime       string         `json:"check_time"`
	Original        int            `json:"original"`
	ActID           int            `json:"act_id"`
	Dispute         interface{}    `json:"dispute"`
	AuthenMark      interface{}    `json:"authenMark"`
	CoverAvid       int            `json:"cover_avid"`
	TopVideoInfo    interface{}    `json:"top_video_info"`
	Type            int            `json:"type"`
}

// data对象 -> articles数组中的对象 -> categories数组中的对象
type CategoryItem struct {
	ID       int    `json:"id"`
	ParentID int    `json:"parent_id"`
	Name     string `json:"name"`
}

// data对象 -> articles数组中的对象 -> author对象
type ListAuthor struct {
	Mid            int            `json:"mid"`
	Name           string         `json:"name"`
	Face           string         `json:"face"`
	Pendant        Pendant        `json:"pendant"`
	OfficialVerify OfficialVerify `json:"official_verify"`
	Nameplate      Nameplate      `json:"nameplate"`
	Vip            Vip            `json:"vip"`
}

// data对象 -> articles数组中的对象 -> author对象 -> pendant对象
type Pendant struct {
	PID    int    `json:"pid"`
	Name   string `json:"name"`
	Image  string `json:"image"`
	Expire int    `json:"expire"`
}

// data对象 -> articles数组中的对象 -> author对象 -> vip对象
type Vip struct {
	Type            int      `json:"type"`
	Status          int      `json:"status"`
	DueDate         int      `json:"due_date"`
	VipPayType      int      `json:"vip_pay_type"`
	ThemeType       int      `json:"theme_type"`
	Label           VipLabel `json:"label"`
	AvatarSubscript int      `json:"avatar_subscript"`
	NicknameColor   string   `json:"nickname_color"`
}

// data对象 -> articles数组中的对象 -> author对象 -> vip对象 -> label对象
type VipLabel struct {
	Path       string `json:"path"`
	Text       string `json:"text"`
	LabelTheme string `json:"label_theme"`
}

// data对象 -> articles数组中的对象 -> tags数组中的对象
type Tag struct {
	TID  int    `json:"tid"`
	Name string `json:"name"`
}

// data对象 -> articles数组中的对象 -> media对象
type Media struct {
	Score    int    `json:"score"`
	MediaID  int    `json:"media_id"`
	Title    string `json:"title"`
	Cover    string `json:"cover"`
	Area     string `json:"area"`
	TypeID   int    `json:"type_id"`
	TypeName string `json:"type_name"`
	Spoiler  int    `json:"spoiler"`
}

// --------------ReadList-------------------
// 根对象
type ReadListResponse struct {
	Code    int          `json:"code"`    // 响应码 0：成功 -400：请求错误
	Message string       `json:"message"` // 0
	TTL     int          `json:"ttl"`     // 1
	Data    ReadListData `json:"data"`    // 信息本体
}

// data对象
type ReadListData struct {
	Lists []List `json:"lists"` // 文集信息列表
	Total int    `json:"total"` // 文集总数
}
