package audio

import (
	"bilibili/tools"
	"fmt"
	"net/http"
	"net/url"
)

// 获取音频榜单每期列表
//
// 参数：
//   - listType (int): 榜单类型，1表示热榜，2表示原创榜
//   - csrf (string): CSRF Token（位于cookie），非必要
//
// 备注：
//   - CSRF Token在cookie中，不一定需要提供
func (a *Audio) GetTopList(listType int, csrfToken string) (*TopListResponse, error) {
	baseURL := "https://api.bilibili.com/x/copyright-music-publicity/toplist/all_period"

	params := map[string]string{
		"list_type": fmt.Sprintf("%d", listType),
	}
	if csrfToken != "" {
		params["csrf"] = csrfToken
	}
	fullURL := tools.URLWithParams(baseURL, params)

	resp, err := a.client.HTTPClient.R().
		SetCookie(&http.Cookie{
			Name:  "SESSDATA",
			Value: a.client.SESSDATA,
		}).
		SetResult(&TopListResponse{}).
		Get(fullURL)

	if err != nil {
		return nil, err
	}

	return resp.Result().(*TopListResponse), nil
}

// 查询音频榜单单期信息
//
// 参数：
//   - listID (int): 榜单ID，见获取音频榜单每期列表
//   - csrf (string): CSRF Token（位于cookie），非必要
//
// 备注：
//   - CSRF Token在cookie中，不一定需要提供
func (a *Audio) GetTopListDetail(listID int, csrfToken string) (*TopListDetailResponse, error) {
	baseURL := "https://api.bilibili.com/x/copyright-music-publicity/toplist/detail"

	params := map[string]string{
		"list_id": fmt.Sprintf("%d", listID),
	}
	if csrfToken != "" {
		params["csrf"] = csrfToken
	}
	fullURL := tools.URLWithParams(baseURL, params)

	resp, err := a.client.HTTPClient.R().
		SetCookie(&http.Cookie{
			Name:  "SESSDATA",
			Value: a.client.SESSDATA,
		}).
		SetResult(&TopListDetailResponse{}).
		Get(fullURL)

	if err != nil {
		return nil, err
	}

	return resp.Result().(*TopListDetailResponse), nil
}

// 获取音频榜单单期内容
//
// 参数：
//   - listID (int): 榜单ID，见获取音频榜单每期列表
//   - csrf (string): CSRF Token（位于cookie），非必要
//
// 备注：
//   - CSRF Token在cookie中，不一定需要提供
func (a *Audio) GetTopListMusic(listID int, csrfToken string) (*TopListMusicResponse, error) {
	baseURL := "https://api.bilibili.com/x/copyright-music-publicity/toplist/music_list"

	params := map[string]string{
		"list_id": fmt.Sprintf("%d", listID),
	}
	if csrfToken != "" {
		params["csrf"] = csrfToken
	}
	fullURL := tools.URLWithParams(baseURL, params)

	resp, err := a.client.HTTPClient.R().
		SetCookie(&http.Cookie{
			Name:  "SESSDATA",
			Value: a.client.SESSDATA,
		}).
		SetResult(&TopListMusicResponse{}).
		Get(fullURL)

	if err != nil {
		return nil, err
	}

	return resp.Result().(*TopListMusicResponse), nil
}

// SubscribeOrUnsubscribeTopList订阅或退订榜单
//
// 参数：
//   - state (int): 操作代码，1表示订阅，2表示退订
//   - listID (int): 榜单ID，非必要，见获取音频榜单每期列表
//   - csrfToken (string): CSRF Token（位于cookie），Cookie方式必要
//
// 备注：
//   - 需要通过 Cookie 进行认证
func (a *Audio) SubscribeOrUnsubscribeTopList(state int, listID int, csrfToken string) (*SubscribeOrUnsubscribeResponse, error) {
	baseURL := "https://api.bilibili.com/x/copyright-music-publicity/toplist/subscribe/update"

	// 准备请求的表单数据
	data := url.Values{}
	data.Set("state", fmt.Sprintf("%d", state))
	if listID != 0 {
		data.Set("list_id", fmt.Sprintf("%d", listID))
	}
	data.Set("csrf", csrfToken)

	var result SubscribeOrUnsubscribeResponse

	// 执行 POST 请求
	resp, err := a.client.HTTPClient.R().
		SetCookie(&http.Cookie{
			Name:  "SESSDATA",
			Value: a.client.SESSDATA,
		}).
		SetBody(data).
		SetResult(&result).
		Post(baseURL)

	if err != nil {
		return nil, err
	}

	return resp.Result().(*SubscribeOrUnsubscribeResponse), nil
}

// --
type TopListResponse struct {
	Code    int          `json:"code"`    // 返回值: 0表示成功, -400表示请求错误
	Message string       `json:"message"` // 错误信息, 默认为0
	Ttl     int          `json:"ttl"`     // TTL值, 恒为1
	Data    *TopListData `json:"data"`    // 信息本体
}

