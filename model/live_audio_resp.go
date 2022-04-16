package model

type LiveAudioDetectSubmitResp struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Result struct {
		TaskID string `json:"taskId"`
		Status int    `json:"status"`
	} `json:"result"`
}

type LiveAudioDetectStopResp struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Result []struct {
		TaskID string `json:"taskId"`
		Result int    `json:"result"`
	} `json:"result"`
}

type LiveAudioDetectResultResp struct {
	Code   int                      `json:"code"`
	Msg    string                   `json:"msg"`
	Result []*LiveAudioDetectResult `json:"result"`
}

// LiveAudioDetectResult 检测结果
type LiveAudioDetectResult struct {
	Antispam *LiveAudioDetectAntispam `json:"antispam" bson:"antispam"` // 否 内容安全检测结果
	Asr      *Asr                     `json:"asr" bson:"asr"`           // 否 语音识别检测结果
	Language *Language                `json:"language" bson:"language"` // 否 语种识别检测结果
}

type LiveAudioDetectAntispam struct {
	TaskID          string                          `json:"taskId" bson:"taskId"`                   // 是 检测任务 ID，示例值："38e08da8d2574df4bd2eca9b5153df72"
	Callback        string                          `json:"callback" bson:"callback"`               // 否 提交时传递的callback
	DataID          string                          `json:"dataId" bson:"dataId"`                   // 否 提交时传递的dataId
	CensorSource    int                             `json:"censorSource" bson:"censorSource"`       // 否 审核来源，0：易盾人审，1：客户人审，2：易盾机审
	Status          int                             `json:"status" bson:"status"`                   // 是 检测状态，1-检测中，2-检测成功，3-检测失败
	FailureReason   int                             `json:"failureReason" bson:"failureReason"`     // 否 检测失败原因，当检测失败时返回，1：下载失败，2：直播流不存在，3：解析失败，4：格式错误
	RiskLevel       int                             `json:"riskLevel" bson:"riskLevel"`             // 否 直播异常风险等级， 0-正常，1-低危，2-中危，3-中高危，4-高危
	RiskScore       int                             `json:"riskScore" bson:"riskScore"`             // 否 直播异常分数
	Duration        int                             `json:"duration" bson:"duration"`               // 否 直播时长字段，直播结束返回直播整体时长，单位s
	Evidences       *LiveAudioDetectEvidence        `json:"evidences" bson:"evidences"`             // 否 直播审核证据信息
	ReviewEvidences *LiveAudioDetectReviewEvidences `json:"reviewEvidences" bson:"reviewEvidences"` // 否 直播墙人审证据信息
}

type LiveAudioDetectEvidence struct {
	Suggestion   int                `json:"suggestion" bson:"suggestion"`     // 是 建议结果 0-通过 1-嫌疑 2-删除
	StartTime    int                `json:"startTime" bson:"startTime"`       // 是	断句开始时间，单位毫秒
	EndTime      int                `json:"endTime" bson:"endTime"`           // 是 断句结束时间，单位毫秒
	Content      string             `json:"content" bson:"content"`           // 是	检测命中内容返回
	Url          string             `json:"url" bson:"url"`                   // 否 需要开启返回直播语音片段播放地址，请联系易盾策略经理
	SpeakerId    string             `json:"speakerId" bson:"speakerId"`       // 否	针对接入SDK监听客户，返回说话人ID信息
	SegmentId    string             `json:"segmentId" bson:"segmentId"`       // 是 断句id
	FrontSegment *FrontAudioSegment `json:"frontSegment" bson:"frontSegment"` // 否 关联证据信息，异常/嫌疑断句命中时返回，命中断句前20s的证据信息，包含前20s的音频语音识别内容及对应url，通过时数据为空
	Labels       []*LiveAudioLabel  `json:"labels" bson:"labels"`             // 否 分类信息，通过时数据为空
}

type LiveAudioLabel struct {
	Label     int                        `json:"label" bson:"label"` // 是 分类信息，100：色情，110：性感低俗，200：广告，210：二维码，260：广告法，300：暴恐，400：违禁，500：涉政，800：恶心类，900：其他，1020：黑屏，1030：挂机，1100：涉价值观
	Level     int                        `json:"level" bson:"level"` // 是 分类级别，0：正常，1：不确定，2：确定
	Rate      float64                    `json:"rate" bson:"rate"`   // 是 置信度分数，0-1之间取值，1为置信度最高，0为置信度最低。若level为正常，置信度越大，说明正常的可能性越高。若level为不确定或确定，置信度越大，说明垃圾的可能性越高
	SubLabels []*LiveAudioDetectSubLabel `json:"subLabels"`          // 是 细分类信息，可能包含多个
}

