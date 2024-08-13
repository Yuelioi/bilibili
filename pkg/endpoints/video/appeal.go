package video

import (
	"net/http"
)

// 获取投诉类型
//
// Authentication:
//   - 认证方式：Cookie（SESSDATA）
func (a *Video) AppealTags() (*AppealTagsResponse, error) {
	baseURL := "https://api.bilibili.com/x/web-interface/archive/appeal/tags"

	resp, err := a.client.HTTPClient.R().
		SetResult(&AppealTagsResponse{}).
		SetCookie(&http.Cookie{
			Name:  "SESSDATA",
			Value: a.client.SESSDATA,
		}).
		Get(baseURL)

	if err != nil {
		return nil, err
	}

	return resp.Result().(*AppealTagsResponse), nil
}

// 投诉稿件
//
// Parameters:
//   - aid (int): 视频的aid
//   - tid (int): 投诉理由tid
//   - desc (string): 投诉理由详细描述
//   - attach (string): 附件（多个附件用逗号隔开，非必要）
//   - buid (string): 风控代码（请求头）
//   - csrf (string): CSRF token（在cookie中）
//
// Authentication:
//   - 认证方式：Cookie（SESSDATA）
func (a *Video) SubmitAppeal(aid int, tid int, desc string, attach string, buid string, csrf string) (*SubmitAppealResponse, error) {
	baseURL := "https://api.bilibili.com/x/web-interface/appeal/v2/submit"

	req := a.client.HTTPClient.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetHeader("Buid", buid).
		SetHeader("Referer", "https://www.bilibili.com/").
		SetBody(map[string]interface{}{
			"aid":    aid,
			"tid":    tid,
			"desc":   desc,
			"attach": attach,
			"csrf":   csrf,
		}).
		SetResult(&SubmitAppealResponse{}).
		SetCookie(&http.Cookie{
			Name:  "buid",
			Value: buid,
		}).
		SetCookie(&http.Cookie{
			Name:  "bili_jct",
			Value: csrf,
		})

	resp, err := req.Post(baseURL)
	if err != nil {
		return nil, err
	}

	return resp.Result().(*SubmitAppealResponse), nil
}

// AppealTagsResponse represents the structure of the API response for retrieving appeal types.
type AppealTagsResponse struct {
	Code    int    `json:"code"`    // 返回值: 0表示成功
	Message string `json:"message"` // 错误信息, 默认为0
	TTL     int    `json:"ttl"`     // TTL, 固定值1
	Data    []struct {
		TID      int    `json:"tid"`      // 类型tid
		Business int    `json:"business"` // 意义不明
		Weight   int    `json:"weight"`   // 权重
		Round    int    `json:"round"`    // 意义不明
		State    int    `json:"state"`    // 意义不明
		Name     string `json:"name"`     // 类型名称
		Remark   string `json:"remark"`   // 类型备注
		Ctime    string `json:"ctime"`    // 意义不明
		Mtime    string `json:"mtime"`    // 意义不明
		Controls []struct {
			TID         int    `json:"tid"`         // 同上
			BID         int    `json:"bid"`         // 意义不明
			Name        string `json:"name"`        // 提示名称
			Title       string `json:"title"`       // 提示标题
			Component   string `json:"component"`   // 需要填入的类型
			Placeholder string `json:"placeholder"` // 文本框占位符
			Required    int    `json:"required"`    // 是否为必填
		} `json:"controls,omitempty"` // 控件提示
	} `json:"data"` // 类型条目
}

// SubmitAppealResponse represents the structure of the API response for submitting an appeal.
type SubmitAppealResponse struct {
	Code    int    `json:"code"`    // 返回码: 0表示成功
	Message string `json:"message"` // 错误信息, 默认为0
	TTL     int    `json:"ttl"`     // TTL, 固定值1
}
