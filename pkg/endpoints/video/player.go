package video

import "fmt"

// GetWebPlayerInfo retrieves web player information.
// Parameters:
//   - aid (int): 视频的 aid (可选)
//   - bvid (string): 视频的 bvid (可选)
//   - cid (int): 视频的 cid (必要)
//   - wRid (string): WBI 签名 (可选)
//   - wts (int64): 当前 unix 时间戳 (可选)
func (v *Video) GetWebPlayerInfo(aid int, bvid string, cid int, wRid string, wts int64) (*WebPlayerInfoResponse, error) {
	baseURL := "https://api.bilibili.com/x/player/wbi/v2"

	// Set query parameters based on provided values
	queryParams := map[string]string{
		"aid":  fmt.Sprintf("%d", aid),
		"bvid": bvid,
		"cid":  fmt.Sprintf("%d", cid),
	}
	if wRid != "" {
		queryParams["w_rid"] = wRid
	}
	if wts != 0 {
		queryParams["wts"] = fmt.Sprintf("%d", wts)
	}

	resp, err := v.client.HTTPClient.R().
		SetQueryParams(queryParams).
		SetResult(&WebPlayerInfoResponse{}).
		Get(baseURL)
	if err != nil {
		return nil, err
	}

	return resp.Result().(*WebPlayerInfoResponse), nil
}

// WebPlayerInfoResponse represents the response structure for the web player info API.
type WebPlayerInfoResponse struct {
	Code    int    `json:"code"`    // 返回值, 0: 成功, -400: 请求错误
	Message string `json:"message"` // 错误信息, 默认为 0
	TTL     int    `json:"ttl"`     // 默认值为 1
	Data    struct {
		AID        int          `json:"aid"`      // 视频 aid
		BVID       string       `json:"bvid"`     // 视频 bvid
		CID        int          `json:"cid"`      // 视频 cid
		DMMask     *DMMask      `json:"dm_mask"`  // webmask 信息 (如果没有这一项，说明这个视频没有防挡功能)
		Subtitle   *WebSubtitle `json:"subtitle"` // 字幕信息 (需要登录，不登录此项内容为 [])
		ViewPoints []struct {
			Content string `json:"content"` // 章节名
			From    int    `json:"from"`    // 开始时间, 单位为秒
			To      int    `json:"to"`      // 结束时间, 单位为秒
			Type    int    `json:"type"`    // 类型, 具体含义视实现而定
			ImgURL  string `json:"imgUrl"`  // 图片资源地址
			LogoURL string `json:"logoUrl"` // Logo 资源地址, 如果为空则为 ""
		} `json:"view_points"` // 章节看点信息
		// 其他字段略去...
	} `json:"data"` // 数据本体
}

// DMMask represents the webmask information in the web player info.
type DMMask struct {
	CID     int    `json:"cid"`      // 视频 cid
	Plat    int    `json:"plat"`     // 未知
	FPS     int    `json:"fps"`      // webmask 取样 fps
	Time    int    `json:"time"`     // 未知
	MaskURL string `json:"mask_url"` // webmask 资源 url
}

// Subtitle represents the subtitle information in the web player info.
type WebSubtitle struct {
	AllowSubmit bool   `json:"allow_submit"` // 是否允许提交字幕
	Lan         string `json:"lan"`          // 语言类型英文字母缩写
	LanDoc      string `json:"lan_doc"`      // 语言类型中文名称
	Subtitles   []struct {
		AIStatus    int    `json:"ai_status"`    // AI 状态
		AIType      int    `json:"ai_type"`      // AI 类型
		ID          int    `json:"id"`           // 字幕 ID
		IDStr       string `json:"id_str"`       // 字符串形式的字幕 ID
		IsLock      bool   `json:"is_lock"`      // 是否锁定
		Lan         string `json:"lan"`          // 语言类型英文字母缩写
		LanDoc      string `json:"lan_doc"`      // 语言类型中文名称
		SubtitleURL string `json:"subtitle_url"` // 字幕资源 URL 地址
		Type        int    `json:"type"`         // 类型
	} `json:"subtitles"` // 字幕列表
}