// TopListData represents the detailed information of the top list data.
type TopListData struct {
	List map[string][]TopListEntry `json:"list"` // 年份索引，键为年份，值为该年中的每期信息
}

// TopListEntry represents the information of a specific top list entry.
type TopListEntry struct {
	ID          int   `json:"ID"`           // 榜单ID
	Period      int   `json:"priod"`        // 榜单期数
	PublishTime int64 `json:"publish_time"` // 发布时间, 秒时间戳
}

//--

type TopListDetailResponse struct {
	Code    int                `json:"code"`    // 返回值: 0表示成功, -400表示请求错误
	Message string             `json:"message"` // 错误信息, 默认为0
	Ttl     int                `json:"ttl"`     // TTL值, 恒为1
	Data    *TopListDetailData `json:"data"`    // 信息本体
}

// TopListDetailData represents the detailed information of a specific top list entry.
type TopListDetailData struct {
	ListenFID   int    `json:"listen_fid"`   // 畅听版歌单收藏夹原始ID
	AllFID      int    `json:"all_fid"`      // 完整版歌单收藏夹原始ID
	FavMID      int    `json:"fav_mid"`      // 绑定收藏夹用户的MID
	CoverURL    string `json:"cover_url"`    // 榜单封面URL
	IsSubscribe bool   `json:"is_subscribe"` // 是否已订阅榜单: true表示已订阅，false表示未订阅
	ListenCount int    `json:"listen_count"` // 平台有版权音频的数量
}

// -
type TopListMusicResponse struct {
	Code    int               `json:"code"`    // 返回值: 0表示成功, -400表示请求错误
	Message string            `json:"message"` // 错误信息, 默认为0
	Ttl     int               `json:"ttl"`     // TTL值, 恒为1
	Data    *TopListMusicData `json:"data"`    // 信息本体
}

// TopListMusicData represents the detailed information of the top list music content.
type TopListMusicData struct {
	List []TopListMusicEntry `json:"list"` // 内容列表
}

// TopListMusicEntry represents the information of a specific top list music entry.
type TopListMusicEntry struct {
	MusicID          string   `json:"music_id"`          // 音频 MAID, 例如 MA409252256362326366
	MusicTitle       string   `json:"music_title"`       // 音频标题
	Singer           string   `json:"singer"`            // 音频作者
	Album            string   `json:"album"`             // 音频专辑
	MV_AID           int      `json:"mv_aid"`            // 音频 MV 的 avid, 若无 MV 则为 0
	MV_BVID          string   `json:"mv_bvid"`           // 音频 MV 的 bvid
	MV_Cover         string   `json:"mv_cover"`          // 音频封面 URL
	Heat             int      `json:"heat"`              // 热度值
	Rank             int      `json:"rank"`              // 排序值, 1 为最高排序
	CanListen        bool     `json:"can_listen"`        // 平台是否有版权: true 表示有，false 表示无
	Recommendation   string   `json:"recommendation"`    // （？）
	CreationAID      int      `json:"creation_aid"`      // 关联稿件 avid
	CreationBVID     string   `json:"creation_bvid"`     // 关联稿件 bvid
	CreationCover    string   `json:"creation_cover"`    // 关联稿件封面 URL
	CreationTitle    string   `json:"creation_title"`    // 关联稿件标题
	CreationUP       int      `json:"creation_up"`       // 关联稿件 UP 主 mid
	CreationNickname string   `json:"creation_nickname"` // 关联稿件 UP 主昵称
	CreationDuration int      `json:"creation_duration"` // 关联稿件时长, 单位为秒
	CreationPlay     int      `json:"creation_play"`     // 关联稿件播放量
	CreationReason   string   `json:"creation_reason"`   // 关联稿件二级分区名
	Achievements     []string `json:"achievements"`      // 获得成就
	MaterialID       int      `json:"material_id"`       // （？）
	MaterialUseNum   int      `json:"material_use_num"`  // （？）
	MaterialDuration int      `json:"material_duration"` // （？）
	MaterialShow     int      `json:"material_show"`     // （？）
	SongType         int      `json:"song_type"`         // （？）
}

// TopListMusicEntryAchievements represents the achievements of a specific top list music entry.
type TopListMusicEntryAchievements struct {
	AchievementText string `json:"achievement"` // 成就文案
}

// --

type SubscribeOrUnsubscribeResponse struct {
	Code    int    `json:"code"`    // 返回值: 0表示成功, -101表示账号未登录, -111表示CSRF验证失败, 400表示请求错误
	Message string `json:"message"` // 错误信息, 默认为0
	Ttl     int    `json:"ttl"`     // TTL值, 恒为1
}
