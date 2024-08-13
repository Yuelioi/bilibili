package video

import (
	"fmt"
	"net/http"
)

// 点赞视频（web端）
//
// 参数：
//   - aid (int): 稿件 avid，任选其一
//   - bvid (string): 稿件 bvid，任选其一
//   - like (int): 操作方式，1表示点赞，2表示取消赞
//   - csrf (string): CSRF Token，位于 Cookie
//
// 备注：
//   - 认证方式：仅可Cookie（SESSDATA）
//   - 需验证 Cookie 中 buvid3 字段存在且正常，否则将触发风控
func (v *Video) Like(aid int, bvid string, like int) (*LikeResponse, error) {
	baseURL := "https://api.bilibili.com/x/web-interface/archive/like"

	// 检查buvid3是否存在
	if v.client.Buvid3 == "" {
		return nil, fmt.Errorf("buvid3 is required but not provided")
	}

	// 构建表单数据
	formData := map[string]string{
		"aid":  fmt.Sprintf("%d", aid),
		"bvid": bvid,
		"like": fmt.Sprintf("%d", like),
		"csrf": v.client.CSRF,
	}

	resp, err := v.client.HTTPClient.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(formData).
		SetCookies([]*http.Cookie{
			&http.Cookie{
				Name:  "SESSDATA",
				Value: v.client.SESSDATA},
			&http.Cookie{
				Name:  "buvid3",
				Value: fmt.Sprint(v.client.Buvid3)}}).
		SetResult(&LikeResponse{}).
		Post(baseURL)

	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	// 检查响应状态码
	if resp.IsError() {
		return nil, fmt.Errorf("request failed with status: %s", resp.Status())
	}
	return resp.Result().(*LikeResponse), nil
}

// 点赞视频（APP端）
//
// 参数：
//   - accessKey (string): APP 登录 Token
//   - aid (int): 稿件 avid
//   - like (int): 操作方式，0表示点赞，1表示取消赞
//
// 备注：
//   - 认证方式：仅可APP
func (v *Video) LikeApp(aid int, like int) (*AppLikeResponse, error) {
	url := "https://app.bilibili.com/x/v2/view/like"

	formData := map[string]string{
		"access_key": v.client.AccessKey,
		"aid":        fmt.Sprintf("%d", aid),
		"like":       fmt.Sprintf("%d", like),
	}

	resp, err := v.client.HTTPClient.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(formData).
		SetResult(&AppLikeResponse{}).
		Post(url)

	if err != nil {
		return nil, err
	}

	return resp.Result().(*AppLikeResponse), nil
}

// 判断视频近期是否被点赞(双端) 没有作用
//
// 参数：
//   - aid (int, optional): 稿件 avid，和 bvid 二选一
//   - bvid (string, optional): 稿件 bvid，和 aid 二选一
//   - access_key (string, optional): APP 登录 Token（APP 方式必要）
//
// 备注：
//   - 认证方式：APP 或 Cookie（SESSDATA）
//   - 该 API 仅能判断视频在近期是否被点赞，不能判断视频是否被点赞。近期的定义不明，但至少半年前点赞的视频，获取到的结果会是0。
//
// Deprecated: Use NewFunction instead.
func (v *Video) HasLike(aid int, bvid string) (*HasLikeResponse, error) {
	baseURL := "https://api.bilibili.com/x/web-interface/archive/has/like"

	formData := map[string]string{
		"aid":  fmt.Sprintf("%d", aid),
		"bvid": bvid,
	}

	if v.client.AccessKey != "" {
		formData["access_key"] = v.client.AccessKey
	}

	resp, err := v.client.HTTPClient.R().
		SetFormData(formData).
		SetCookie(&http.Cookie{
			Name:  "SESSDATA",
			Value: v.client.SESSDATA,
		}).
		SetResult(&HasLikeResponse{}).
		Get(baseURL)

	if err != nil {
		return nil, err
	}

	return resp.Result().(*HasLikeResponse), nil
}

// Dislike sends a request to dislike or cancel dislike for a video.
//
// Parameters:
//   - aid (int): 视频aid
//   - dislike (int): 操作类型, 0表示点踩, 1表示取消点踩
//
// Authentication:
//   - 认证方式：仅可App，使用access_key进行认证
func (v *Video) DislikeApp(aid int, dislike int) (*DislikeResponse, error) {
	baseURL := "https://app.biliapi.net/x/v2/view/dislike"

	formData := map[string]string{
		"access_key": v.client.AccessKey,
		"aid":        fmt.Sprintf("%d", aid),
		"dislike":    fmt.Sprintf("%d", dislike),
	}

	resp, err := v.client.HTTPClient.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(formData).
		SetResult(&DislikeResponse{}).
		Post(baseURL)

	if err != nil {
		return nil, err
	}

	return resp.Result().(*DislikeResponse), nil
}

