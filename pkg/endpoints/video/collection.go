package video

import (
	"fmt"
)

// 获取视频合集信息
//
// Parameters:
//   - mid (int): 用户的mid
//   - seasonID (int): 视频合集 ID
//   - pageNum (int): 页码索引, 默认为1, 可选
//   - pageSize (int): 单页内容数量, 默认为20, 可选
//   - sortReverse (bool): 排序方式, true表示升序, false表示默认排序, 可选
//   - gaiaVToken (string): 风控验证 Token, 可选
//   - webLocation (string): 页面位置, 可选
//   - wRid (string): WBI 签名, 可选
//   - wts (int): UNIX 秒级时间戳, 可选
//
// Authentication:
//   - 需要验证 referer
//   - 需要 User-Agent
func (v *Video) SeasonsArchives(mid int, seasonID int, sortReverse bool, pageNum, pageSize int, gaiaVToken, webLocation, wRid string, wts int) (*SeasonArchivesResponse, error) {
	baseURL := "https://api.bilibili.com/x/polymer/web-space/seasons_archives_list"

	defaultPageNum := 1
	defaultPageSize := 30

	if pageNum == 0 {
		pageNum = defaultPageNum
	}
	if pageSize == 0 {
		pageSize = defaultPageSize
	}

	formData := map[string]string{
		"mid":          fmt.Sprintf("%d", mid),
		"season_id":    fmt.Sprintf("%d", seasonID),
		"sort_reverse": fmt.Sprintf("%v", sortReverse),
		"page_num":     fmt.Sprintf("%d", pageNum),
		"page_size":    fmt.Sprintf("%d", pageSize),
		"gaia_vtoken":  gaiaVToken,
		"web_location": webLocation,
		"w_rid":        wRid,
		"wts":          fmt.Sprintf("%d", wts),
	}

	resp, err := v.client.HTTPClient.R().
		SetHeader("Referer", "https://www.bilibili.com").
		SetHeader("User-Agent", v.client.UserAgent).
		SetQueryParams(formData).
		SetResult(&SeasonArchivesResponse{}).Get(baseURL)
	if err != nil {
		return nil, err
	}

	return resp.Result().(*SeasonArchivesResponse), nil
}

// 只获取系列视频
//
// Parameters:
//   - mid (int): 用户的 mid
//   - pageNum (int): 页码索引, 默认为1, 可选
//   - pageSize (int): 单页内容数量,默认为20, 可选 (最大为20?)
//   - gaiaVToken (string): 风控验证 Token, 可选
//   - wRid (string): WBI 签名, 可选
//   - wts (int): UNIX 秒级时间戳, 可选
//
// Authentication:
//   - 需要验证 referer
//   - 需要 User-Agent
func (v *Video) SeasonsSeries(mid int, pageNum, pageSize int, gaiaVToken, wRid string, wts int) (*SeriesListResponse, error) {
	baseURL := "https://api.bilibili.com/x/polymer/web-space/home/seasons_series"

	defaultPageNum := 1
	defaultPageSize := 20

	if pageNum == 0 {
		pageNum = defaultPageNum
	}
	if pageSize == 0 {
		pageSize = defaultPageSize
	}

	formData := map[string]string{
		"mid":         fmt.Sprintf("%d", mid),
		"page_num":    fmt.Sprintf("%d", pageNum),
		"page_size":   fmt.Sprintf("%d", pageSize),
		"gaia_vtoken": gaiaVToken,
		"w_rid":       wRid,
		"wts":         fmt.Sprintf("%d", wts),
	}

	resp, err := v.client.HTTPClient.R().
		SetHeader("Referer", "https://www.bilibili.com").
		SetHeader("User-Agent", v.client.UserAgent).
		SetQueryParams(formData).
		SetResult(&SeriesListResponse{}).Get(baseURL)

	if err != nil {
		return nil, err
	}

	return resp.Result().(*SeriesListResponse), nil
}

// 获取系列和合集视频
//
// Parameters:
//   - mid (int): 用户的 mid
//   - pageNum (int): 页码, 默认1, 可选
//   - pageSize (int): 每页数量, 默认20, 可选
//   - wRid (string): WBI 签名, 可选
//   - wts (int): UNIX 秒级时间戳, 可选
//   - webLocation (string): 页面位置, 可选
//
// Authentication:
//   - User-Agent: 必须为正常浏览器
func (v *Video) SeasonsSeriesList(mid int, pageNum, pageSize int, wRid string, wts int, webLocation string) (*SeasonsSeriesListResponse, error) {
	baseURL := "https://api.bilibili.com/x/polymer/web-space/seasons_series_list"

	defaultPageNum := 1
	defaultPageSize := 20

	if pageNum == 0 {
		pageNum = defaultPageNum
	}
	if pageSize == 0 {
		pageSize = defaultPageSize
	}

	formData := map[string]string{
		"mid":          fmt.Sprintf("%d", mid),
		"page_num":     fmt.Sprintf("%d", pageNum),
		"page_size":    fmt.Sprintf("%d", pageSize),
		"w_rid":        wRid,
		"wts":          fmt.Sprintf("%d", wts),
		"web_location": webLocation,
	}

	resp, err := v.client.HTTPClient.R().
		SetHeader("Referer", "https://www.bilibili.com").
		SetHeader("User-Agent", v.client.UserAgent).
		SetQueryParams(formData).
		SetResult(&SeasonsSeriesListResponse{}).Get(baseURL)
	if err != nil {
		return nil, err
	}

	return resp.Result().(*SeasonsSeriesListResponse), nil
}

