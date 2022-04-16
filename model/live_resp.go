package model

type LiveDetectSubmitResp struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Result struct {
		TaskID string `json:"taskId"`
		Status bool   `json:"status"`
	} `json:"result"`
}

type LiveDetectStopResp struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Result []struct {
		TaskID string `json:"taskId"`
		Result int    `json:"result"`
	} `json:"result"`
}

type LiveDetectResultResp struct {
	Code   int                 `json:"code"`
	Msg    string              `json:"msg"`
	Result []*LiveDetectResult `json:"result"`
}

type LiveDetectResult struct {
	Antispam *LiveDetectAntispam `json:"antispam" bson:"antispam"` // 否 内容安全检测结果
}

type LiveDetectAntispam struct {
	TaskID        string               `json:"taskId" bson:"taskId"`               // 是 检测任务 ID，示例值："38e08da8d2574df4bd2eca9b5153df72"
	Callback      string               `json:"callback" bson:"callback"`           // 否 提交时传递的callback
	DataID        string               `json:"dataId" bson:"dataId"`               // 否 提交时传递的dataId
	Status        int                  `json:"status" bson:"status"`               // 是 检测状态，1-检测中，2-检测成功，3-检测失败
	FailureReason int                  `json:"failureReason" bson:"failureReason"` // 否 检测失败原因，当检测失败时返回，1：下载失败，2：直播流不存在，3：解析失败，4：格式错误
	RiskLevel     int                  `json:"riskLevel" bson:"riskLevel"`         // 否 直播异常风险等级， 0-正常，1-低危，2-中危，3-中高危，4-高危
	Duration      int                  `json:"duration" bson:"duration"`           // 否 直播时长字段，直播结束返回直播整体时长，单位s
	Evidences     *LiveDetectEvidences `json:"evidences" bson:"evidences"`         // 否 直播审核证据信息
}

type LiveDetectEvidences struct {
	LiveDetectVideo *LiveDetectVideo `json:"video" bson:"video"`
	LiveDetectAudio *LiveDetectAudio `json:"audio" bson:"audio"`
}

type LiveDetectAudio struct {
	Suggestion   int                     `json:"suggestion" bson:"suggestion"`     // 是 建议结果 0-通过 1-嫌疑 2-删除
	StartTime    int64                   `json:"startTime" bson:"startTime"`       // 是	断句开始时间，单位毫秒
	EndTime      int64                   `json:"endTime" bson:"endTime"`           // 是 断句结束时间，单位毫秒
	Content      string                  `json:"content" bson:"content"`           // 是	检测命中内容返回
	Url          string                  `json:"url" bson:"url"`                   // 否 需要开启返回直播语音片段播放地址，请联系易盾策略经理
	SpeakerId    string                  `json:"speakerId" bson:"speakerId"`       // 否	针对接入SDK监听客户，返回说话人ID信息
	FrontSegment *FrontAudioSegment      `json:"frontSegment" bson:"frontSegment"` // 否 关联证据信息，异常/嫌疑断句命中时返回，命中断句前20s的证据信息，包含前20s的音频语音识别内容及对应url，通过时数据为空
	Labels       []*LiveDetectAudioLabel `json:"labels" bson:"labels"`             // 否 分类信息，通过时数据为空
	SegmentId    string                  `json:"segmentId" bson:"segmentId"`       // 是	断句id
}

type LiveDetectVideo struct {
	Evidence *LiveDetectEvidenceDetail `json:"evidence" bson:"evidence"` // 否 证据信息
	Labels   []*LiveDetectLabel        `json:"labels" bson:"labels"`     // 否 分类信息
}