// 投币视频（web端）
//
// Parameters:
//   - aid (int): 稿件avid, 与bvid任选一个
//   - bvid (string): 稿件bvid, 与aid任选一个
//   - multiply (int): 投币数量, 上限为2
//   - select_like (int): 是否附加点赞, 0表示不点赞, 1表示同时点赞
//   - csrf (string): CSRF Token, 从Cookie中获取
//
// Authentication:
//   - 认证方式：仅可Cookie，使用SESSDATA进行认证
func (v *Video) Coin(aid int, bvid string, multiply int, selectLike int) (*CoinResponse, error) {
	baseURL := "https://api.bilibili.com/x/web-interface/coin/add"

	// 检查buvid3是否存在
	if v.client.Buvid3 == "" {
		return nil, fmt.Errorf("buvid3 is required but not provided")
	}

	formData := map[string]string{
		"aid":         fmt.Sprintf("%d", aid),
		"bvid":        bvid,
		"multiply":    fmt.Sprintf("%d", multiply),
		"select_like": fmt.Sprintf("%d", selectLike),
		"csrf":        v.client.CSRF,
	}

	resp, err := v.client.HTTPClient.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(formData).
		SetCookies([]*http.Cookie{
			{
				Name:  "SESSDATA",
				Value: v.client.SESSDATA},
			{
				Name:  "buvid3",
				Value: fmt.Sprint(v.client.Buvid3)}}).
		SetResult(&CoinResponse{}).
		Post(baseURL)

	if err != nil {
		return nil, err
	}

	return resp.Result().(*CoinResponse), nil
}

// 投币视频（APP端）
//
// 参数：
//   - aid (int): 稿件 avid
//   - multiply (int): 投币数量，上限为2
//   - selectLike (int): 附加点赞，0表示不点赞，1表示同时点赞，默认为0
//
// 备注：
//   - 认证方式：仅可APP
func (v *Video) CoinApp(aid int, multiply int, selectLike int) (*CoinAppResponse, error) {
	baseURL := "https://app.biliapi.com/x/v2/view/coin/add"

	formData := map[string]string{
		"access_key":  v.client.AccessKey,
		"aid":         fmt.Sprintf("%d", aid),
		"multiply":    fmt.Sprintf("%d", multiply),
		"select_like": fmt.Sprintf("%d", selectLike),
	}

	resp, err := v.client.HTTPClient.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(formData).
		SetResult(&CoinAppResponse{}).
		Post(baseURL)

	if err != nil {
		return nil, err
	}

	return resp.Result().(*CoinAppResponse), nil
}

// 判断视频是否被投币（双端）
//
// 参数：
//   - aid (int): 稿件 avid，avid 与 bvid 任选一个
//   - bvid (string): 稿件 bvid，avid 与 bvid 任选一个
//
// 备注：
//   - 认证方式：APP或Cookie（SESSDATA）
func (v *Video) CoinsStatus(aid int, bvid string) (*CoinsStatusResponse, error) {
	baseURL := "https://api.bilibili.com/x/web-interface/archive/coins"

	formData := map[string]string{
		"aid":  fmt.Sprintf("%d", aid),
		"bvid": bvid,
	}

	if v.client.AccessKey != "" {
		formData["access_key"] = v.client.AccessKey
	}

	resp, err := v.client.HTTPClient.R().
		SetCookie(&http.Cookie{
			Name:  "SESSDATA",
			Value: v.client.SESSDATA,
		}).
		SetQueryParams(formData).
		SetResult(&CoinsStatusResponse{}).
		Get(baseURL)

	if err != nil {
		return nil, err
	}

	return resp.Result().(*CoinsStatusResponse), nil
}

