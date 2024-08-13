package audio

import (
	"fmt"
	"net/http"
)

// 获取音频流URL (web端)
//
// 参数：
//   - sid (int): 音频auid
//
// 备注：
//   - 本接口仅能获取192K音质的音频
//   - web端无法播放完整付费歌曲，付费歌曲为30s试听片段
func (a *Audio) GetAudioURL(sid int) (*AudioURLResponse, error) {
	baseURL := "https://www.bilibili.com/audio/music-service-c/web/url"

	formData := map[string]string{
		"sid": fmt.Sprintf("%d", sid),
	}

	resp, err := a.client.HTTPClient.R().
		SetFormData(formData).
		SetCookie(&http.Cookie{
			Name:  "SESSDATA",
			Value: a.client.SESSDATA,
		}).
		SetResult(&AudioURLResponse{}).
		Get(baseURL)

	if err != nil {
		return nil, err
	}

	return resp.Result().(*AudioURLResponse), nil
}

// 获取音频流URL（可获取付费音频）
//
// 参数：
//   - accessKey (string): APP登录Token，APP方式必要
//   - songID (int): 音频auid，必须
//   - quality (int): 音质代码，必须
//   - privilege (int): 必须为2
//   - mid (int): 当前用户mid，任意值
//   - platform (string): 平台标识，任意值
//
// 备注：
//   - 付费音乐需要有带大会员或音乐包的账号登录，否则为试听片段
//   - 无损音质需要登录的用户为会员
func (a *Audio) GetPaidAudioURL(accessKey string, songID int, quality int, privilege int, mid int, platform string) (*PaidAudioURLResponse, error) {
	baseURL := "https://api.bilibili.com/audio/music-service-c/url"

	formData := map[string]string{
		"access_key": accessKey,
		"songid":     fmt.Sprintf("%d", songID),
		"quality":    fmt.Sprintf("%d", quality),
		"privilege":  fmt.Sprintf("%d", privilege),
		"mid":        fmt.Sprintf("%d", mid),
		"platform":   platform,
	}

	resp, err := a.client.HTTPClient.R().
		SetFormData(formData).
		SetCookie(&http.Cookie{
			Name:  "SESSDATA",
			Value: a.client.SESSDATA,
		}).
		SetResult(&PaidAudioURLResponse{}).
		Get(baseURL)

	if err != nil {
		return nil, err
	}

	return resp.Result().(*PaidAudioURLResponse), nil
}

// ----------------------

// AudioURLResponse represents the structure of the API response for getting the audio stream URL.
type AudioURLResponse struct {
	Code int           `json:"code"` // 返回值: 0表示成功, 7201006表示未找到或已下架, 72000000表示请求错误
	Msg  string        `json:"msg"`  // 错误信息, 默认为"success"
	Data *AudioURLData `json:"data"` // 数据本体
}

// AudioURLData represents the detailed information of the audio stream.
type AudioURLData struct {
	SID       int          `json:"sid"`       // 音频auid
	Type      int          `json:"type"`      // 音质标识: -1表示试听片段（192K），1表示192K
	Info      string       `json:"info"`      // 作用尚不明确, 默认为空
	Timeout   int          `json:"timeout"`   // 有效时长, 单位为秒, 一般为3小时
	Size      int          `json:"size"`      // 文件大小, 单位为字节, 当type为-1时size为0
	Cdns      []string     `json:"cdns"`      // 音频流URL数组
	Qualities *interface{} `json:"qualities"` // 恒为null
	Title     *interface{} `json:"title"`     // 恒为null
	Cover     *interface{} `json:"cover"`     // 恒为null
}

//---

type PaidAudioURLResponse struct {
	Code int               `json:"code"` // 返回值: 0表示成功, 7201006表示未找到或已下架, 72000000表示请求错误
	Msg  string            `json:"msg"`  // 错误信息, 默认为"success"
	Data *PaidAudioURLData `json:"data"` // 数据本体
}

// PaidAudioURLData represents the detailed information of the paid audio stream.
type PaidAudioURLData struct {
	SID       int                `json:"sid"`       // 音频auid
	Type      int                `json:"type"`      // 音质标识: -1表示试听片段（192K），0表示128K，1表示192K，2表示320K，3表示FLAC
	Info      string             `json:"info"`      // 作用尚不明确, 默认为空
	Timeout   int                `json:"timeout"`   // 有效时长, 单位为秒, 一般为3小时
	Size      int                `json:"size"`      // 文件大小, 单位为字节, 当type为-1时size为0
	Cdns      []string           `json:"cdns"`      // 音频流URL数组
	Qualities []AudioQualityInfo `json:"qualities"` // 音质列表
	Title     string             `json:"title"`     // 音频标题
	Cover     string             `json:"cover"`     // 音频封面图片url
}

// AudioQualityInfo represents the information of a specific audio quality.
type AudioQualityInfo struct {
	Type        int    `json:"type"`        // 音质代码
	Desc        string `json:"desc"`        // 音质名称
	Size        int    `json:"size"`        // 该音质的文件大小, 单位为字节
	Bps         string `json:"bps"`         // 比特率标签
	Tag         string `json:"tag"`         // 音质标签
	Require     int    `json:"require"`     // 是否需要会员权限: 0表示不需要, 1表示需要
	RequireDesc string `json:"requiredesc"` // 会员权限标签
}
