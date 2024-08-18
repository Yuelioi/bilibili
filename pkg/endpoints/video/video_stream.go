package video

import (
	"fmt"
	"net/http"
)

var VideoQualityMap = map[int]string{
	6:   "240P 极速",     //  - 仅 MP4 格式支持，仅 platform=html5 时有效"
	16:  "360P 流畅",     //
	32:  "480P 清晰",     //
	64:  "720P 高清",     // - WEB 端默认值，B站前端需要登录才能选择，但直接发送请求可以不登录就拿到 720P 的取流地址，无 720P 时则为 720P60"
	74:  "720P60 高帧率",  // - 登录认证",
	80:  "1080P 高清",    // - TV 端与 APP 端默认值，登录认证"
	112: "1080P+ 高码率",  // - 大会员认证"
	116: "1080P60 高帧率", // - 大会员认证"
	120: "4K 超清",       // - 需要 fnval&128=128 且 fourk=1，大会员认证"
	125: "HDR 真彩色",     // - 仅支持 DASH 格式，需 fnval&64=64，大会员认证"
	126: "杜比视界",        // - 仅支持 DASH 格式，需 fnval&512=512，大会员认证"
	127: "8K 超高清",      // - 仅支持 DASH 格式，需 fnval&1024=1024，大会员认证"
}

// 视频编码代码
var VideoCodecMap = map[int]string{
	7:  "AVC",
	12: "HEVC",
	13: "AV1",
}

// 视频伴音音质代码
var AudioQualityMap = map[int]string{
	30216: "64K",
	30232: "132K",
	30280: "192K",
	30250: "杜比全景声",
	30251: "Hi-Res无损",
}

// Parameters:
//   - avid (int): 稿件 avid（avid 与 bvid 任选一个）
//   - bvid (string): 稿件 bvid（avid 与 bvid 任选一个）
//   - cid (int): 视频 cid
//   - qn (int): 视频清晰度选择（非必要，默认值为32 - 480P）
//   - fnval (int): 视频流格式标识（非必要，默认值为1 - MP4格式）
//   - fnver (int): 视频流版本标识（非必要，默认值为0）
//   - fourk (int): 是否允许 4K 视频（非必要，0为最高画质1080P，1为最高画质4K）
//   - session (string): 播放会话标识（非必要，从视频播放页的 HTML 中获取）
//   - otype (string): 输出格式（非必要，固定为json）
//   - platform (string): 播放平台（非必要，pc表示web播放，html5表示移动端HTML5播放）
//   - high_quality (int): 是否高画质（非必要，platform=html5时，1表示画质为1080p）
func (v *Video) Stream(avid int, bvid string, cid int, qn int) (*StreamResponse, error) {

	baseURL := "https://api.bilibili.com/x/player/playurl"

	formData := map[string]string{
		"avid":  fmt.Sprintf("%d", avid),
		"bvid":  bvid,
		"cid":   fmt.Sprintf("%d", cid),
		"qn":    fmt.Sprintf("%d", qn),
		"fnval": "4048",
		"fnver": "0",
		"fourk": "1",
		// "session":      session,
		"otype": "json",
		// "platform": platform,
		// "high_quality": fmt.Sprintf("%d", high_quality),
	}

	resp, err := v.client.HTTPClient.R().
		SetQueryParams(formData).
		SetResult(&StreamResponse{}).
		SetCookie(&http.Cookie{
			Name:  "SESSDATA",
			Value: v.client.SESSDATA,
		}).
		SetHeader("Referer", "https://www.bilibili.com").
		Get(baseURL)
	if err != nil {
		return nil, err
	}

	return resp.Result().(*StreamResponse), nil
}

// VideoDetailResponse represents the JSON response structure for the video stream API.
type StreamResponse struct {
	Code    int         `json:"code"`    // 返回值，0：成功，-400：请求错误，-404：无视频
	Message string      `json:"message"` // 错误信息，默认为0
	Ttl     int         `json:"ttl"`     // 默认值为1
	Data    *StreamData `json:"data"`    // 数据本体，如果无效则为null
}