// 收藏视频（双端）
//
// Parameters:
//   - rid (int): 视频的aid
//   - addMediaIDs (string): 需要加入的收藏夹mlid，多个用逗号分隔（2选1）
//   - delMediaIDs (string): 需要取消的收藏夹mlid，多个用逗号分隔（2选1）
//
// Authentication:
//   - 认证方式：APP或Cookie（SESSDATA），Cookie方式时需要验证referer为.bilibili.com域名下
//   - 使用access_key进行APP认证
//   - 使用csrf token进行Cookie认证
func (v *Video) Collect(rid int, addMediaIDs, delMediaIDs string) (*CollectResponse, error) {
	baseURL := "https://api.bilibili.com/medialist/gateway/coll/resource/deal"

	formData := map[string]string{
		"rid":           fmt.Sprintf("%d", rid),
		"type":          "2",
		"csrf":          v.client.CSRF,
		"add_media_ids": addMediaIDs,
		"del_media_ids": delMediaIDs,
	}

	if v.client.AccessKey != "" {
		formData["access_key"] = v.client.AccessKey
	}

	resp, err := v.client.HTTPClient.R().
		SetFormData(formData).
		SetCookie(&http.Cookie{
			Name:  "SESSDATA",
			Value: v.client.SESSDATA,
		}).
		SetHeader("Referer", "https://www.bilibili.com").
		SetResult(&CollectResponse{}).Post(baseURL)

	if err != nil {
		return nil, err
	}

	return resp.Result().(*CollectResponse), nil
}

// 收藏视频（Web端）
//
// Parameters:
//   - rid (int): 视频的aid
//   - addMediaIDs (string): 需要加入的收藏夹mlid，多个用逗号分隔（可选）
//   - delMediaIDs (string): 需要取消的收藏夹mlid，多个用逗号分隔（可选）
//
// Authentication:
//   - 认证方式：Cookie（SESSDATA），需要设置csrf
func (v *Video) CollectWeb(rid int, addMediaIDs, delMediaIDs string) (*WebCollectResponse, error) {
	baseURL := "https://api.bilibili.com/x/v3/fav/resource/deal"

	formData := map[string]string{
		"rid":           fmt.Sprintf("%d", rid),
		"type":          "2",
		"add_media_ids": addMediaIDs,
		"del_media_ids": delMediaIDs,
		"csrf":          v.client.CSRF,
		// "platform":      "web",
		// "eab_x":         "1",
		// "ga":            "1",
		// "gaia_source":   "web_normal",
	}

	req := v.client.HTTPClient.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(formData).
		SetResult(&WebCollectResponse{}).
		SetCookie(&http.Cookie{
			Name:  "SESSDATA",
			Value: v.client.SESSDATA,
		}).
		SetHeader("Referer", "https://www.bilibili.com")

	resp, err := req.Post(baseURL)
	if err != nil {
		return nil, err
	}

	return resp.Result().(*WebCollectResponse), nil
}

// IsFavoured checks whether a video is favoured.
//
// Parameters:
//   - aid (int or string): 视频的aid或bvid，任选其一
//
// Authentication:
//   - 认证方式：APP（使用access_key）或Cookie（SESSDATA）
func (v *Video) IsFavoured(aid interface{}) (*FavouredResponse, error) {
	baseURL := "https://api.bilibili.com/x/v2/fav/video/favoured"

	formData := map[string]string{
		"aid": fmt.Sprintf("%v", aid),
	}
	if v.client.AccessKey != "" {
		formData["access_key"] = v.client.AccessKey
	}

	resp, err := v.client.HTTPClient.R().
		SetQueryParams(formData).
		SetCookie(&http.Cookie{
			Name:  "SESSDATA",
			Value: v.client.SESSDATA,
		}).
		SetResult(&FavouredResponse{}).Get(baseURL)

	if err != nil {
		return nil, err
	}

	return resp.Result().(*FavouredResponse), nil
}

// TripleLike performs the triple action of liking, coin, and favouring a video.
//
// Parameters:
//   - aid (int): 视频的aid（可选）
//   - bvid (string): 视频的bvid（可选）
//
// Authentication:
//   - 认证方式：Cookie（SESSDATA），需要设置csrf token
func (v *Video) TripleLike(aid int, bvid string) (*TripleLikeResponse, error) {
	baseURL := "https://api.bilibili.com/x/web-interface/archive/like/triple"

	formData := map[string]string{
		"aid":  fmt.Sprintf("%d", aid),
		"bvid": bvid,
		"csrf": v.client.CSRF,
	}

	resp, err := v.client.HTTPClient.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(formData).
		SetResult(&TripleLikeResponse{}).
		SetCookie(&http.Cookie{
			Name:  "SESSDATA",
			Value: v.client.SESSDATA,
		}).
		Post(baseURL)

	if err != nil {
		return nil, err
	}

	return resp.Result().(*TripleLikeResponse), nil
}

