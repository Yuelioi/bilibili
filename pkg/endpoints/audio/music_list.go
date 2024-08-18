package audio

import (
	"fmt"
	"net/http"
)

// 查询自己创建的歌单
//
// 参数：
//   - pn (int): 页码
//   - ps (int): 每页项数
//
// 备注：
//   - 请求方式：GET
//   - 认证方式：Cookie（SESSDATA）
func (a *Audio) CreatedCollections(pn, ps int) (*CreatedCollectionsResponse, error) {
	baseURL := "https://www.bilibili.com/audio/music-service-c/web/collections/list"

	formData := map[string]string{
		"pn": fmt.Sprintf("%d", pn),
		"ps": fmt.Sprintf("%d", ps),
	}

	resp, err := a.client.HTTPClient.R().
		SetFormData(formData).
		SetCookies([]*http.Cookie{
			{
				Name:  "SESSDATA",
				Value: a.client.SESSDATA},
			{
				Name:  "DedeUserID",
				Value: fmt.Sprint(a.client.DedeUserID)}}).
		SetResult(&CreatedCollectionsResponse{}).
		Get(baseURL)

	if err != nil {
		return nil, err
	}

	return resp.Result().(*CreatedCollectionsResponse), nil
}

// 查询音频收藏夹（默认歌单）信息
//
// 参数：
//   - sid (int): 音频收藏夹mlid，必须为默认收藏夹mlid
//
// 备注：
//   - 认证方式：Cookie（SESSDATA）
//   - 鉴权方式：Cookie中DedeUserID存在且不为0
func (a *Audio) CollectionInfo(sid int) (*CollectionInfoResponse, error) {
	baseURL := "https://www.bilibili.com/audio/music-service-c/web/collections/info"

	formData := map[string]string{
		"sid": fmt.Sprintf("%d", sid),
	}

	resp, err := a.client.HTTPClient.R().
		SetFormData(formData).
		SetCookies([]*http.Cookie{
			{
				Name:  "SESSDATA",
				Value: a.client.SESSDATA},
			{
				Name:  "DedeUserID",
				Value: fmt.Sprint(a.client.DedeUserID)}}).
		SetResult(&CollectionInfoResponse{}).
		Get(baseURL)

	if err != nil {
		return nil, err
	}

	return resp.Result().(*CollectionInfoResponse), nil
}

// 查询热门歌单
//
// 参数：
//   - pn (int): 页码
//   - ps (int): 每页项数
//
// 备注：
//   - 认证方式：Cookie（SESSDATA）
func (a *Audio) HotPlaylists(pn, ps int) (*HotPlaylistsResponse, error) {
	baseURL := "https://www.bilibili.com/audio/music-service-c/web/menu/hit"

	formData := map[string]string{
		"pn": fmt.Sprintf("%d", pn),
		"ps": fmt.Sprintf("%d", ps),
	}

	resp, err := a.client.HTTPClient.R().
		SetFormData(formData).
		SetCookie(&http.Cookie{
			Name:  "SESSDATA",
			Value: a.client.SESSDATA,
		}).
		SetResult(&HotPlaylistsResponse{}).
		Get(baseURL)

	if err != nil {
		return nil, err
	}

	return resp.Result().(*HotPlaylistsResponse), nil
}

// 查询热门榜单
//
// 参数：
//   - pn (int): 页码
//   - ps (int): 每页项数
//
// 备注：
//   - 认证方式：Cookie（SESSDATA）
func (a *Audio) HotRank(pn, ps int) (*HotRankResponse, error) {
	baseURL := "https://www.bilibili.com/audio/music-service-c/web/menu/rank"

	formData := map[string]string{
		"pn": fmt.Sprintf("%d", pn),
		"ps": fmt.Sprintf("%d", ps),
	}

	resp, err := a.client.HTTPClient.R().
		SetFormData(formData).
		SetCookie(&http.Cookie{
			Name:  "SESSDATA",
			Value: a.client.SESSDATA,
		}).
		SetResult(&HotRankResponse{}).
		Get(baseURL)

	if err != nil {
		return nil, err
	}

	return resp.Result().(*HotRankResponse), nil
}

// -----------------------

// CreatedCollectionsResponse represents the structure of the API response for querying created collections.
type CreatedCollectionsResponse struct {
	Code int                    `json:"code"` // 返回值: 0表示成功, 72000000表示参数错误, 72010002表示未登录
	Msg  string                 `json:"msg"`  // 错误信息, 默认为 "success"
	Data *CreatedCollectionsObj `json:"data"` // 信息本体, 正确时为对象, 错误时为null
}

