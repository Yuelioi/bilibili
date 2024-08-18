package audio

import (
	"fmt"
	"net/http"
)

// 查询歌曲基本信息
//
// 参数：
//   - sid (int): 音频auid
//
// 备注：
//   - 认证方式：Cookie（SESSDATA）
func (a *Audio) SongInfo(sid int) (*SongInfoResponse, error) {
	baseURL := "https://www.bilibili.com/audio/music-service-c/web/song/info"

	formData := map[string]string{
		"sid": fmt.Sprintf("%d", sid),
	}

	resp, err := a.client.HTTPClient.R().
		SetFormData(formData).
		SetCookie(&http.Cookie{
			Name:  "SESSDATA",
			Value: a.client.SESSDATA,
		}).
		SetResult(&SongInfoResponse{}).
		Get(baseURL)

	if err != nil {
		return nil, err
	}

	return resp.Result().(*SongInfoResponse), nil
}

// 查询歌曲TAG
//
// 参数：
//   - sid (int): 音频auid
//
// 备注：
//   - 请求方式：GET
func (a *Audio) SongTags(sid int) (*SongTagsResponse, error) {
	baseURL := "https://www.bilibili.com/audio/music-service-c/web/tag/song"

	formData := map[string]string{
		"sid": fmt.Sprintf("%d", sid),
	}

	resp, err := a.client.HTTPClient.R().
		SetFormData(formData).
		SetResult(&SongTagsResponse{}).
		Get(baseURL)

	if err != nil {
		return nil, err
	}

	return resp.Result().(*SongTagsResponse), nil
}

// 查询歌曲创作成员列表
//
// 参数：
//   - sid (int): 音频auid
//
// 备注：
//   - 请求方式：GET
func (a *Audio) SongMembers(sid int) (*SongMembersResponse, error) {
	baseURL := "https://www.bilibili.com/audio/music-service-c/web/member/song"

	formData := map[string]string{
		"sid": fmt.Sprintf("%d", sid),
	}

	resp, err := a.client.HTTPClient.R().
		SetFormData(formData).
		SetResult(&SongMembersResponse{}).
		Get(baseURL)

	if err != nil {
		return nil, err
	}

	return resp.Result().(*SongMembersResponse), nil
}

// 获取歌曲歌词
//
// 参数：
//   - sid (int): 音频auid
//
// 备注：
//   - 请求方式：GET
func (a *Audio) SongLyric(sid int) (*SongLyricResponse, error) {
	baseURL := "https://www.bilibili.com/audio/music-service-c/web/song/lyric"

	formData := map[string]string{
		"sid": fmt.Sprintf("%d", sid),
	}

	resp, err := a.client.HTTPClient.R().
		SetFormData(formData).
		SetResult(&SongLyricResponse{}).
		Get(baseURL)

	if err != nil {
		return nil, err
	}

	return resp.Result().(*SongLyricResponse), nil
}

// -----------------------------------------

type SongInfoResponse struct {
	Code int           `json:"code"` // 返回值: 0表示成功, 72000000表示参数错误, 7201006表示该音频不存在或已被下架, 72010027表示版权音乐重定向
	Msg  string        `json:"msg"`  // 错误信息, 默认为 "success"
	Data *SongInfoData `json:"data"` // 信息本体，当出错时为null
}

