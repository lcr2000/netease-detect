package model

type AudioCheckSubmitResponse struct {
	BaseResp
	Result *AudioCheckSubmitResult `json:"result"`
}

type AudioCheckSubmitResult struct {
	TaskID       string `json:"taskId"`
	Status       int    `json:"status"` // 检测结果，0：成功，1：失败
	DealingCount int    `json:"dealingCount"`
}

type AudioDetectResp struct {
	BaseResp
	AntiSpam []*AudioAntiSpam `json:"antispam"`
}

type AudioAntiSpam struct {
	TaskID    string        `json:"taskId"`
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

// AudioSyncDetectResp 点播音频同步检测
type AudioSyncDetectResp struct {
	Code   int          `json:"code"`
	Msg    string       `json:"msg"`
	Result *AudioResult `json:"result"`
}

type AudioResult struct {
	Antispam *Antispam      `json:"antispam"`
	Language *AudioLanguage `json:"language"`
	Asr      *AudioAsr      `json:"asr"`
	Voice    *Voice         `json:"voice"`
}

type AudioAsr struct {
	TaskID   string            `json:"taskId"`
	DataID   string            `json:"dataId"`
	Callback string            `json:"callback"`
	Details  []*AudioAsrDetail `json:"details"`
}

type AudioAsrDetail struct {
	StartTime int    `json:"startTime"`
	EndTime   int    `json:"endTime"`
	Content   string `json:"content"`
}

type Voice struct {
	TaskID   string `json:"taskId"`
	DataID   string `json:"dataId"`
	Callback string `json:"callback"`
	Detail   struct {
		MainGender string `json:"mainGender"`
	} `json:"detail"`
}

type Antispam struct {
	TaskID        string        `json:"taskId"`
	Status        int           `json:"status"`
	FailureReason int           `json:"failureReason"`
	Suggestion    int           `json:"suggestion"`
	ResultType    int           `json:"resultType"`
	Segments      []*Segments   `json:"segments"`
	Callback      string        `json:"callback"`
	DataID        string        `json:"dataId"`
	CensorSource  int           `json:"censorSource"`
	Duration      int           `json:"duration"`
	CensorTime    int64         `json:"censorTime"`
	CustomAction  int           `json:"customAction"`
	CensorLabels  []interface{} `json:"censorLabels"`
}

type Segments struct {
	StartTime int      `json:"startTime"`
	EndTime   int      `json:"endTime"`
	Content   string   `json:"content"`
	Labels    []*Label `json:"labels"`
}

type Label struct {
	Label     int         `json:"label"`
	Level     int         `json:"level"`
	SubLabels []*SubLabel `json:"subLabels,omitempty"`
}

type SubLabel struct {
	SubLabel string  `json:"subLabel"`
	Details  *Detail `json:"details"`
}

type Detail struct {
	HitInfos []*HitInfo `json:"hitInfos"`
	Keywords []*Keyword `json:"keywords"`
	LibInfos []*LibInfo `json:"libInfos"`
}

type HitInfo struct {
	Value string `json:"value"`
}

type Keyword struct {
	Word string `json:"word"`
}

type LibInfo struct {
	ListType int    `json:"listType"`
	Entity   string `json:"entity"`
}

type AudioLanguage struct {
	TaskID   string         `json:"taskId"`
	DataID   string         `json:"dataId"`
	Callback string         `json:"callback"`
	Details  []*AudioDetail `json:"details"`
}

type AudioDetail struct {
	Type     string     `json:"type"`
	Segments []*Segment `json:"segments"`
}

type AudioSegment struct {
	StartTime int `json:"startTime"`
	EndTime   int `json:"endTime"`
}
