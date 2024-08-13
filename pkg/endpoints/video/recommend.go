package video

import (
	"fmt"
	"net/http"
)

// 获取单视频推荐列表（web端）
// Parameters:
//   - aid (int): 视频的 aid (可选)
//   - bvid (string): 视频的 bvid (可选)
func (v *Video) GetRelatedVideos(aid int, bvid string) (*RelatedVideosResponse, error) {
	baseURL := "https://api.bilibili.com/x/web-interface/archive/related"

	queryParams := map[string]string{
		"aid":  fmt.Sprintf("%d", aid),
		"bvid": bvid,
	}

	resp, err := v.client.HTTPClient.R().
		SetQueryParams(queryParams).
		SetResult(&RelatedVideosResponse{}).
		Get(baseURL)
	if err != nil {
		return nil, err
	}

	return resp.Result().(*RelatedVideosResponse), nil
}

// 获取首页视频推荐列表（web端）
//
// Parameters:
//   - freshType (int): 相关性 (默认为 4)
//   - ps (int): 单页返回的记录条数 (默认为 12, 最大值为 30)
//   - freshIdx (int): 当前翻页号 (以 1 开始)
//   - freshIdx1H (int): 当前翻页号(一小时前?)
//   - brush (int): 刷子? (以 1 开始)
//   - fetchRow (int): 本次抓取的最后一行行号
//   - webLocation (int): 网页位置 (主页为 1430650)
//   - yNum (int): 普通列数
//   - lastYNum (int): 总列数
//   - feedVersion (string): V8
//   - homepageVer (int): 首页版本 (默认为 1)
//   - screen (string): 浏览器视口大小 (水平在前垂直在后以减号分割)
//   - seoInfo (string): SEO信息
//   - lastShowlist (string): 上次抓取的视频av号列表
//   - uniqID (string): ??? (作用尚不明确)
//   - wRid (string): WBI 签名
//   - wts (int64): UNIX 时间戳
//
// 认证方式：Cookie（SESSDATA）
// 最多获取30条推荐视频,直播及推荐边栏
func (v *Video) GetHomePageRecommendations(
	freshType int,
	ps int,
	freshIdx int,
	freshIdx1H int,
	brush int,
	fetchRow int,
	webLocation int,
	yNum int,
	lastYNum int,
	feedVersion string,
	homepageVer int,
	screen string,
	seoInfo string,
	lastShowlist string,
	uniqID string,
	wRid string,
	wts int64,
) (*HomePageRcmdResponse, error) {
	baseURL := "https://api.bilibili.com/x/web-interface/wbi/index/top/feed/rcmd"

	queryParams := map[string]string{
		"fresh_type":    fmt.Sprintf("%d", freshType),
		"ps":            fmt.Sprintf("%d", ps),
		"fresh_idx":     fmt.Sprintf("%d", freshIdx),
		"fresh_idx_1h":  fmt.Sprintf("%d", freshIdx1H),
		"brush":         fmt.Sprintf("%d", brush),
		"fetch_row":     fmt.Sprintf("%d", fetchRow),
		"web_location":  fmt.Sprintf("%d", webLocation),
		"y_num":         fmt.Sprintf("%d", yNum),
		"last_y_num":    fmt.Sprintf("%d", lastYNum),
		"feed_version":  feedVersion,
		"homepage_ver":  fmt.Sprintf("%d", homepageVer),
		"screen":        screen,
		"seo_info":      seoInfo,
		"last_showlist": lastShowlist,
		"uniq_id":       uniqID,
		"w_rid":         wRid,
		"wts":           fmt.Sprintf("%d", wts),
	}

	resp, err := v.client.HTTPClient.R().
		SetQueryParams(queryParams).
		SetCookie(&http.Cookie{
			Name:  "SESSDATA",
			Value: v.client.SESSDATA,
		}).
		SetResult(&HomePageRcmdResponse{}).
		Get(baseURL)
	if err != nil {
		return nil, err
	}

	return resp.Result().(*HomePageRcmdResponse), nil
}

