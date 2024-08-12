package login

import (
	"net/http"
	"net/url"
)

// Wbi 签名 获取keys
func (a *Login) UserKeys() (*Keys, error) {
	baseURL := "https://api.bilibili.com/x/web-interface/nav"
	fullURL := baseURL

	resp, err := a.client.HTTPClient.R().
		SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36 Edg/127.0.0.0").
		SetResult(&NavUserInfoResponse{}).
		Get(fullURL)

	if err != nil {
		return nil, err
	}

	res := resp.Result().(*NavUserInfoResponse)

	keys := &Keys{
		extract(res.Data.WbiImg.ImgURL),
		extract(res.Data.WbiImg.SubURL),
	}

	return keys, nil
}

// Wbi 签名新链接
func (a *Login) SignAndGenerateURL(urlStr string) (string, error) {
	urlObj, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}
	imgKey, subKey := getWbiKeysCached(a)
	query := urlObj.Query()
	params := map[string]string{}
	for k, v := range query {
		params[k] = v[0]
	}
	newParams := encWbi(params, imgKey, subKey)
	for k, v := range newParams {
		query.Set(k, v)
	}
	urlObj.RawQuery = query.Encode()
	newUrlStr := urlObj.String()
	return newUrlStr, nil
}

// 导航栏用户信息
//
// 备注：
//   - 认证方式：仅可Cookie（SESSDATA）
func (a *Login) NavUserInfo() (*NavUserInfoResponse, error) {
	baseURL := "https://api.bilibili.com/x/web-interface/nav"
	fullURL := baseURL

	resp, err := a.client.HTTPClient.R().
		SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36 Edg/127.0.0.0").
		SetCookie(&http.Cookie{
			Name:  "SESSDATA",
			Value: a.client.SESSDATA,
		}).
		SetResult(&NavUserInfoResponse{}).
		Get(fullURL)

	if err != nil {
		return nil, err
	}

	return resp.Result().(*NavUserInfoResponse), nil
}

// 登录用户状态数（双端）
//
// 备注：
//   - 认证方式：Cookie（SESSDATA）或APP
func (a *Login) UserState() (*UserStateResponse, error) {
	fullURL := "https://api.bilibili.com/x/web-interface/nav/stat"

	resp, err := a.client.HTTPClient.R().
		// SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36 Edg/127.0.0.0").
		SetCookie(&http.Cookie{
			Name:  "SESSDATA",
			Value: a.client.SESSDATA,
		}).
		SetResult(&UserStateResponse{}).
		Get(fullURL)

	if err != nil {
		return nil, err
	}

	return resp.Result().(*UserStateResponse), nil
}

type Keys struct {
	ImgURL string `json:"img_url"`
	SubURL string `json:"sub_url"`
}

// UserInfoResponse 根对象
type NavUserInfoResponse struct {
	Code    int          `json:"code"`    // 返回值 0：成功 -101：账号未登录
	Message string       `json:"message"` // 错误信息 默认为0
	TTL     int          `json:"ttl"`     // 1
	Data    UserInfoData `json:"data"`    // 信息本体
}

