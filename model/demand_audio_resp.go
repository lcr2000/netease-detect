package model

// VideoCheckSubmitResp 视频点播Response
type AudioCheckSubmitResponse struct {
	BaseResp
	Result *AudioCheckSubmitResult `json:"result"`
}

type AudioCheckSubmitResult struct {
	TaskId       string `json:"taskId"`
	Status       int    `json:"status"` // 检测结果，0：成功，1：失败
	DealingCount int    `json:"dealingCount"`
}

type AudioDetectResp struct {
	BaseResp
	AntiSpam []*AudioAntiSpam `json:"antispam"`
}

type AudioAntiSpam struct {
	TaskId    string        `json:"taskId"`
	AsrStatus int           `json:"asrStatus"`
	AsrResult int           `json:"asrResult"`
	Action    int           `json:"action"` // 0：通过，1：嫌疑，2：不通过
	Labels    []*AudioLabel `json:"labels"`
	Duration  int           `json:"duration"` // 音频时长，单位s
	Callback  string        `json:"callback"`
}

type AudioLabel struct {
	Label     int              `json:"label"` // 100：色情，200：广告，260：广告法，300：暴恐，400：违禁，500：涉政，600：谩骂，1100：涉价值观
	Level     int              `json:"level"` // 0：通过，1：嫌疑，2：不通过
	SubLabels []*AudioSubLabel `json:"subLabels"`
}

type AudioSubLabel struct {
	SubLabel string                `json:"subLabel"`
	Details  *AudioSubLabelDetails `json:"details"`
}

type AudioSubLabelDetails struct {
	HitType  int                     `json:"hitType"`
	Hint     []*AudioSubLabelHint    `json:"hint"`
	HitInfos []*AudioSubLabelHitInfo `json:"hitInfos"`
}

type AudioSubLabelHitInfo struct {
	HitType  int    `json:"hitType"`
	HitClues string `json:"hitClues"`
}

type AudioSubLabelHint struct {
	Value string `json:"value"`
}

type AudioHitLabel struct {
	Label     int                       `json:"label"` // 100：色情，200：广告，260：广告法，300：暴恐，400：违禁，500：涉政，600：谩骂，1100：涉价值观
	LabelName string                    `json:"label_name"`
	SubLabels map[int]*AudioHitSubLabel `json:"subLabels"`
}

type AudioHitSubLabel struct {
	SubLabel     int                   `json:"subLabel"`
	SubLabelName string                `json:"sub_label_name"`
	Details      *AudioSubLabelDetails `json:"details"`
}
