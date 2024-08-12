package video

import (
	"bilibili/tools"
	"fmt"
	"net/http"
)

// GetVideoDetail fetches the detailed information of a video by either AID or BVID.
//
// Parameters:
//   - aid (int): 视频的aid
//   - bvid (string): 视频的bvid
//
// Authentication:
//   - 认证方式：Cookie（SESSDATA）
func (a *Video) GetVideoDetail(aid int, bvid string) (*VideoDetailResponse, error) {
	var baseURL string
	if aid != 0 {
		baseURL = "https://api.bilibili.com/x/web-interface/view"
	} else if bvid != "" {
		baseURL = "https://api.bilibili.com/x/web-interface/wbi/view"
	} else {
		return nil, fmt.Errorf("either aid or bvid must be provided")
	}

	params := map[string]string{}
	if aid != 0 {
		params["aid"] = fmt.Sprintf("%d", aid)
	}
	if bvid != "" {
		params["bvid"] = bvid
	}

	fullURL := tools.URLWithParams(baseURL, params)

	req := a.client.HTTPClient.R().
		SetHeader("Content-Type", "application/json").
		SetResult(&VideoDetailResponse{}).
		SetCookie(&http.Cookie{
			Name:  "SESSDATA",
			Value: a.client.SESSDATA,
		}).
		SetHeader("Referer", "https://www.bilibili.com")

	resp, err := req.Get(fullURL)
	if err != nil {
		return nil, err
	}

	return resp.Result().(*VideoDetailResponse), nil
}

// VideoDetailResponse represents the structure of the API response for video details.
type VideoDetailResponse struct {
	Code    int    `json:"code"`    // 返回值: 0表示成功, -400表示请求错误, -403表示权限不足, -404表示无视频
	Message string `json:"message"` // 错误信息, 默认为0
	TTL     int    `json:"ttl"`     // TTL, 固定值1
	Data    struct {
		BVID               string       `json:"bvid"`                 // 稿件bvid
		AID                int          `json:"aid"`                  // 稿件avid
		Videos             int          `json:"videos"`               // 稿件分P总数, 默认为1
		TID                int          `json:"tid"`                  // 分区tid
		TName              string       `json:"tname"`                // 子分区名称
		Copyright          int          `json:"copyright"`            // 视频类型: 1表示原创, 2表示转载
		Pic                string       `json:"pic"`                  // 稿件封面图片url
		Title              string       `json:"title"`                // 稿件标题
		PubDate            int64        `json:"pubdate"`              // 稿件发布时间, 秒级时间戳
		CTime              int64        `json:"ctime"`                // 用户投稿时间, 秒级时间戳
		Desc               string       `json:"desc"`                 // 视频简介
		DescV2             []DescV2Item `json:"desc_v2"`              // 新版视频简介
		State              int          `json:"state"`                // 视频状态
		Attribute          int          `json:"attribute"`            // 已弃用
		Duration           int          `json:"duration"`             // 稿件总时长(所有分P), 单位为秒
		Forward            int          `json:"forward"`              // 撞车视频跳转avid
		MissionID          int          `json:"mission_id"`           // 稿件参与的活动id
		RedirectURL        string       `json:"redirect_url"`         // 重定向url，仅番剧或影视视频存在此字段
		Rights             Rights       `json:"rights"`               // 视频属性标志
		Owner              Owner        `json:"owner"`                // 视频UP主信息
		Stat               Stat         `json:"stat"`                 // 视频状态数
		Dynamic            string       `json:"dynamic"`              // 视频同步发布的动态的文字内容
		CID                int          `json:"cid"`                  // 视频1P cid
		Dimension          Dimension    `json:"dimension"`            // 视频1P分辨率
		Premiere           interface{}  `json:"premiere"`             // null
		TeenageMode        int          `json:"teenage_mode"`         // 未成年人模式
		IsChargeableSeason bool         `json:"is_chargeable_season"` // 是否为付费季
		IsStory            bool         `json:"is_story"`             // 是否为故事片
		NoCache            bool         `json:"no_cache"`             // 作用尚不明确
		Pages              []Page       `json:"pages"`                // 视频分P列表
		Subtitle           Subtitle     `json:"subtitle"`             // 视频CC字幕信息
		UgcSesson          string       // TODO
		Staff              []Staff      `json:"staff"`             // 合作成员列表
		IsSeasonDisplay    bool         `json:"is_season_display"` // 是否为季显示
		UserGarb           interface{}  `json:"user_garb"`         // 用户装扮信息
		HonorReply         HonorReply   `json:"honor_reply"`       // 荣誉回复信息
		LikeIcon           string       `json:"like_icon"`         // 点赞图标
		ArgueInfo          ArgueInfo    `json:"argue_info"`        // 争议/警告信息
	} `json:"data"`
}