// CreatedCollectionsObj represents the main data object in the response.
type CreatedCollectionsObj struct {
	CurPage   int                 `json:"curPage"`   // 当前页码
	PageCount int                 `json:"pageCount"` // 总计页数
	TotalSize int                 `json:"totalSize"` // 总计收藏夹数
	PageSize  int                 `json:"pageSize"`  // 当前页面项数
	Data      []CollectionItemObj `json:"data"`      // 歌单列表
}

// CollectionItemObj represents an individual collection item.
type CollectionItemObj struct {
	ID        int               `json:"id"`        // 音频收藏夹mlid
	UID       int               `json:"uid"`       // 创建用户mid
	Uname     string            `json:"uname"`     // 创建用户昵称
	Title     string            `json:"title"`     // 歌单标题
	Type      int               `json:"type"`      // 收藏夹属性: 0表示普通收藏夹, 1表示默认收藏夹
	Published int               `json:"published"` // 是否公开: 0表示不公开, 1表示公开
	Cover     string            `json:"cover"`     // 歌单封面图片url
	Ctime     int               `json:"ctime"`     // 歌单创建时间, 时间戳
	Song      int               `json:"song"`      // 歌单中的音乐数量
	Desc      string            `json:"desc"`      // 歌单备注信息
	Sids      []int             `json:"sids"`      // 歌单中的音乐列表
	MenuID    int               `json:"menuId"`    // 音频收藏夹对应的歌单amid
	Statistic CollectionStatObj `json:"statistic"` // 歌单状态数信息
}

// CollectionStatObj represents the statistics of a collection.
type CollectionStatObj struct {
	Sid     int  `json:"sid"`     // 音频收藏夹对应的歌单amid
	Play    int  `json:"play"`    // 播放数
	Collect int  `json:"collect"` // 收藏数
	Comment *int `json:"comment"` // 评论数
	Share   int  `json:"share"`   // 分享数
}

//------

type CollectionInfoResponse struct {
	Code int                 `json:"code"` // 返回值: 0表示成功, 72000000表示参数错误, 72010002表示未登录
	Msg  string              `json:"msg"`  // 错误信息, 默认为"success"
	Data *CollectionInfoData `json:"data"` // 信息本体, 正确时为对象，错误时为null
}

// CollectionInfoData represents the detailed information of the audio collection.
type CollectionInfoData struct {
	ID        int                 `json:"id"`        // 音频收藏夹mlid
	UID       int                 `json:"uid"`       // 创建用户mid
	Uname     string              `json:"uname"`     // 创建用户昵称
	Title     string              `json:"title"`     // 默认歌单，恒为"默认歌单"
	Type      int                 `json:"type"`      // 恒为1
	Published int                 `json:"published"` // 是否公开: 0表示不公开, 1表示公开
	Cover     string              `json:"cover"`     // 歌单封面图片url
	Ctime     int                 `json:"ctime"`     // 歌单创建时间, 时间戳
	Song      int                 `json:"song"`      // 歌单中的音乐数量
	Desc      string              `json:"desc"`      // 空, 恒为空
	Sids      []int               `json:"sids"`      // 歌单中的音乐, 按照歌单顺序排列的音频auid数组
	MenuID    int                 `json:"menuId"`    // 音频收藏夹对应的歌单amid, 与普通歌单不同
	Statistic CollectionStatistic `json:"statistic"` // 歌单状态数信息
}

// CollectionStatistic represents the statistics of the audio collection.
type CollectionStatistic struct {
	SID     int  `json:"sid"`     // 音频收藏夹对应的歌单amid
	Play    int  `json:"play"`    // 播放次数, 恒为0
	Collect int  `json:"collect"` // 收藏次数, 恒为0
	Comment *int `json:"comment"` // 评论数, 恒为null
	Share   int  `json:"share"`   // 分享次数, 恒为0
}

// ---
type HotPlaylistsResponse struct {
	Code int               `json:"code"` // 返回值: 0表示成功, 72000000表示参数错误, 72010002表示未登录
	Msg  string            `json:"msg"`  // 错误信息, 默认为"success"
	Data *HotPlaylistsData `json:"data"` // 信息本体, 正确时为对象，错误时为null
}