// AppTripleLike performs the triple action of liking, coin, and favouring a video on the app.
//
// Parameters:
//   - aid (int): 视频的aid
//
// Authentication:
//   - 认证方式：APP（使用access_key）
func (v *Video) TripleLikeApp(aid int) (*AppTripleLikeResponse, error) {
	baseURL := "https://app.biliapi.net/x/v2/view/like/triple"

	formData := map[string]string{
		"aid":        fmt.Sprintf("%d", aid),
		"access_key": v.client.AccessKey,
	}

	req := v.client.HTTPClient.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(formData).
		SetResult(&AppTripleLikeResponse{})

	resp, err := req.Post(baseURL)
	if err != nil {
		return nil, err
	}

	return resp.Result().(*AppTripleLikeResponse), nil
}

// 分享视频 （Web端）?貌似会提示账号异常
//
// Parameters:
//   - aid (int): 视频的aid
//
// Authentication:
//   - 认证方式：Cookie（需要设置csrf token）
func (v *Video) Share(aid int) (*ShareResponse, error) {
	baseURL := "https://api.bilibili.com/x/web-interface/share/add"

	formData := map[string]string{
		"aid":    fmt.Sprintf("%d", aid),
		"csrf":   v.client.CSRF,
		"eab_x":  "2",
		"ramval": "1",
		"source": "web_normal",
		"ga":     "1",
	}

	resp, err := v.client.HTTPClient.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(formData).
		SetResult(&ShareResponse{}).
		Post(baseURL)
	if err != nil {
		return nil, err
	}

	return resp.Result().(*ShareResponse), nil
}

//--

// LikeResponse represents the structure of the API response for liking a video.
type LikeResponse struct {
	Code    int    `json:"code"`    // 返回值: 0表示成功, -101表示账号未登录, -111表示csrf校验失败, -400表示请求错误, -403表示账号异常, 10003表示不存在该稿件, 65004表示取消点赞失败, 65006表示重复点赞
	Message string `json:"message"` // 错误信息, 默认为"0"
	TTL     int    `json:"ttl"`     // 恒为1
}

//-

type AppLikeResponse struct {
	Code    int      `json:"code"`    // 返回值: 0表示成功, -101表示账号未登录, -400表示请求错误, -403表示账号异常, 10003表示不存在该稿件
	Message string   `json:"message"` // 错误信息, 默认为0
	TTL     int      `json:"ttl"`     // 固定值1
	Data    LikeData `json:"data"`    // 数据本体
}

// LikeData represents the data structure in the LikeResponse.
type LikeData struct {
	Toast string `json:"toast"` // 提示信息内容
}

// -

type HasLikeResponse struct {
	Code    int    `json:"code"`    // 返回值: 0表示成功, -400表示请求错误, -101表示账号未登录
	Message string `json:"message"` // 错误信息, 默认为0
	TTL     int    `json:"ttl"`     // 恒为1
	Data    int    `json:"data"`    // 被点赞标志: 0表示未点赞, 1表示已点赞
}

// -
type DislikeResponse struct {
	Code    int    `json:"code"`    // 返回值: 0表示成功, -101表示账号未登录, -400表示请求错误, -404表示啥都木有, 65005表示取消踩失败未点踩过, 65007表示已踩过
	Message string `json:"message"` // 错误信息, 默认为"0"
	TTL     int    `json:"ttl"`     // 固定值: 1
}

// -
type CoinResponse struct {
	Code    int    `json:"code"`    // 返回值: 0表示成功, -101表示账号未登录, -102表示账号被封停, -104表示硬币不足, -111表示csrf校验失败, -400表示请求错误, -403表示账号异常, 10003表示不存在该稿件, 34002表示不能给自己投币, 34003表示非法的投币数量, 34004表示投币间隔太短, 34005表示超过投币上限
	Message string `json:"message"` // 错误信息, 默认为0
	TTL     int    `json:"ttl"`     // TTL, 固定值1
	Data    struct {
		Like bool `json:"like"` // 是否点赞成功, true表示成功, false表示失败
	} `json:"data"`
}

// -
type CoinAppResponse struct {
	Code    int    `json:"code"`    // 返回值: 0表示成功, -101表示账号未登录, -102表示账号被封停, -104表示硬币不足, -400表示请求错误, 10003表示不存在该稿件, 34002表示不能给自己投币, 34003表示非法的投币数量, 34004表示投币间隔太短, 34005表示超过投币上限
	Message string `json:"message"` // 错误信息, 默认为0
	TTL     int    `json:"ttl"`     // TTL, 固定值1
	Data    struct {
		Like bool `json:"like"` // 是否点赞成功, true表示成功, false表示失败
	} `json:"data"`
}