type LiveAudioDetectSubLabel struct {
	SubLabel string                    `json:"subLabel" bson:"subLabel"` // 是 细分类，详细编码请参考下方对应细分类编码对照表
	Rate     float64                   `json:"rate" bson:"rate"`         // 是 置信度分数，0-1之间取值，1为置信度最高，0为置信度最低
	Details  LiveDetectSubLabelDetails `json:"details" bson:"details"`   // 否 命中的详细对象信息
}

// LiveAudioDetectReviewEvidences 直播墙人审证据信息 否
type LiveAudioDetectReviewEvidences struct {
	Action        int       `json:"action" bson:"action"`               // 是 审核操作, 1为忽略，2为警告，3为断流，4为提示，10为机器检测结束
	ActionTime    int64     `json:"actionTime" bson:"actionTime"`       // 是 操作时间，UNIX_TIME时间戳，毫秒为单位
	SpamType      int       `json:"spamType" bson:"spamType"`           // 是 违规类型, 100-色情, 115-音乐内容低俗, 116-言论低俗涉黄, 121-发出涉黄声音, 122-ASMR, 200-广告, 211-商业推广, 212-提及竞品, 300-暴恐, 400-违禁, 500-涉政, 511-影响政府形象, 512-邪教迷信, 513-涉军事, 514-涉宗教, 515-国歌、严肃歌曲, 516-涉及少数民族,517-涉及政治敏感话题 800-不文明, 826-传播负面情绪, 827-违反公序良俗, 828-侵害他人隐私, 1000-其他, 1023-无营养, 1027-话题炒作, 1030-挂机, 1050-自定义
	SpeakerID     string    `json:"speakerId" bson:"speakerId"`         // 否 针对接入SDK监听客户，针对房间内具体说话人的处罚，返回说话人ID信息，若为空，则针对房间维度处罚
	SpamDetail    string    `json:"spamDetail" bson:"spamDetail"`       // 否 违规详细说明
	CensorAccount string    `json:"censorAccount" bson:"censorAccount"` // 是 审核员账号
	WarnCount     int       `json:"warnCount" bson:"warnCount"`         // 否 警告次数
	PromptCount   int       `json:"promptCount" bson:"promptCount"`     // 是 提示次数
	Segments      []Segment `json:"segments" bson:"segments"`           // 否 人审断句信息, 可为空
}

// Segment 人审断句信息, 可为空
type Segment struct {
	StartTime int `json:"startTime" bson:"startTime"` // 断句开始时间，单位毫秒
	EndTime   int `json:"endTime" bson:"endTime"`     // 断句结束时间 ，单位毫秒
}

// Asr 语音识别检测结果
type Asr struct {
	TaskId    string `json:"taskId" bson:"taskId"`       // 是	检测任务 ID，示例值："38e08da8d2574df4bd2eca9b5153df72"
	StartTime int    `json:"startTime" bson:"startTime"` // 是	断句开始时间，单位毫秒
	EndTime   int    `json:"endTime" bson:"endTime"`     // 是	断句结束时间，单位毫秒
	Content   string `json:"content" bson:"content"`     // 是	语音识别结果
	SpeakerId string `json:"speakerId" bson:"speakerId"` // 否	针对接入SDK监听客户，返回说话人ID信息
}

// Language 语种识别检测结果
type Language struct {
	TaskId    string `json:"taskId" bson:"taskId"`       // 是	检测任务 ID，示例值："38e08da8d2574df4bd2eca9b5153df72"
	StartTime int    `json:"startTime" bson:"startTime"` // 是	断句开始时间，单位毫秒
	EndTime   int    `json:"endTime" bson:"endTime"`     // 是	断句结束时间，单位毫秒
	Content   string `json:"content" bson:"content"`     // 是	语种识别结果
	Callback  string `json:"callback" bson:"callback"`   // 否	提交时传递的callback
	SegmentId string `json:"segmentId" bson:"segmentId"` // 是	断句id
}

type LiveAudioDetectFeedbackResp struct {
	Code   int      `json:"code"`
	Msg    string   `json:"msg"`
	Result []Result `json:"result"`
}

type Result struct {
	TaskID string `json:"taskId"`
	Result int    `json:"result"`
}