// DataObj represents the data object within the VideoDetailResponse.
type StreamData struct {
	From              string          `json:"from"`               // 来源（示例：local）
	Result            string          `json:"result"`             // 结果（示例：suee）
	Message           string          `json:"message"`            // 消息（示例：空）
	Quality           int             `json:"quality"`            // 清晰度标识，含义见上表
	Format            string          `json:"format"`             // 视频格式（示例：mp4/flv）
	Timelength        int             `json:"timelength"`         // 视频长度（单位为毫秒）
	AcceptFormat      string          `json:"accept_format"`      // 支持的全部格式，每项用,分隔
	AcceptDescription []string        `json:"accept_description"` // 支持的清晰度列表（文字说明）
	AcceptQuality     []int           `json:"accept_quality"`     // 支持的清晰度列表（代码）
	VideoCodecid      int             `json:"video_codecid"`      // 默认选择视频流的编码id
	SeekParam         string          `json:"seek_param"`         // Seek 参数（示例：start）
	SeekType          string          `json:"seek_type"`          // Seek 类型（示例：offset（DASH / FLV），second（MP4））
	Dash              *Dash           `json:"dash"`               // DASH 流信息，仅 DASH 格式存在此字段
	SupportFormats    []SupportFormat `json:"support_formats"`    // 支持格式的详细信息
	HighFormat        interface{}     `json:"high_format"`        // (示例：null)
	LastPlayTime      int             `json:"last_play_time"`     // 上次播放进度，单位为毫秒
	LastPlayCid       int             `json:"last_play_cid"`      // 上次播放分P的 cid
}

// SupportFormat represents the details of a supported format.
type SupportFormat struct {
	Quality        int      `json:"quality"`         // 视频清晰度代码，含义见上表
	Format         string   `json:"format"`          // 视频格式
	NewDescription string   `json:"new_description"` // 格式描述
	DisplayDesc    string   `json:"display_desc"`    // 格式描述
	Superscript    string   `json:"superscript"`     // 未知字段
	Codecs         []string `json:"codecs"`          // 可用编码格式列表
}

type Dash struct {
	Duration       int      `json:"duration"`        // 视频长度，单位为秒
	MinBufferTime  float64  `json:"minBufferTime"`   // 缓冲时间，单位为秒
	MinBuffer_time float64  `json:"min_buffer_time"` // 同上
	Video          []Stream `json:"video"`           // 视频流信息数组
	Audio          []Stream `json:"audio"`           // 伴音流信息数组，视频没有音轨时为nil
	Dolby          *Dolby   `json:"dolby"`           // 杜比全景声伴音信息
	Flac           *Flac    `json:"flac"`            // 无损音轨伴音信息，视频没有无损音轨时为nil
}
type Stream struct {
	ID             int      `json:"id"`             // 音视频清晰度代码
	BaseURL        string   `json:"baseUrl"`        // 默认流 URL
	Base_url       string   `json:"base_url"`       // 同上
	BackupURL      []string `json:"backupUrl"`      // 备用流 URL 数组
	Backup_url     []string `json:"backup_url"`     // 同上
	Bandwidth      int      `json:"bandwidth"`      // 所需最低带宽，单位为 Byte
	MimeType       string   `json:"mimeType"`       // 格式 mimetype 类型
	Mime_type      string   `json:"mime_type"`      // 同上
	Codecs         string   `json:"codecs"`         // 编码/音频类型
	Width          int      `json:"width"`          // 视频宽度，仅视频流存在该字段
	Height         int      `json:"height"`         // 视频高度，仅视频流存在该字段
	FrameRate      string   `json:"frameRate"`      // 视频帧率，仅视频流存在该字段
	Frame_rate     string   `json:"frame_rate"`     // 同上
	Sar            string   `json:"sar"`            // Sample Aspect Ratio（单个像素的宽高比），音频流该值恒为空
	StartWithSap   int      `json:"startWithSap"`   // Stream Access Point（流媒体访问位点），音频流该值恒为空
	Start_with_sap int      `json:"start_with_sap"` // 同上
	SegmentBase    *Segment `json:"SegmentBase"`    // URL 对应 m4s 文件中，头部的位置，音频流该值恒为空
	Segment_base   *Segment `json:"segment_base"`   // 同上
	Codecid        int      `json:"codecid"`        // 码流编码标识代码，音频流该值恒为0
}

type Segment struct {
	Initialization string `json:"initialization"` // 如：0-821，表示开头 820 个字节
	IndexRange     string `json:"index_range"`    // 如：822-1309，记录关键帧的时间戳及其在文件中的位置
}

type Dolby struct {
	Type  int      `json:"type"`  // 杜比音效类型 1：普通杜比音效，2：全景杜比音效
	Audio []Stream `json:"audio"` // 杜比伴音流列表
}

type Flac struct {
	Display bool   `json:"display"` // 是否在播放器显示切换 Hi-Res 无损音轨按钮
	Audio   Stream `json:"audio"`   // 音频流信息
}