// -
// CoinsStatusResponse represents the structure of the API response for checking video coin status.
type CoinsStatusResponse struct {
	Code    int    `json:"code"`    // 返回值: 0表示成功, -400表示请求错误, -101表示账号未登录
	Message string `json:"message"` // 错误信息, 默认为0
	TTL     int    `json:"ttl"`     // TTL, 固定值1
	Data    struct {
		Multiply int `json:"multiply"` // 投币枚数, 未投币为0
	} `json:"data"`
}

// -

type CollectResponse struct {
	Code    int    `json:"code"`    // 返回值: 0表示成功, -101表示账号未登录, -111表示csrf校验失败, -400表示请求错误, -403表示访问权限不足, 10003表示不存在该稿件, 11010表示内容不存在, 11201表示已经收藏过了, 11202表示已经取消收藏了, 11203表示达到收藏上限, 72010017表示参数错误
	Message string `json:"message"` // 错误信息, 默认为"success"
	Data    struct {
		Prompt bool `json:"prompt"` // 是否为未关注用户收藏, false表示否, true表示是
	} `json:"data"`
}

// -

type WebCollectResponse struct {
	Code    int    `json:"code"`    // 返回值: 0表示成功, -101表示账号未登录, -111表示csrf校验失败, 2001000表示参数错误
	Message string `json:"message"` // 错误信息, 默认为0
	TTL     int    `json:"ttl"`     // TTL, 固定值1
	Data    struct {
		Prompt     bool      `json:"prompt"`      // 是否为未关注用户收藏, false表示否, true表示是
		GaData     *struct{} `json:"ga_data"`     // 作用尚不明确，可能为null
		ToastMsg   string    `json:"toast_msg"`   // 空，作用尚不明确
		SuccessNum int       `json:"success_num"` // 作用尚不明确
	} `json:"data"`
}

// -
type FavouredResponse struct {
	Code    int    `json:"code"`    // 返回值: 0表示成功, -400表示请求错误, -101表示账号未登录
	Message string `json:"message"` // 错误信息, 默认为0
	TTL     int    `json:"ttl"`     // TTL, 固定值1
	Data    struct {
		Count    int  `json:"count"`    // 作用尚不明确
		Favoured bool `json:"favoured"` // 是否收藏, true表示已收藏, false表示未收藏
	} `json:"data"`
}

// TripleLikeResponse represents the structure of the API response for the triple like action.
type TripleLikeResponse struct {
	Code    int    `json:"code"`    // 返回值: 0表示成功, -101表示账号未登录, -111表示csrf校验失败, -400表示请求错误, 10003表示不存在该稿件, -403表示账号异常
	Message string `json:"message"` // 错误信息, 默认为0
	TTL     int    `json:"ttl"`     // TTL, 固定值1
	Data    struct {
		Like     bool `json:"like"`     // 是否点赞成功, true表示成功, false表示失败
		Coin     bool `json:"coin"`     // 是否投币成功, true表示成功, false表示失败
		Fav      bool `json:"fav"`      // 是否收藏成功, true表示成功, false表示失败
		Multiply int  `json:"multiply"` // 投币枚数, 默认为2
	} `json:"data"`
}

// /-
type AppTripleLikeResponse struct {
	Code    int    `json:"code"`    // 返回值: 0表示成功, -101表示账号未登录, -400表示请求错误, 10003表示不存在该稿件
	Message string `json:"message"` // 错误信息, 默认为0
	TTL     int    `json:"ttl"`     // TTL, 固定值1
	Data    struct {
		Like     bool `json:"like"`     // 是否点赞成功, true表示成功, false表示失败
		Coin     bool `json:"coin"`     // 是否投币成功, true表示成功, false表示失败
		Fav      bool `json:"fav"`      // 是否收藏成功, true表示成功, false表示失败
		Multiply int  `json:"multiply"` // 投币枚数, 默认为2
	} `json:"data"`
}

// ShareResponse represents the structure of the API response for sharing a video.
type ShareResponse struct {
	Code    int    `json:"code"`    // 返回值: 0表示成功, -101表示账号未登录, -111表示csrf校验失败, -400表示请求错误
	Message string `json:"message"` // 错误信息, 默认为0
	TTL     int    `json:"ttl"`     // TTL, 固定值1
	Data    int    `json:"data"`    // 当前分享数
}