type LiveDetectEvidenceDetail struct {
	Suggestion int                   `json:"suggestion" bson:"suggestion"` // 是 建议结果 0-通过 1-嫌疑 2-删除
	Type       int                   `json:"type" bson:"type"`             // 是 证据信息类型，1-图片，2-视频
	URL        string                `json:"url" bson:"url"`               // 是 证据信息内容
	BeginTime  int64                 `json:"beginTime" bson:"beginTime"`   // 是 直播当前时间点，单位为毫秒
	EndTime    int64                 `json:"endTime" bson:"endTime"`       // 是 直播当前时间点，单位为毫秒
	SpeakerId  string                `json:"speakerId" bson:"speakerId"`   // 否 针对接入SDK监听客户，返回说话人ID信息
	FrontPics  []*LiveDetectFrontPic `json:"frontPics" bson:"frontPics"`   // 是 关联信息-命中前截图信息
}

type LiveDetectFrontPic struct {
	URL string `json:"url"` // 是 命中前截图信息
}

type LiveDetectAudioLabel struct {
	Label     int                        `json:"label" bson:"label"`         // 是 分类信息，100：色情，110：性感低俗，200：广告，210：二维码，260：广告法，300：暴恐，400：违禁，500：涉政，800：恶心类，900：其他，1020：黑屏，1030：挂机，1100：涉价值观
	Level     int                        `json:"level" bson:"level"`         // 是 分类级别，0：正常，1：不确定，2：确定
	Rate      float64                    `json:"rate" bson:"rate"`           // 是 置信度分数，0-1之间取值，1为置信度最高，0为置信度最低。若level为正常，置信度越大，说明正常的可能性越高。若level为不确定或确定，置信度越大，说明垃圾的可能性越高
	SubLabels []*LiveDetectAudioSubLabel `json:"subLabels" bson:"subLabels"` // 是 细分类信息，可能包含多个
}

type LiveDetectLabel struct {
	Label     int                   `json:"label" bson:"label"`         // 是 分类信息，100：色情，110：性感低俗，200：广告，210：二维码，260：广告法，300：暴恐，400：违禁，500：涉政，800：恶心类，900：其他，1020：黑屏，1030：挂机，1100：涉价值观
	Level     int                   `json:"level" bson:"level"`         // 是 分类级别，0：正常，1：不确定，2：确定
	Rate      float64               `json:"rate" bson:"rate"`           // 是 置信度分数，0-1之间取值，1为置信度最高，0为置信度最低。若level为正常，置信度越大，说明正常的可能性越高。若level为不确定或确定，置信度越大，说明垃圾的可能性越高
	SubLabels []*LiveDetectSubLabel `json:"subLabels" bson:"subLabels"` // 是 细分类信息，可能包含多个
}

type LiveDetectAudioSubLabel struct {
	SubLabel string                     `json:"subLabel" bson:"subLabel"` // 是 细分类，详细编码请参考下方对应细分类编码对照表
	Rate     float64                    `json:"rate" bson:"rate"`         // 是 置信度分数，0-1之间取值，1为置信度最高，0为置信度最低
	Details  *LiveDetectSubLabelDetails `json:"details" bson:"details"`   // 否 命中的详细对象信息
}

// LiveDetectSubLabel 网易坑爹的，audio和video subLabel数据类型不一致
type LiveDetectSubLabel struct {
	SubLabel int                        `json:"subLabel" bson:"subLabel"` // 是 细分类，详细编码请参考下方对应细分类编码对照表
	Rate     float64                    `json:"rate" bson:"rate"`         // 是 置信度分数，0-1之间取值，1为置信度最高，0为置信度最低
	Details  *LiveDetectSubLabelDetails `json:"details" bson:"details"`   // 否 命中的详细对象信息
}

type LiveDetectSubLabelDetails struct {
	HitInfos []*LiveDetectHitInfo `json:"hitInfos" bson:"hitInfos"` // 是 命中的线索信息
}

type LiveDetectHitInfo struct {
	Value string `json:"value" bson:"value"` // 是 图片中包含的可识别内容
	Group string `json:"group" bson:"group"` // 否 value对应的分组名称，用于对value的解释
}

type FrontAudioSegment struct {
	Url     string `json:"url" bson:"url"`         // 是 音频断句url
	Content string `json:"content" bson:"content"` // 是 音频断句语音识别结果
}
