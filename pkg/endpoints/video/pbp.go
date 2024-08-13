package video

import "fmt"

// 获取弹幕趋势顶点列表
// Parameters:
//   - cid (int): 视频的cid
//   - aid (int): 视频的aid (可选)
//   - bvid (string): 视频的bvid (可选)
func (v *Video) GetHighEnergyProgress(cid int, aid int, bvid string) (*HighEnergyProgressResponse, error) {
	baseURL := "https://bvc.bilivideo.com/pbp/data"

	// Set query parameters based on provided values
	queryParams := map[string]string{
		"aid":  fmt.Sprintf("%d", aid),
		"bvid": bvid,
		"cid":  fmt.Sprintf("%d", cid),
	}

	resp, err := v.client.HTTPClient.R().
		SetQueryParams(queryParams).
		SetResult(&HighEnergyProgressResponse{}).
		Get(baseURL)
	if err != nil {
		return nil, err
	}

	return resp.Result().(*HighEnergyProgressResponse), nil
}

type HighEnergyProgressResponse struct {
	StepSec int    `json:"step_sec"` // 采样间隔时间, 单位为秒, 由视频时长决定
	TagStr  string `json:"tagstr"`   // ？？？ 作用尚不明确
	Events  struct {
		Default []int `json:"default"` // 顶点值列表
	} `json:"events"` // 数据本体
	Debug string `json:"debug"` // 调试信息, json字符串
}
