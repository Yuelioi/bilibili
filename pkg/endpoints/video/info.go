package video

import (
	"fmt"
	"net/http"
)

// 获取视频信息(web端)
//
// Parameters:
//   - aid (int): 视频的aid
//   - bvid (string): 视频的bvid
//
// Authentication:
//   - 认证方式：Cookie（SESSDATA）限制游客访问的视频需要登录
//   - 鉴权方式：Wbi 签名(本api未使用)
func (v *Video) Info(aid int, bvid string) (*VideoInfoResponse, error) {

	baseURL := "https://api.bilibili.com/x/web-interface/view"
	// baseURL = "https://api.bilibili.com/x/web-interface/wbi/view"

	formData := map[string]string{
		"aid":  fmt.Sprintf("%d", aid),
		"bvid": bvid,
	}

	resp, err := v.client.HTTPClient.R().
		SetQueryParams(formData).
		SetResult(&VideoInfoResponse{}).
		SetCookie(&http.Cookie{
			Name:  "SESSDATA",
			Value: v.client.SESSDATA,
		}).Get(baseURL)
	if err != nil {
		return nil, err
	}

	return resp.Result().(*VideoInfoResponse), nil
}

// 获取视频超详细信息(web端)
//
// Parameters:
//   - aid (int): 视频的aid
//   - bvid (string): 视频的bvid
//
// Authentication:
//   - 认证方式：Cookie（SESSDATA）限制游客访问的视频需要登录
//   - 鉴权方式：Wbi 签名(本api未使用)
func (v *Video) Detail(aid int, bvid string) (*VideoDetailResponse, error) {

	baseURL := "https://api.bilibili.com/x/web-interface/view/detail"
	// baseURL = "https://api.bilibili.com/x/web-interface/wbi/view/detail"

	formData := map[string]string{
		"aid":  fmt.Sprintf("%d", aid),
		"bvid": bvid,
	}

	resp, err := v.client.HTTPClient.R().
		SetQueryParams(formData).
		SetResult(&VideoDetailResponse{}).
		SetCookie(&http.Cookie{
			Name:  "SESSDATA",
			Value: v.client.SESSDATA,
		}).Get(baseURL)
	if err != nil {
		return nil, err
	}

	return resp.Result().(*VideoDetailResponse), nil
}

// 获取视频简介
//
// Parameters:
//   - aid (int): 视频的aid
//   - bvid (string): 视频的bvid
func (v *Video) Description(aid int, bvid string) (*DescResponse, error) {

	baseURL := "https://api.bilibili.com/x/web-interface/archive/desc"

	formData := map[string]string{
		"aid":  fmt.Sprintf("%d", aid),
		"bvid": bvid,
	}

	resp, err := v.client.HTTPClient.R().
		SetQueryParams(formData).
		SetResult(&DescResponse{}).
		Get(baseURL)
	if err != nil {
		return nil, err
	}

	return resp.Result().(*DescResponse), nil
}

// 查询视频分P列表 (avid/bvid转cid)
//
// Parameters:
//   - aid (int): 视频的aid
//   - bvid (string): 视频的bvid
func (v *Video) PageList(aid int, bvid string) (*PageListResponse, error) {

	baseURL := "https://api.bilibili.com/x/player/pagelist"

	formData := map[string]string{
		"aid":  fmt.Sprintf("%d", aid),
		"bvid": bvid,
	}

	resp, err := v.client.HTTPClient.R().
		SetQueryParams(formData).
		SetResult(&PageListResponse{}).
		Get(baseURL)
	if err != nil {
		return nil, err
	}

	return resp.Result().(*PageListResponse), nil
}

// VideoInfoResponse represents the structure of the API response for video infos.
type VideoInfoResponse struct {
	Code    int       `json:"code"`    // 返回值: 0表示成功, -400表示请求错误, -403表示权限不足, -404表示无视频
	Message string    `json:"message"` // 错误信息, 默认为0
	TTL     int       `json:"ttl"`     // TTL, 固定值1
	Data    VideoData `json:"data"`
}

