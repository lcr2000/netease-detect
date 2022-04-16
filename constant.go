package netease_detect

const (
	LiveDetectSubmitUrl        = "http://as.dun.163.com/v3/livewallsolution/check"            // 直播音视频检测提交
	LiveDetectResultUrl        = "http://as.dun.163.com/v3/livewallsolution/callback/results" // 获取直播音视频检测结果
	LiveDetectStopUrl          = "http://as.dun.163.com/v1/livewallsolution/feedback"         // 直播音视频停止检测
	LiveAudioDetectSubmitUrl   = "http://as.dun.163.com/v4/liveaudio/check"                   // 直播音频检测提交
	LiveAudioDetectStopUrl     = "http://as.dun.163.com/v1/liveaudio/feedback"                // 直播音频停止检测
	LiveAudioDetectResultUrl   = "http://as.dun.163.com/v4/liveaudio/callback/results"        // 获取直播音频检测结果
	LiveAudioDetectFeedbackUrl = "http://as.dun.163.com/v1/audio/feedback"                    // 音频检测反馈接口
	DemandAudioDetectResultUrl = "https://as.dun.163.com/v3/audio/callback/results"           // 点播音频检测结果
	DemandAudioDetectSubmitUrl = "http://as.dun.163.com/v3/audio/submit"                      // 点播音频信息提交接口
	DemandVideoDetectResultUrl = "http://as.dun.163.com/v3/video/callback/results"            // 点播视频检测结果
	DemandVideoDetectSubmitUrl = "http://as.dun.163.com/v3/video/submit"                      // 点播视频信息提交接口
	ImageDetectUrl             = "http://as.dun.163.com/v4/image/check"                       // 图片在线检测
)

// 各个接口版本号
const (
	ApiVersionV1 = "v1"
	ApiVersionV3 = "v3"
	ApiVersionV4 = "v4"
)

const (
	StopDetectStatus = 100 // 停止检测
)

const (
	CallSuccessCode = 200 // 接口调用成功状态码
)