// DescV2Item represents the structure of a description item in the desc_v2 array.
type DescV2Item struct {
	RawText string `json:"raw_text"` // 简介内容
	Type    int    `json:"type"`     // 类型: 1表示普通, 2表示@他人
	BizID   int    `json:"biz_id"`   // 被@用户的mid
}

// Rights represents the video rights information.
type Rights struct {
	BP            int `json:"bp"`              // 是否允许承包
	Elec          int `json:"elec"`            // 是否支持充电
	Download      int `json:"download"`        // 是否允许下载
	Movie         int `json:"movie"`           // 是否电影
	Pay           int `json:"pay"`             // 是否PGC付费
	HD5           int `json:"hd5"`             // 是否有高码率
	NoReprint     int `json:"no_reprint"`      // 是否显示“禁止转载”标志
	Autoplay      int `json:"autoplay"`        // 是否自动播放
	UGCPay        int `json:"ugc_pay"`         // 是否UGC付费
	IsCooperation int `json:"is_cooperation"`  // 是否为联合投稿
	UGCPayPreview int `json:"ugc_pay_preview"` // 0 作用尚不明确
	NoBackground  int `json:"no_background"`   // 0 作用尚不明确
	CleanMode     int `json:"clean_mode"`      // 0 作用尚不明确
	IsSteinGate   int `json:"is_stein_gate"`   // 是否为互动视频
	Is360         int `json:"is_360"`          // 是否为全景视频
	NoShare       int `json:"no_share"`        // 0 作用尚不明确
	ArcPay        int `json:"arc_pay"`         // 0 作用尚不明确
	FreeWatch     int `json:"free_watch"`      // 0 作用尚不明确
}

// Owner represents the owner (UP主) information.
type Owner struct {
	MID  int    `json:"mid"`  // UP主mid
	Name string `json:"name"` // UP主昵称
	Face string `json:"face"` // UP主头像
}

// Stat represents the statistics of the video.
type Stat struct {
	AID        int    `json:"aid"`        // 稿件avid
	View       int    `json:"view"`       // 播放数
	Danmaku    int    `json:"danmaku"`    // 弹幕数
	Reply      int    `json:"reply"`      // 评论数
	Favorite   int    `json:"favorite"`   // 收藏数
	Coin       int    `json:"coin"`       // 投币数
	Share      int    `json:"share"`      // 分享数
	NowRank    int    `json:"now_rank"`   // 当前排名
	HisRank    int    `json:"his_rank"`   // 历史最高排行
	Like       int    `json:"like"`       // 获赞数
	Dislike    int    `json:"dislike"`    // 点踩数
	Evaluation string `json:"evaluation"` // 视频评分
	VT         int    `json:"vt"`         // 作用尚不明确，恒为0
}

// Page represents a page of the video.
type Page struct {
	CID       int       `json:"cid"`       // 分P cid
	Page      int       `json:"page"`      // 分P序号，从1开始
	From      string    `json:"from"`      // 视频来源
	Part      string    `json:"part"`      // 分P标题
	Duration  int       `json:"duration"`  // 分P持续时间，单位为秒
	VID       string    `json:"vid"`       // 站外视频vid，仅站外视频有效
	Weblink   string    `json:"weblink"`   // 站外视频跳转url，仅站外视频有效
	Dimension Dimension `json:"dimension"` // 当前分P分辨率
}

