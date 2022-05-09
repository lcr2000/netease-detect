package netease

const (
	LiveSubmitURL        = "http://as.dun.163.com/v3/livewallsolution/check"            // 直播音视频检测提交
	LiveResultURL        = "http://as.dun.163.com/v3/livewallsolution/callback/results" // 获取直播音视频检测结果
	LiveStopURL          = "http://as.dun.163.com/v1/livewallsolution/feedback"         // 直播音视频停止检测
	LiveAudioSubmitURL   = "http://as.dun.163.com/v4/liveaudio/check"                   // 直播音频检测提交
	LiveAudioStopURL     = "http://as.dun.163.com/v1/liveaudio/feedback"                // 直播音频停止检测
	LiveAudioResultURL   = "http://as.dun.163.com/v4/liveaudio/callback/results"        // 获取直播音频检测结果
	LiveAudioFeedbackURL = "http://as.dun.163.com/v1/audio/feedback"                    // 音频检测反馈接口
	AudioResultURL       = "https://as.dun.163.com/v3/audio/callback/results"           // 点播音频异步检测结果
	AudioSubmitURL       = "http://as.dun.163.com/v3/audio/submit"                      // 点播音频异步检测接口
	AudioURL             = "http://as.dun.163.com/v2/audio/check"                       // 点播音频同步检测
	VideoResultURL       = "http://as.dun.163.com/v3/video/callback/results"            // 点播视频检测结果
	VideoSubmitURL       = "http://as.dun.163.com/v3/video/submit"                      // 点播视频信息提交接口
	ImageURL             = "http://as.dun.163.com/v4/image/check"                       // 图片在线检测
)

// 各个接口版本号
const (
	APIVersionV1 = "v1"
	APIVersionV2 = "v2"
	APIVersionV3 = "v3"
	APIVersionV4 = "v4"
)

const (
	// StopDetectStatus 停止检测
	StopDetectStatus = 100
)

const (
	// CallSuccessCode 接口调用成功状态码
	CallSuccessCode = 200
)

const (
	DataCheckTypeURL    = 0
	DataCheckTypeBase64 = 1
)