// 获取短视频模式视频列表
//
// Parameters:
//
//   - fnval (int): 视频流格式标识 (默认为 272)
//   - fnver (int): 视频流版本标识 (恒为 1)
//   - forceHost (int): 源url类型 (0:无限制, 1:使用http, 2:使用https)
//   - fourk (int): 是否允许 4K 视频 (0: 1080P, 1: 4K)
//   - guidance (int): 指导 (默认为 0)
//   - httpsUrlReq (int): https 请求标识 (默认为 0)
//   - inlineDanmu (int): 弹幕内嵌方式 (默认为 2)
//   - inlineSound (int): 音频内嵌方式 (默认为 1)
//   - interestId (int): 兴趣 ID (默认为 0)
//   - loginEvent (int): 登录状态 (0: 登录, 1: 未登录)
//   - mobiApp (string): 设备类型 (默认为 android)
//   - network (string): 网络类型 (默认为 wifi)
//   - openEvent (int): 开放事件 (非必要)
//   - platform (string): 设备平台 (默认为 android)
//   - pull (bool): 拉取标识 (默认为 false)
//   - qn (int): 画质 (默认为 32)
//   - recsysMode (int): 推荐系统模式 (默认为 0)
//   - sLocale (string): 语言 (默认为 zh_CN)
//   - videoMode (int): 视频模式 (默认为 1)
//   - voiceBalance (int): 音量均衡 (默认为 1)
//
// 认证方式：
//   - Cookie（SESSDATA）
func (v *Video) GetShortVideoList(
	fnval int,
	fnver int,
	forceHost int,
	fourk int,
	guidance int,
	httpsUrlReq int,
	inlineDanmu int,
	inlineSound int,
	interestId int,
	loginEvent int,
	mobiApp string,
	network string,
	openEvent int,
	platform string,
	pull bool,
	qn int,
	recsysMode int,
	sLocale string,
	videoMode int,
	voiceBalance int,
) (*ShortVideoResponse, error) {
	baseURL := "https://app.bilibili.com/x/v2/feed/index"

	queryParams := map[string]string{
		"fnval":         fmt.Sprintf("%d", fnval),
		"fnver":         fmt.Sprintf("%d", fnver),
		"force_host":    fmt.Sprintf("%d", forceHost),
		"fourk":         fmt.Sprintf("%d", fourk),
		"guidance":      fmt.Sprintf("%d", guidance),
		"https_url_req": fmt.Sprintf("%d", httpsUrlReq),
		"inline_danmu":  fmt.Sprintf("%d", inlineDanmu),
		"inline_sound":  fmt.Sprintf("%d", inlineSound),
		"interest_id":   fmt.Sprintf("%d", interestId),
		"login_event":   fmt.Sprintf("%d", loginEvent),
		"mobi_app":      mobiApp,
		"network":       network,
		"open_event":    fmt.Sprintf("%d", openEvent),
		"platform":      platform,
		"pull":          fmt.Sprintf("%t", pull),
		"qn":            fmt.Sprintf("%d", qn),
		"recsys_mode":   fmt.Sprintf("%d", recsysMode),
		"s_locale":      sLocale,
		"video_mode":    fmt.Sprintf("%d", videoMode),
		"voice_balance": fmt.Sprintf("%d", voiceBalance),
	}

	resp, err := v.client.HTTPClient.R().
		SetQueryParams(queryParams).
		SetCookie(&http.Cookie{
			Name:  "SESSDATA",
			Value: v.client.SESSDATA,
		}).
		SetResult(&ShortVideoResponse{}).
		Get(baseURL)
	if err != nil {
		return nil, err
	}

	return resp.Result().(*ShortVideoResponse), nil
}

type RelatedVideosResponse struct {
	Code    int       `json:"code"`    // 返回值, 0: 成功, -400: 请求错误
	Message string    `json:"message"` // 错误信息, 默认为 0
	TTL     int       `json:"ttl"`     // 默认值为 1
	Data    []Related `json:"data"`    // 推荐列表
}

type HomePageRcmdResponse struct {
	Code    int          `json:"code"`    // 返回值, 0: 成功, -400: 请求错误
	Message string       `json:"message"` // 错误信息, 默认为 0
	TTL     int          `json:"ttl"`     // 默认值为 1
	Data    HomePageData `json:"data"`    // 数据本体
}

// HomePageData represents the data section of the homepage recommendation response.
type HomePageData struct {
	Item             []HomePageItem `json:"item"`                     // 推荐列表
	SideBarColumn    []HomePageItem `json:"side_bar_column"`          // 边栏列表
	MID              int            `json:"mid"`                      // 用户mid, 未登录为0
	PreloadExposePct float64        `json:"preload_expose_pct"`       // 用于预加载
	PreloadFloorPct  float64        `json:"preload_floor_expose_pct"` // 用于预加载
	// Additional fields can be added based on actual response requirements...
}