// UserInfoData data对象
type UserInfoData struct {
	IsLogin            bool           `json:"isLogin"`              // 是否已登录 false：未登录 true：已登录
	EmailVerified      int            `json:"email_verified"`       // 是否验证邮箱地址 0：未验证 1：已验证
	Face               string         `json:"face"`                 // 用户头像 url
	LevelInfo          LevelInfo      `json:"level_info"`           // 等级信息
	Mid                int64          `json:"mid"`                  // 用户 mid
	MobileVerified     int            `json:"mobile_verified"`      // 是否验证手机号 0：未验证 1：已验证
	Money              float64        `json:"money"`                // 拥有硬币数
	Moral              int            `json:"moral"`                // 当前节操值 上限为70
	Official           Official       `json:"official"`             // 认证信息
	OfficialVerify     OfficialVerify `json:"officialVerify"`       // 认证信息 2
	Pendant            Pendant        `json:"pendant"`              // 头像框信息
	Scores             int            `json:"scores"`               // （？）
	Uname              string         `json:"uname"`                // 用户昵称
	VipDueDate         int64          `json:"vipDueDate"`           // 会员到期时间 毫秒 时间戳
	VipStatus          int            `json:"vipStatus"`            // 会员开通状态 0：无 1：有
	VipType            int            `json:"vipType"`              // 会员类型 0：无 1：月度大会员 2：年度及以上大会员
	VipPayType         int            `json:"vip_pay_type"`         // 会员开通状态 0：无 1：有
	VipThemeType       int            `json:"vip_theme_type"`       // （？）
	VipLabel           VipLabel       `json:"vip_label"`            // 会员标签
	VipAvatarSubscript int            `json:"vip_avatar_subscript"` // 是否显示会员图标 0：不显示 1：显示
	VipNicknameColor   string         `json:"vip_nickname_color"`   // 会员昵称颜色 颜色码
	Wallet             Wallet         `json:"wallet"`               // B币钱包信息
	HasShop            bool           `json:"has_shop"`             // 是否拥有推广商品 false：无 true：有
	ShopURL            string         `json:"shop_url"`             // 商品推广页面 url
	AllowanceCount     int            `json:"allowance_count"`      // （？）
	AnswerStatus       int            `json:"answer_status"`        // （？）
	IsSeniorMember     int            `json:"is_senior_member"`     // 是否硬核会员 0：非硬核会员 1：硬核会员
	WbiImg             WbiImg         `json:"wbi_img"`              // Wbi 签名实时口令 该字段即使用户未登录也存在
	IsJury             bool           `json:"is_jury"`              // 是否风纪委员 true：风纪委员 false：非风纪委员
}

// LevelInfo data中的level_info对象
type LevelInfo struct {
	CurrentLevel int    `json:"current_level"` // 当前等级
	CurrentMin   int    `json:"current_min"`   // 当前等级经验最低值
	CurrentExp   int    `json:"current_exp"`   // 当前经验
	NextExp      string `json:"next_exp"`      // 升级下一等级需达到的经验 当用户等级为Lv6时，值为--，代表无穷大
}

// Official data中的official对象
type Official struct {
	Role  int    `json:"role"`  // 认证类型 见用户认证类型一览
	Title string `json:"title"` // 认证信息 无为空
	Desc  string `json:"desc"`  // 认证备注 无为空
	Type  int    `json:"type"`  // 是否认证 -1：无 0：认证
}

// OfficialVerify data中的official_verify对象
type OfficialVerify struct {
	Type int    `json:"type"` // 是否认证 -1：无 0：认证
	Desc string `json:"desc"` // 认证信息 无为空
}

// Pendant data中的pendant对象
type Pendant struct {
	PID    int    `json:"pid"`    // 挂件id
	Name   string `json:"name"`   // 挂件名称
	Image  string `json:"image"`  // 挂件图片url
	Expire int64  `json:"expire"` // （？）
}

// VipLabel data中的vip_label对象
type VipLabel struct {
	Path       string `json:"path"`        // （？）
	Text       string `json:"text"`        // 会员名称
	LabelTheme string `json:"label_theme"` // 会员标签 vip：大会员 annual_vip：年度大会员 ten_annual_vip：十年大会员 hundred_annual_vip：百年大会员
}

// Wallet data中的wallet对象
type Wallet struct {
	Mid           int64 `json:"mid"`             // 登录用户mid
	BCoinBalance  int   `json:"bcoin_balance"`   // 拥有B币数
	CouponBalance int   `json:"coupon_balance"`  // 每月奖励B币数
	CouponDueTime int64 `json:"coupon_due_time"` // （？）
}

// WbiImg data中的wbi_img对象
type WbiImg struct {
	ImgURL string `json:"img_url"` // Wbi 签名参数 imgKey的伪装 url 详见文档 Wbi 签名
	SubURL string `json:"sub_url"` // Wbi 签名参数 subKey的伪装 url 详见文档 Wbi 签名
}

// ------------------------

// UserInfoResponse 根对象
type UserStateResponse struct {
	Code    int           `json:"code"`    // 返回值 0：成功 -101：账号未登录
	Message string        `json:"message"` // 错误信息 默认为0
	TTL     int           `json:"ttl"`     // 1
	Data    UserStateData `json:"data"`    // 信息本体
}

// UserInfoResponse 根对象
type UserStateData struct {
	Following    int `json:"following"`     // 关注数
	Follower     int `json:"follower"`      // 粉丝数
	DynamicCount int `json:"dynamic_count"` // 发布动态数
}
