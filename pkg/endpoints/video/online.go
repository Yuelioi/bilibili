package video

import "fmt"

// 获取视频在线人数 (Web端)
//
// Parameters:
//   - aid (int): 视频的aid (可选)
//   - bvid (string): 视频的bvid (可选)
//   - cid (int): 视频的cid (必要)
func (v *Video) OnlineTotal(aid int, bvid string, cid int) (*OnlineTotalResponse, error) {

	baseURL := "https://api.bilibili.com/x/player/online/total"

	formData := map[string]string{
		"aid":  fmt.Sprintf("%d", aid),
		"bvid": bvid,
		"cid":  fmt.Sprintf("%d", cid),
	}

	resp, err := v.client.HTTPClient.R().
		SetQueryParams(formData).
		SetResult(&OnlineTotalResponse{}).
		Get(baseURL)
	if err != nil {
		return nil, err
	}

	return resp.Result().(*OnlineTotalResponse), nil
}

// 获取视频在线人数_APP端
// Parameters:
//   - aid (int): 视频的aid
//   - appkey (string): APP密钥
//   - cid (int): 视频的cid
//   - ts (int64): 当前时间戳
//   - sign (string): APP签名
func (v *Video) AppOnlineTotal(aid int, appkey string, cid int, ts int64, sign string) (*AppOnlineTotalResponse, error) {
	baseURL := "https://app.bilibili.com/x/v2/view/video/online"

	formData := map[string]string{
		"aid":    fmt.Sprintf("%d", aid),
		"appkey": appkey,
		"cid":    fmt.Sprintf("%d", cid),
		"ts":     fmt.Sprintf("%d", ts),
		"sign":   sign,
	}

	resp, err := v.client.HTTPClient.R().
		SetQueryParams(formData).
		SetResult(&AppOnlineTotalResponse{}).
		Get(baseURL)
	if err != nil {
		return nil, err
	}

	return resp.Result().(*AppOnlineTotalResponse), nil
}

type OnlineTotalResponse struct {
	Code    int    `json:"code"`    // 返回值: 0：成功 -400：请求错误 -404：无视频
	Message string `json:"message"` // 错误信息，默认为0
	Ttl     int    `json:"ttl"`     // 1
	Data    struct {
		Total      string `json:"total"` // 所有终端总计人数, 例如10万+
		Count      string `json:"count"` // web端实时在线人数
		ShowSwitch struct {
			Total bool `json:"total"` // 展示所有终端总计人数
			Count bool `json:"count"` // 展示web端实时在线人数
		} `json:"show_switch"` // 数据显示控制
	} `json:"data"` // 信息本体
}

type AppOnlineTotalResponse struct {
	Code    int    `json:"code"`    // 返回值: 0：成功 -400：请求错误 -404：无视频
	Message string `json:"message"` // 错误信息，默认为0
	Ttl     int    `json:"ttl"`     // 1
	Data    struct {
		Online string `json:"online"` // 所有终端总计人数, 例如10万+人在看
	} `json:"data"` // 信息本体
}