// HomePageItem represents a single recommended video or live stream.
type HomePageItem struct {
	BVID       string             `json:"bvid"`        // 视频bvid
	CID        int                `json:"cid"`         // 稿件cid
	Title      string             `json:"title"`       // 标题
	Pic        string             `json:"pic"`         // 封面
	Pic4_3     string             `json:"pic_4_3"`     // 封面(4:3)
	Goto       string             `json:"goto"`        // 目标类型 (av: 视频, ogv: 边栏, live: 直播)
	ID         int                `json:"id"`          // 视频aid / 直播间id
	Pubdate    int                `json:"pubdate"`     // 发布时间
	Duration   int                `json:"duration"`    // 视频时长
	Owner      HomePageOwner      `json:"owner"`       // UP主信息
	Stat       HomePageStat       `json:"stat"`        // 视频状态信息
	RcmdReason HomePageRcmdReason `json:"rcmd_reason"` // 推荐理由
	URI        string             `json:"uri"`         // 目标页 URI
	// Additional fields can be added based on actual response requirements...
}

// HomePageOwner represents the owner (UP主) of the video or live stream.
type HomePageOwner struct {
	Face string `json:"face"` // 头像URL
	MID  int    `json:"mid"`  // UP主mid
	Name string `json:"name"` // UP昵称
}

// HomePageStat represents the status of the video.
type HomePageStat struct {
	View     int `json:"view"`     // 播放数
	Danmaku  int `json:"danmaku"`  // 弹幕数
	Reply    int `json:"reply"`    // 评论数
	Favorite int `json:"favorite"` // 收藏数
	Coin     int `json:"coin"`     // 投币数
	Share    int `json:"share"`    // 分享数
	Like     int `json:"like"`     // 点赞数
	// Additional fields can be added based on actual response requirements...
}

// HomePageRcmdReason represents the reason for recommending the video.
type HomePageRcmdReason struct {
	ReasonType int    `json:"reason_type"` // 原因类型, 0: 无, 1: 已关注, 3: 高点赞量
	Content    string `json:"content"`     // 原因描述, 当 reason_type 为 3 时存在
}

type ShortVideoResponse struct {
	Code    int            `json:"code"`    // 返回值, 0: 成功, -400: 请求错误
	Message string         `json:"message"` // 错误信息, 默认为 0
	TTL     int            `json:"ttl"`     // 默认值为 1
	Data    ShortVideoData `json:"data"`    // 数据本体
}

// ShortVideoData represents the data section of the short video mode response.
type ShortVideoData struct {
	Config interface{}      `json:"config"` // 一些界面相关的内容
	Items  []ShortVideoItem `json:"items"`  // 视频列表
}

// ShortVideoItem represents a single short video item.
type ShortVideoItem struct {
	CanPlay                      int                  `json:"can_play"`                         // 字面意思
	CardGoto                     string               `json:"card_goto"`                        // av
	CardType                     string               `json:"card_type"`                        // 卡片类型, 视频为small_cover_v2
	Cover                        string               `json:"cover"`                            // 封面url
	CoverLeft1ContentDescription string               `json:"cover_left_1_content_description"` // 播放量
	CoverLeft2ContentDescription string               `json:"cover_left_2_content_description"` // 弹幕数
	CoverLeftText1               string               `json:"cover_left_text_1"`                // 播放量
	CoverLeftText2               string               `json:"cover_left_text_2"`                // 弹幕数
	CoverRightContentDescription string               `json:"cover_right_content_description"`  // 视频长度
	CoverRightText               string               `json:"cover_right_text"`                 // 视频长度
	DescButton                   ShortVideoDescButton `json:"desc_button"`                      // up主信息
	Param                        string               `json:"param"`                            // 视频aid
	PlayerArgs                   ShortVideoPlayerArgs `json:"player_args"`                      // 视频信息
	TalkBack                     string               `json:"talk_back"`                        // 可能用于弹幕返回信息
	Title                        string               `json:"title"`                            // 标题
	URI                          string               `json:"uri"`                              // 跳转链接
}

// ShortVideoDescButton represents the UP主 information in the short video item.
type ShortVideoDescButton struct {
	Event string `json:"event"` // 事件
	Text  string `json:"text"`  // up名称
	Type  int    `json:"type"`  // 类型
	URI   string `json:"uri"`   // 跳转链接
}

// ShortVideoPlayerArgs represents video information in the short video item.
type ShortVideoPlayerArgs struct {
	Aid      int    `json:"aid"`      // 视频aid
	Cid      int    `json:"cid"`      // 视频cid
	Duration int    `json:"duration"` // 视频长度 (秒数)
	Type     string `json:"type"`     // 类型
}