// Dimension represents the dimension of a video page.
type Dimension struct {
	Width  int `json:"width"`  // 当前分P宽度
	Height int `json:"height"` // 当前分P高度
	Rotate int `json:"rotate"` // 是否将宽高对换: 0正常, 1对换
}

// Subtitle represents subtitle information.
type Subtitle struct {
	AllowSubmit bool           `json:"allow_submit"` // 是否允许提交字幕
	List        []SubtitleItem `json:"list"`         // 字幕列表
}

// SubtitleItem represents a subtitle item.
type SubtitleItem struct {
	ID          int            `json:"id"`           // 字幕id
	Lan         string         `json:"lan"`          // 字幕语言
	LanDoc      string         `json:"lan_doc"`      // 字幕语言名称
	IsLock      bool           `json:"is_lock"`      // 是否锁定
	AuthorMID   int            `json:"author_mid"`   // 字幕上传者mid
	SubtitleURL string         `json:"subtitle_url"` // json格式字幕文件url
	Author      SubtitleAuthor `json:"author"`       // 字幕上传者信息
}

// SubtitleAuthor represents the author of the subtitle.
type SubtitleAuthor struct {
	MID           int    `json:"mid"`             // 字幕上传者mid
	Name          string `json:"name"`            // 字幕上传者昵称
	Sex           string `json:"sex"`             // 字幕上传者性别: 男/女/保密
	Face          string `json:"face"`            // 字幕上传者头像url
	Sign          string `json:"sign"`            // 字幕上传者签名
	Rank          int    `json:"rank"`            // 作用尚不明确
	Birthday      int    `json:"birthday"`        // 生日，作用尚不明确
	IsFakeAccount int    `json:"is_fake_account"` // 是否假账号
	IsDeleted     int    `json:"is_deleted"`      // 是否已删除
}

// Staff represents a staff member of the video.
type Staff struct {
	MID        int           `json:"mid"`         // 成员mid
	Title      string        `json:"title"`       // 成员名称
	Name       string        `json:"name"`        // 成员昵称
	Face       string        `json:"face"`        // 成员头像url
	VIP        StaffVIP      `json:"vip"`         // 成员大会员状态
	Official   StaffOfficial `json:"official"`    // 成员认证信息
	Follower   int           `json:"follower"`    // 成员粉丝数
	LabelStyle int           `json:"label_style"` // 标签样式
}

// StaffVIP represents the VIP status of a staff member.
type StaffVIP struct {
	Type      int `json:"type"`       // 成员会员类型: 0无, 1月会员, 2年会员
	Status    int `json:"status"`     // 会员状态: 0无, 1有
	ThemeType int `json:"theme_type"` // 作用尚不明确
}

// StaffOfficial represents the official certification of a staff member.
type StaffOfficial struct {
	Role  int    `json:"role"`  // 成员认证级别
	Title string `json:"title"` // 成员认证名
	Desc  string `json:"desc"`  // 成员认证备注
	Type  int    `json:"type"`  // 成员认证类型: -1无, 0有
}

// HonorReply represents honor reply information.
type HonorReply struct {
	Honor []HonorItem `json:"honor"` // 荣誉信息
}

// HonorItem represents an honor item.
type HonorItem struct {
	AID                int    `json:"aid"`                  // 当前稿件aid
	Type               int    `json:"type"`                 // 荣誉类型: 1入站必刷, 2每周必看, 3全站排行榜, 4热门
	Desc               string `json:"desc"`                 // 描述
	WeeklyRecommendNum int    `json:"weekly_recommend_num"` // 每周推荐数量
}

// ArgueInfo represents argue or warning information.
type ArgueInfo struct {
	ArgueLink string `json:"argue_link"` // 作用尚不明确
	ArgueMsg  string `json:"argue_msg"`  // 警告/争议提示信息
	ArgueType int    `json:"argue_type"` // 作用尚不明确
}
