package audio

import (
	"fmt"
	"net/http"
)

// 获取音频统计信息
//
// 参数：
//   - sid (int): 音频auid，必须
//
// 备注：
//   - 唯缺投币数（音频投币数）
//   - 需要音频的唯一标识符来获取统计信息
func (a *Audio) GetSongStats(sid int) (*SongStatsResponse, error) {
	baseURL := "https://www.bilibili.com/audio/music-service-c/web/stat/song"

	formData := map[string]string{
		"sid": fmt.Sprintf("%d", sid),
	}

	resp, err := a.client.HTTPClient.R().
		SetFormData(formData).
		SetCookie(&http.Cookie{
			Name:  "SESSDATA",
			Value: a.client.SESSDATA,
		}).
		SetResult(&SongStatsResponse{}).
		Get(baseURL)

	if err != nil {
		return nil, err
	}

	return resp.Result().(*SongStatsResponse), nil
}

// SongStatsResponse represents the structure of the API response for getting song statistics.
type SongStatsResponse struct {
	Code int            `json:"code"` // 返回值: 0表示成功, 72000000表示参数错误
	Msg  string         `json:"msg"`  // 错误信息, 默认为"success"
	Data *SongStatsData `json:"data"` // 信息本体
}

// SongStatsData represents the detailed information of the song statistics.
type SongStatsData struct {
	SID     int `json:"sid"`     // 音频auid
	Play    int `json:"play"`    // 播放次数
	Collect int `json:"collect"` // 收藏数
	Comment int `json:"comment"` // 评论数
	Share   int `json:"share"`   // 分享数
}