// HotPlaylistsData represents the detailed information of the hot playlists.
type HotPlaylistsData struct {
	CurPage   int            `json:"curPage"`   // 当前页码
	PageCount int            `json:"pageCount"` // 总计页数
	TotalSize int            `json:"totalSize"` // 总计收藏夹数
	PageSize  int            `json:"pageSize"`  // 当前页面项数
	Playlists []PlaylistInfo `json:"data"`      // 歌单列表
}

// PlaylistInfo represents the details of a single playlist.
type PlaylistInfo struct {
	MenuID    int               `json:"menuId"`    // 音频收藏夹对应的歌单amid
	UID       int               `json:"uid"`       // 创建用户mid
	Uname     string            `json:"uname"`     // 创建用户昵称
	Title     string            `json:"title"`     // 歌单标题
	Cover     string            `json:"cover"`     // 歌单封面图片url
	Intro     string            `json:"intro"`     // 歌单介绍
	Type      int               `json:"type"`      // 歌单属性: 1为普通歌单, 2为置顶歌单, 5为PGC歌单
	Off       int               `json:"off"`       // 歌单是否公开: 0为公开, 1为私密
	Ctime     int               `json:"ctime"`     // 歌单创建时间, 时间戳
	Curtime   int               `json:"curtime"`   // 当前时间, 时间戳
	Statistic PlaylistStatistic `json:"statistic"` // 歌单状态数信息
	Snum      int               `json:"snum"`      // 歌单包含歌曲个数
}

// PlaylistStatistic represents the statistics of a playlist.
type PlaylistStatistic struct {
	SID     int `json:"sid"`     // 音频收藏夹对应的歌单amid
	Play    int `json:"play"`    // 播放数
	Collect int `json:"collect"` // 收藏数
	Comment int `json:"comment"` // 评论数
	Share   int `json:"share"`   // 分享数
}

// ---
type HotRankResponse struct {
	Code int          `json:"code"` // 返回值: 0表示成功, 72000000表示参数错误, 72010002表示未登录
	Msg  string       `json:"msg"`  // 错误信息, 默认为"success"
	Data *HotRankData `json:"data"` // 信息本体, 正确时为对象，错误时为null
}

// HotRankData represents the detailed information of the hot rank.
type HotRankData struct {
	CurPage   int        `json:"curPage"`   // 当前页码
	PageCount int        `json:"pageCount"` // 总计页数
	TotalSize int        `json:"totalSize"` // 总计收藏夹数
	PageSize  int        `json:"pageSize"`  // 当前页面项数
	Playlists []RankInfo `json:"data"`      // 歌单列表
}

// RankInfo represents the details of a single rank entry.
type RankInfo struct {
	MenuID    int           `json:"menuId"`    // 音频收藏夹对应的歌单amid
	UID       int           `json:"uid"`       // 创建用户mid
	Uname     string        `json:"uname"`     // 创建用户昵称
	Title     string        `json:"title"`     // 歌单标题
	Cover     string        `json:"cover"`     // 歌单封面图片url
	Intro     string        `json:"intro"`     // 歌单介绍
	Type      int           `json:"type"`      // 歌单属性: 1为普通歌单, 2为置顶歌单, 5为PGC歌单
	Off       int           `json:"off"`       // 歌单是否公开: 0为公开, 1为私密
	Ctime     int           `json:"ctime"`     // 歌单创建时间, 时间戳
	Curtime   int           `json:"curtime"`   // 当前时间, 时间戳
	Statistic RankStatistic `json:"statistic"` // 歌单状态数信息
	Snum      int           `json:"snum"`      // 歌单包含歌曲个数
	Audios    []AudioInfo   `json:"audios"`    // 歌单中的音乐信息(部分)
}

// RankStatistic represents the statistics of a rank entry.
type RankStatistic struct {
	SID     int `json:"sid"`     // 音频收藏夹对应的歌单amid
	Play    int `json:"play"`    // 收藏数
	Collect int `json:"collect"` // 点赞数
	Comment int `json:"comment"` // 评论数
	Share   int `json:"share"`   // 分享数
}

// AudioInfo represents the information of a single audio item within a rank entry.
type AudioInfo struct {
	ID       int    `json:"id"`       // 音频id
	Title    string `json:"title"`    // 音频标题
	Duration int    `json:"duration"` // 音频时长, 单位：秒(s)
}