// 查询指定系列
//
// Parameters:
//   - seriesID (int): 系列 ID
//
// Authentication:
//   - 无需特殊认证
func (v *Video) Series(seriesID int) (*SeriesResponse, error) {
	baseURL := "https://api.bilibili.com/x/series/series"

	formData := map[string]string{
		"series_id": fmt.Sprintf("%d", seriesID),
	}

	resp, err := v.client.HTTPClient.R().
		SetQueryParams(formData).
		SetResult(&SeriesResponse{}).Get(baseURL)

	if err != nil {
		return nil, err
	}

	return resp.Result().(*SeriesResponse), nil
}

// 获取指定系列视频
//
// Parameters:
//   - mid (int): 用户的 mid
//   - seriesID (int): 系列 ID
//   - sort (string): 排序方式, 可选值为 "desc" 或 "asc"
//   - pn (int): 页码, 默认为 1
//   - ps (int): 每页数量, 默认为 20
//   - currentMid (int): 当前用户 mid, 用于播放进度追踪
//
// Authentication:
//   - 无需特殊认证
func (v *Video) Archives(mid, seriesID int, sort string, pn, ps, currentMid int) (*SeriesArchivesResponse, error) {
	baseURL := "https://api.bilibili.com/x/series/archives"

	defaultPageNum := 1
	defaultPageSize := 20

	if pn == 0 {
		pn = defaultPageNum
	}
	if ps == 0 {
		ps = defaultPageSize
	}

	formData := map[string]string{
		"mid":         fmt.Sprintf("%d", mid),
		"series_id":   fmt.Sprintf("%d", seriesID),
		"only_normal": fmt.Sprintf("%t", true), // 作用尚不明确, 默认为 true
		"sort":        sort,
		"pn":          fmt.Sprintf("%d", pn),
		"ps":          fmt.Sprintf("%d", ps),
		"current_mid": fmt.Sprintf("%d", currentMid),
	}

	fmt.Printf("formData: %v\n", formData)

	req := v.client.HTTPClient.R().
		SetQueryParams(formData).
		SetResult(&SeriesArchivesResponse{})

	resp, err := req.Get(baseURL)
	if err != nil {
		return nil, err
	}

	return resp.Result().(*SeriesArchivesResponse), nil
}

// SeasonArchivesResponse represents the structure of the API response for fetching season archives.
type SeasonArchivesResponse struct {
	Code    int    `json:"code"`    // 返回值: 0表示成功
	Message string `json:"message"` // 错误信息, 默认为0
	TTL     int    `json:"ttl"`     // TTL, 固定值1
	Data    struct {
		Aids     []int          `json:"aids"`     // 稿件 avid 列表
		Archives []Archive      `json:"archives"` // 合集中的视频
		Meta     ArchiveMeta    `json:"meta"`     // 合集元数据
		Page     PaginationInfo `json:"page"`     // 分页信息
	} `json:"data"`
}

// Archive represents the structure of each archive item.
type Archive struct {
	Aid              int         `json:"aid"`               // 稿件 avid
	Bvid             string      `json:"bvid"`              // 稿件 bvid
	Ctime            int         `json:"ctime"`             // 创建时间 Unix 时间戳
	Duration         int         `json:"duration"`          // 视频时长 单位为秒
	EnableVt         bool        `json:"enable_vt"`         // 是否支持互动视频
	InteractiveVideo bool        `json:"interactive_video"` // 是否是互动视频
	Pic              string      `json:"pic"`               // 封面 URL
	PlaybackPosition int         `json:"playback_position"` // 播放位置 单位为 %
	Pubdate          int         `json:"pubdate"`           // 发布日期 Unix 时间戳
	Stat             ArchiveStat `json:"stat"`              // 稿件信息
	State            int         `json:"state"`             // 状态值
	Title            string      `json:"title"`             // 稿件标题
	UgcPay           int         `json:"ugc_pay"`           // UGC 付费标识
	VtDisplay        string      `json:"vt_display"`        // 旧接口无
}

type SeasonArchive struct {
	Aid              int         `json:"aid"`               // 稿件 avid
	Bvid             string      `json:"bvid"`              // 稿件 bvid
	Ctime            int         `json:"ctime"`             // 创建时间 Unix 时间戳
	Duration         int         `json:"duration"`          // 视频时长 单位为秒
	EnableVt         int         `json:"enable_vt"`         // 是否支持互动视频
	InteractiveVideo bool        `json:"interactive_video"` // 是否是互动视频
	Pic              string      `json:"pic"`               // 封面 URL
	PlaybackPosition int         `json:"playback_position"` // 播放位置 单位为 %
	Pubdate          int         `json:"pubdate"`           // 发布日期 Unix 时间戳
	Stat             ArchiveStat `json:"stat"`              // 稿件信息
	State            int         `json:"state"`             // 状态值
	Title            string      `json:"title"`             // 稿件标题
	UgcPay           int         `json:"ugc_pay"`           // UGC 付费标识
	VtDisplay        string      `json:"vt_display"`        // 旧接口无
}