type SongInfoData struct {
	ID         int            `json:"id"`         // 音频auid
	UID        int            `json:"uid"`        // UP主mid
	Uname      string         `json:"uname"`      // UP主昵称
	Author     string         `json:"author"`     // 作者名
	Title      string         `json:"title"`      // 歌曲标题
	Cover      string         `json:"cover"`      // 封面图片url
	Intro      string         `json:"intro"`      // 歌曲简介
	Lyric      string         `json:"lyric"`      // lrc歌词url
	Crtype     int            `json:"crtype"`     // 作用尚不明确, 默认为1
	Duration   int            `json:"duration"`   // 歌曲时间长度, 单位为秒
	Passtime   int            `json:"passtime"`   // 歌曲发布时间, 时间戳
	Curtime    int            `json:"curtime"`    // 当前请求时间, 时间戳
	Aid        int            `json:"aid"`        // 关联稿件avid, 无为0
	Bvid       string         `json:"bvid"`       // 关联稿件bvid, 无为空
	Cid        int            `json:"cid"`        // 关联视频cid, 无为0
	Msid       int            `json:"msid"`       // 作用尚不明确, 默认为0
	Attr       int            `json:"attr"`       // 作用尚不明确, 默认为0
	Limit      int            `json:"limit"`      // 作用尚不明确, 默认为0
	ActivityID int            `json:"activityId"` // 作用尚不明确, 默认为0
	Limitdesc  string         `json:"limitdesc"`  // 作用尚不明确, 默认为空
	Ctime      interface{}    `json:"ctime"`      // 作用尚不明确, 默认为null
	Statistic  *SongStatistic `json:"statistic"`  // 歌曲的状态数
	VipInfo    *VipInfo       `json:"vipInfo"`    // UP主会员状态
	CollectIds []int          `json:"collectIds"` // 歌曲所在的收藏夹mlid, 需要登录(SESSDATA)
	CoinNum    int            `json:"coin_num"`   // 投币数
}

type SongStatistic struct {
	Sid     int `json:"sid"`     // 音频auid
	Play    int `json:"play"`    // 播放次数
	Collect int `json:"collect"` // 收藏数
	Comment int `json:"comment"` // 评论数
	Share   int `json:"share"`   // 分享数
}

type VipInfo struct {
	Type       int `json:"type"`         // 会员类型: 0表示无, 1表示月会员, 2表示年会员
	Status     int `json:"status"`       // 会员状态: 0表示无, 1表示有
	DueDate    int `json:"due_date"`     // 会员到期时间, 时间戳 毫秒
	VipPayType int `json:"vip_pay_type"` // 会员开通状态: 0表示无, 1表示有
}

// ----------------------------------------

type SongTagsResponse struct {
	Code int          `json:"code"` // 返回值: 0表示成功
	Msg  string       `json:"msg"`  // 错误信息, 默认为 "success"
	Data []SongTagObj `json:"data"` // TAG列表, 无为空
}

// SongTagObj represents an individual tag object in the song tags list.
type SongTagObj struct {
	Type    string `json:"type"`    // TAG类型, 作用尚不明确
	Subtype int    `json:"subtype"` // 子类型, 作用尚不明确
	Key     int    `json:"key"`     // TAG id, 作用尚不明确
	Info    string `json:"info"`    // TAG名
}

// ----------------------------------------
type SongMembersResponse struct {
	Code int                 `json:"code"` // 返回值: 0表示成功
	Msg  string              `json:"msg"`  // 错误信息, 默认为 "success"
	Data []SongMemberTypeObj `json:"data"` // 成员类型列表, 无为空
}

// SongMemberTypeObj represents an individual member type in the song members list.
type SongMemberTypeObj struct {
	Type int             `json:"type"` // 成员类型代码: 1表示歌手, 2表示作词, 3表示作曲等
	List []SongMemberObj `json:"list"` // 成员列表
}

// SongMemberObj represents an individual member in the members list.
type SongMemberObj struct {
	Mid      int    `json:"mid"`       // 成员的mid, 作用尚不明确
	Name     string `json:"name"`      // 成员名
	MemberID int    `json:"member_id"` // 成员id, 作用尚不明确
}

// ----------------------------------------

type SongLyricResponse struct {
	Code int    `json:"code"` // 返回值: 0表示成功
	Msg  string `json:"msg"`  // 错误信息, 默认为 "success"
	Data string `json:"data"` // 歌词信息, lrc格式, 成功时为歌词内容, 错误时为null
}