type VideoData struct {
	BVID               string       `json:"bvid"`                 // 稿件bvid
	AID                int          `json:"aid"`                  // 稿件avid
	Videos             int          `json:"videos"`               // 稿件分P总数, 默认为1
	TID                int          `json:"tid"`                  // 分区tid
	TName              string       `json:"tname"`                // 子分区名称
	Copyright          int          `json:"copyright"`            // 视频类型: 1表示原创, 2表示转载
	Pic                string       `json:"pic"`                  // 稿件封面图片url
	Title              string       `json:"title"`                // 稿件标题
	PubDate            int          `json:"pubdate"`              // 稿件发布时间, 秒级时间戳
	CTime              int          `json:"ctime"`                // 用户投稿时间, 秒级时间戳
	Desc               string       `json:"desc"`                 // 视频简介
	DescV2             []DescV2Item `json:"desc_v2"`              // 新版视频简介
	State              int          `json:"state"`                // 视频状态
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
	UgcSeason          UGCSeason    `json:"ugc_season"`           // 合集信息
	Staff              []Staff      `json:"staff"`                // 合作成员列表
	IsSeasonDisplay    bool         `json:"is_season_display"`    // 是否为季显示
	UserGarb           interface{}  `json:"user_garb"`            // 用户装扮信息
	HonorReply         HonorReply   `json:"honor_reply"`          // 荣誉回复信息
	LikeIcon           string       `json:"like_icon"`            // 点赞图标
	ArgueInfo          ArgueInfo    `json:"argue_info"`           // 争议/警告信息
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
	Favorite   int    `json:"favorite"`   // 收藏数(仅 data>stat)
	Fav        int    `json:"fav"`        // 收藏数(仅 data>ugc_season>episodes>stat)
	Coin       int    `json:"coin"`       // 投币数
	Share      int    `json:"share"`      // 分享数
	NowRank    int    `json:"now_rank"`   // 当前排名
	HisRank    int    `json:"his_rank"`   // 历史最高排行
	Like       int    `json:"like"`       // 获赞数
	Dislike    int    `json:"dislike"`    // 点踩数
	Evaluation string `json:"evaluation"` // 视频评分
	VT         int    `json:"vt"`         // 作用尚不明确，恒为0
	VV         int    `json:"vv"`
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

type UGCSeason struct {
	ID        int       `json:"id"`         // 作品的唯一标识符
	Title     string    `json:"title"`      // 作品的标题
	Cover     string    `json:"cover"`      // 作品封面图片的URL
	MID       int       `json:"mid"`        // 作品作者的唯一标识符
	Intro     string    `json:"intro"`      // 作品的简介
	SignState int       `json:"sign_state"` // 作品的签约状态
	Attribute int       `json:"attribute"`  // 作品的属性标识符
	Sections  []Section `json:"sections"`   // 作品的章节列表
}

type Section struct {
	SeasonID int       `json:"season_id"` // 章节所属的作品ID
	ID       int       `json:"id"`        // 章节的唯一标识符
	Title    string    `json:"title"`     // 章节的标题
	Type     int       `json:"type"`      // 章节的类型标识符
	Episodes []Episode `json:"episodes"`  // 章节中的剧集列表
}

type Episode struct {
	SeasonID  int    `json:"season_id"`  // 剧集所属的作品ID
	SectionID int    `json:"section_id"` // 剧集所属的章节ID
	ID        int    `json:"id"`         // 剧集的唯一标识符
	AID       int    `json:"aid"`        // 视频的唯一标识符
	CID       int    `json:"cid"`        // 视频内容的唯一标识符
	Title     string `json:"title"`      // 剧集的标题
	Attribute int    `json:"attribute"`  // 剧集的属性标识符
	Arc       Arc    `json:"arc"`        // 视频的详细信息
	Page      Page   `json:"page"`       // 视频的分页信息
	BVID      string `json:"bvid"`       // 视频的BVID（B站视频唯一标识符）
}

type Arc struct {
	AID       int       `json:"aid"`       // 视频的唯一标识符
	Videos    int       `json:"videos"`    // 视频的数量
	TypeID    int       `json:"type_id"`   // 视频的类型ID
	TypeName  string    `json:"type_name"` // 视频的类型名称
	Copyright int       `json:"copyright"` // 视频的版权类型
	Pic       string    `json:"pic"`       // 视频封面的URL
	Title     string    `json:"title"`     // 视频的标题
	Pubdate   int       `json:"pubdate"`   // 视频的发布时间
	Ctime     int       `json:"ctime"`     // 视频创建的时间
	Desc      string    `json:"desc"`      // 视频的描述
	State     int       `json:"state"`     // 视频的状态
	Duration  int       `json:"duration"`  // 视频的时长，单位为秒
	Rights    Rights    `json:"rights"`    // 视频的权限信息
	Author    Author    `json:"author"`    // 视频的作者信息
	Stat      Stat      `json:"stat"`      // 视频的统计数据
	Dynamic   string    `json:"dynamic"`   // 视频的动态内容
	Dimension Dimension `json:"dimension"` // 视频的维度信息（分辨率等）
}

type Author struct {
	MID  int    `json:"mid"`  // 作者的唯一标识符
	Name string `json:"name"` // 作者的名称
	Face string `json:"face"` // 作者头像的URL
}

type VideoDetailResponse struct {
	Code    int    `json:"code"`    // 返回值，0：成功，-400：请求错误，-403：权限不足，-404：无视频，62002：稿件不可见，62004：稿件审核中
	Message string `json:"message"` // 错误信息，默认为0
	TTL     int    `json:"ttl"`     // 时间戳，默认为1
	Data    struct {
		View      VideoData   `json:"view"`       // 视频基本信息
		Card      Card        `json:"card"`       // 视频 UP 主信息
		Tags      []Tag       `json:"tags"`       // 视频 TAG 信息
		Reply     Reply       `json:"reply"`      // 视频热评信息
		Related   []Related   `json:"related"`    // 推荐视频信息
		Spec      interface{} `json:"spec"`       // 作用尚不明确
		HotShare  HotShare    `json:"hot_share"`  // 作用尚不明确
		Elec      interface{} `json:"elec"`       // 作用尚不明确
		Recommend interface{} `json:"recommend"`  // 作用尚不明确
		ViewAddit ViewAddit   `json:"view_addit"` // 作用尚不明确

	} `json:"data"` // 信息本体
}

// Card 表示视频 UP 主信息
type Card struct {
	Card         CardInfo `json:"card"`          // UP 主名片信息
	Space        Space    `json:"space"`         // 主页头图
	Following    bool     `json:"following"`     // 是否关注此用户，true：已关注，false：未关注
	ArchiveCount int      `json:"archive_count"` // 用户稿件数
	ArticleCount int      `json:"article_count"` // 用户专栏数
	Follower     int      `json:"follower"`      // 粉丝数
	LikeNum      int      `json:"like_num"`      // UP 主获赞次数
}

// CardInfo 表示 Card 中的 card 对象
type CardInfo struct {
	Mid            string         `json:"mid"`              // 用户 mid
	Name           string         `json:"name"`             // 用户昵称
	Approve        bool           `json:"approve"`          // 作用尚不明确，默认为 false
	Sex            string         `json:"sex"`              // 用户性别，男、女、保密
	Rank           string         `json:"rank"`             // 作用尚不明确，默认为 10000
	Face           string         `json:"face"`             // 用户头像链接
	FaceNft        int            `json:"face_nft"`         // 是否为 nft 头像，0 不是，1 是
	DisplayRank    string         `json:"DisplayRank"`      // 作用尚不明确，默认为 0
	Regtime        int            `json:"regtime"`          // 作用尚不明确，默认为 0
	Spacesta       int            `json:"spacesta"`         // 作用尚不明确，默认为 0
	Birthday       string         `json:"birthday"`         // 作用尚不明确，默认为空
	Place          string         `json:"place"`            // 作用尚不明确，默认为空
	Description    string         `json:"description"`      // 作用尚不明确，默认为空
	Article        int            `json:"article"`          // 作用尚不明确，默认为 0
	Attentions     []interface{}  `json:"attentions"`       // 作用尚不明确，默认为空
	Fans           int            `json:"fans"`             // 粉丝数
	Friend         int            `json:"friend"`           // 关注数
	Attention      int            `json:"attention"`        // 关注数
	Sign           string         `json:"sign"`             // 签名
	LevelInfo      LevelInfo      `json:"level_info"`       // 等级
	Pendant        Pendant        `json:"pendant"`          // 挂件
	Nameplate      Nameplate      `json:"nameplate"`        // 勋章
	Official       Official       `json:"Official"`         // 认证信息
	OfficialVerify OfficialVerify `json:"official_verify"`  // 认证信息2
	Vip            Vip            `json:"vip"`              // 大会员状态
	IsSeniorMember int            `json:"is_senior_member"` // 是否为硬核会员，0：否，1：是
}

// LevelInfo 表示 Card 中的 level_info 对象
type LevelInfo struct {
	CurrentLevel int `json:"current_level"` // 当前等级，0-6级
	CurrentMin   int `json:"current_min"`   // 作用尚不明确，默认为 0
	CurrentExp   int `json:"current_exp"`   // 作用尚不明确，默认为 0
	NextExp      int `json:"next_exp"`      // 作用尚不明确，默认为 0
}

// Pendant 表示 Card 中的 pendant 对象
type Pendant struct {
	Pid    int    `json:"pid"`    // 挂件 id
	Name   string `json:"name"`   // 挂件名称
	Image  string `json:"image"`  // 挂件图片 url
	Expire int    `json:"expire"` // 作用尚不明确，默认为 0
}

// Nameplate 表示 Card 中的 nameplate 对象
type Nameplate struct {
	Nid        int    `json:"nid"`         // 勋章 id，详细说明有待补充
	Name       string `json:"name"`        // 勋章名称
	Image      string `json:"image"`       // 挂件图片 url 正常
	ImageSmall string `json:"image_small"` // 勋章图片 url 小
	Level      string `json:"level"`       // 勋章等级
	Condition  string `json:"condition"`   // 勋章条件
}

// Official 表示 Card 中的 Official 对象
type Official struct {
	Role  int    `json:"role"`  // 认证类型，见用户认证类型一览
	Title string `json:"title"` // 认证信息，无为空
	Desc  string `json:"desc"`  // 认证备注，无为空
	Type  int    `json:"type"`  // 是否认证，-1：无，0：认证
}

// OfficialVerify 表示 Card 中的 official_verify 对象
type OfficialVerify struct {
	Type int    `json:"type"` // 是否认证，-1：无，0：认证
	Desc string `json:"desc"` // 认证信息，无为空
}

// Vip 表示 Card 中的 vip 对象
type Vip struct {
	Type               int    `json:"type"`                 // 会员类型，0：无，1：月大会员，2：年度及以上大会员
	Status             int    `json:"status"`               // 会员状态，0：无，1：有
	DueDate            int    `json:"due_date"`             // 会员过期时间，Unix 时间戳（毫秒）
	VipPayType         int    `json:"vip_pay_type"`         // 支付类型，0：未支付，1：已支付
	ThemeType          int    `json:"theme_type"`           // 作用尚不明确，默认为 0
	Label              Label  `json:"label"`                // 会员标签
	AvatarSubscript    int    `json:"avatar_subscript"`     // 是否显示会员图标，0：不显示，1：显示
	NicknameColor      string `json:"nickname_color"`       // 会员昵称颜色，颜色码
	Role               int    `json:"role"`                 // 大角色类型，1：月度大会员，3：年度大会员，7：十年大会员，15：百年大会员
	AvatarSubscriptURL string `json:"avatar_subscript_url"` // 大会员角标地址
	TVVipStatus        int    `json:"tv_vip_status"`        // 电视大会员状态，0：未开通
	TVVipPayType       int    `json:"tv_vip_pay_type"`      // 电视大会员支付类型
}

// Label 表示 Vip 中的 label 对象
type Label struct {
	Path                  string `json:"path"`                      // 作用尚不明确，默认为空
	Text                  string `json:"text"`                      // 会员类型文案
	LabelTheme            string `json:"label_theme"`               // 会员标签
	TextColor             string `json:"text_color"`                // 会员标签文本颜色
	BgStyle               int    `json:"bg_style"`                  // 背景样式
	BgColor               string `json:"bg_color"`                  // 会员标签背景颜色
	BorderColor           string `json:"border_color"`              // 会员标签边框颜色，未使用
	UseImgLabel           bool   `json:"use_img_label"`             // 是否使用图片标签
	ImgLabelUriHans       string `json:"img_label_uri_hans"`        // 简体版图片标签 URI
	ImgLabelUriHant       string `json:"img_label_uri_hant"`        // 繁体版图片标签 URI
	ImgLabelUriHansStatic string `json:"img_label_uri_hans_static"` // 简体版图片标签静态 URI
	ImgLabelUriHantStatic string `json:"img_label_uri_hant_static"` // 繁体版图片标签静态 URI
}

// Space 表示 Card 中的 space 对象
type Space struct {
	SImg string `json:"s_img"` // 主页头图 url 小图
	LImg string `json:"l_img"` // 主页头图 url 正常
}

// TODO 视频 TAG 信息
type Tag struct {
	// 根据具体字段定义
}

// TODO Reply 表示视频热评信息
type Reply struct {
	// 根据具体字段定义
}

// TODO Reply 表示视频热评信息
type Related struct {
	Aid         int         `json:"aid"`           // 视频AV号
	Videos      int         `json:"videos"`        // 视频集数
	Tid         int         `json:"tid"`           // 分区ID
	Tname       string      `json:"tname"`         // 分区名称
	Copyright   int         `json:"copyright"`     // 版权标识
	Pic         string      `json:"pic"`           // 视频封面图片地址
	Title       string      `json:"title"`         // 视频标题
	Pubdate     int         `json:"pubdate"`       // 发布时间，Unix时间戳
	Ctime       int         `json:"ctime"`         // 创建时间，Unix时间戳
	Desc        string      `json:"desc"`          // 视频描述
	State       int         `json:"state"`         // 状态
	Duration    int         `json:"duration"`      // 视频时长，单位秒
	MissionID   int         `json:"mission_id"`    // 任务ID
	Rights      Rights      `json:"rights"`        // 视频权限信息
	Owner       Owner       `json:"owner"`         // 视频UP主信息
	Stat        Stat        `json:"stat"`          // 视频统计信息
	Dynamic     string      `json:"dynamic"`       // 动态信息
	Cid         int         `json:"cid"`           // 视频CID
	Dimension   Dimension   `json:"dimension"`     // 视频维度信息
	SeasonID    int         `json:"season_id"`     // 剧集ID
	ShortLinkV2 string      `json:"short_link_v2"` // 短链接
	FirstFrame  string      `json:"first_frame"`   // 视频第一帧图片地址
	PubLocation string      `json:"pub_location"`  // 发布位置
	Cover43     string      `json:"cover43"`       // 4:3比例封面图片地址
	Bvid        string      `json:"bvid"`          // 视频BV号
	SeasonType  int         `json:"season_type"`   // 剧集类型
	IsOgv       bool        `json:"is_ogv"`        // 是否为OGV内容
	OgvInfo     interface{} `json:"ogv_info"`      // OGV信息，可能为空
	RcmdReason  string      `json:"rcmd_reason"`   // 推荐原因
	EnableVt    int         `json:"enable_vt"`     // 是否启用VT
	AiRcmd      interface{} `json:"ai_rcmd"`       // AI推荐信息，可能为空
}

// HotShare 表示 hot_share 对象
type HotShare struct {
	Show bool          `json:"show"` // 作用尚不明确，默认为 false
	List []interface{} `json:"list"` // 作用尚不明确，默认为空
}

// ViewAddit 表示 view_addit 对象
type ViewAddit struct {
	Field63 bool `json:"63"` // 作用尚不明确
	Field64 bool `json:"64"` // 作用尚不明确
	Field69 bool `json:"69"` // 作用尚不明确
	Field71 bool `json:"71"` // 作用尚不明确
	Field72 bool `json:"72"` // 作用尚不明确
}

type DescResponse struct {
	Code    int    `json:"code"`    // 返回值，0：成功
	Message string `json:"message"` // 错误信息，默认为0
	TTL     int    `json:"ttl"`     // 时间戳，默认为1
	Data    string `json:"data"`    // 视频简介
}

// Response 表示根对象，包含了返回码、错误信息、TTL以及数据列表
type PageListResponse struct {
	Code    int         `json:"code"`    // 返回值，0：成功，-400：请求错误，-404：无视频
	Message string      `json:"message"` // 错误信息，默认为0
	Ttl     int         `json:"ttl"`     // TTL，默认为1
	Data    []VideoPart `json:"data"`    // 分P列表
}

// VideoPart 表示data数组中的对象，包含了分P的详细信息
type VideoPart struct {
	Cid        int       `json:"cid"`         // 当前分P cid
	Page       int       `json:"page"`        // 当前分P页码
	From       string    `json:"from"`        // 视频来源，vupload：普通上传（B站），hunan：芒果TV，qq：腾讯
	Part       string    `json:"part"`        // 当前分P标题
	Duration   int       `json:"duration"`    // 当前分P持续时间，单位为秒
	Vid        string    `json:"vid"`         // 站外视频vid
	Weblink    string    `json:"weblink"`     // 站外视频跳转url
	Dimension  Dimension `json:"dimension"`   // 当前分P分辨率
	FirstFrame string    `json:"first_frame"` // 分P封面
}