// ArchiveStat represents the statistics of an archive.
type ArchiveStat struct {
	View int `json:"view"` // 稿件播放量
	Vt   int `json:"vt"`   // 旧接口无
}

// ArchiveMeta represents the metadata of the archive collection.
type ArchiveMeta struct {
	Category    int    `json:"category"`    // 分类 ID
	Cover       string `json:"cover"`       // 合集封面 URL
	Description string `json:"description"` // 合集描述
	Mid         int    `json:"mid"`         // UP 主 ID
	Name        string `json:"name"`        // 合集标题
	Ptime       int    `json:"ptime"`       // 发布时间 Unix 时间戳
	SeasonID    int    `json:"season_id"`   // 合集 ID
	Total       int    `json:"total"`       // 合集内视频数量
}

// PaginationInfo represents the pagination information.
type PaginationInfo struct {
	PageNum  int `json:"page_num"`  // 分页页码
	PageSize int `json:"page_size"` // 单页个数
	Total    int `json:"total"`     // 合集内视频总数量
}

//--

// SeriesListResponse represents the structure of the API response for fetching series lists.
type SeriesListResponse struct {
	Code    int    `json:"code"`    // 返回值: 0表示成功, -352表示请求被风控, -400表示请求错误
	Message string `json:"message"` // 错误信息, 默认为0
	TTL     int    `json:"ttl"`     // TTL, 固定值1
	Data    struct {
		ItemsLists struct {
			Page        PaginationInfo `json:"page"`         // 分页信息
			SeasonsList []interface{}  `json:"seasons_list"` // 空
			SeriesList  []Series       `json:"series_list"`  // 系列列表
		} `json:"items_lists"`
	} `json:"data"`
}

// Series represents the structure of each series item.
type Series struct {
	Archives   []Archive  `json:"archives"`    // 系列视频列表
	Meta       SeriesMeta `json:"meta"`        // 系列元数据
	RecentAids []int      `json:"recent_aids"` // 系列视频 aid 列表
}

// SeriesMeta represents the metadata of the series.
type SeriesMeta struct {
	Category     int      `json:"category"`       // 分类 ID
	Cover        string   `json:"cover"`          // 系列封面 URL
	Creator      string   `json:"creator"`        // 创建者
	Ctime        int      `json:"ctime"`          // 创建时间 Unix 时间戳
	Description  string   `json:"description"`    // 系列描述
	Keywords     []string `json:"keywords"`       // 系列关键词列表
	LastUpdateTs int      `json:"last_update_ts"` // 最近更新时间 Unix 时间戳
	Mid          int      `json:"mid"`            // UP 主 ID
	Mtime        int      `json:"mtime"`          // 修改时间 Unix 时间戳
	Name         string   `json:"name"`           // 系列标题
	RawKeywords  string   `json:"raw_keywords"`   // 原始系列关键词
	SeriesID     int      `json:"series_id"`      // 系列 ID
	State        int      `json:"state"`          // 状态值
	Total        int      `json:"total"`          // 系列视频数量
}

// -
type SeasonsSeriesListResponse struct {
	Code    int    `json:"code"`    // 返回值: 0表示成功, -352表示请求被风控, -400表示请求错误
	Message string `json:"message"` // 错误信息, 默认为0
	TTL     int    `json:"ttl"`     // TTL, 固定值1
	Data    struct {
		ItemsLists struct {
			Page        PaginationInfo `json:"page"`         // 分页信息
			SeasonsList []Series       `json:"seasons_list"` // 视频合集列表
			SeriesList  []Series       `json:"series_list"`  // 系列列表
		} `json:"items_lists"`
	} `json:"data"`
}

//-

type SeriesResponse struct {
	Code    int    `json:"code"`    // 返回值: 0表示成功
	Message string `json:"message"` // 错误信息, 默认为0
	TTL     int    `json:"ttl"`     // TTL, 固定值1
	Data    struct {
		Meta       SeriesMeta `json:"meta"`        // 系列信息
		RecentAids []int      `json:"recent_aids"` // 系列 aid 列表
	} `json:"data"`
}

// -
// SeriesArchivesResponse represents the structure of the API response for fetching videos from a specific series.
type SeriesArchivesResponse struct {
	Code    int    `json:"code"`    // 返回值: 0表示成功
	Message string `json:"message"` // 错误信息, 默认为0
	TTL     int    `json:"ttl"`     // TTL, 固定值1
	Data    struct {
		Aids     []int           `json:"aids"`     // 视频 aid 列表
		Page     PaginationInfo  `json:"page"`     // 页码信息
		Archives []SeasonArchive `json:"archives"` // 视频信息列表
	} `json:"data"`
}
