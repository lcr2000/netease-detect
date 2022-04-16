package model

// VideoCheckSubmitResp 视频点播Response
type VideoCheckSubmitResp struct {
	BaseResp
	Result VideoCheckSubmitResult `json:"result"`
}

type VideoCheckSubmitResult struct {
	TaskId       string `json:"taskId"`
	Status       int    `json:"status"` // 检测结果，0：成功，1：失败
	DealingCount int    `json:"dealingCount"`
}

type VideoDetectResp struct {
	BaseResp
	Result []*VideoDetectResult `json:"result"`
}

type VideoDetectResult struct {
	Status       int                         `json:"status"` // 视频处理状态码，定义为：0：检测成功，110：请求重复，120：参数错误，130：解析错误，140：数据类型错误，160：视频大小超限（>5G）
	TaskId       string                      `json:"taskId"`
	CensorSource int                         `json:"censorSource"`
	CensorTime   int64                       `json:"censorTime"`
	Level        int                         `json:"level"`
	Evidences    []VideoDetectResultEvidence `json:"evidences"`
	Callback     string                      `json:"callback"`
}

type VideoDetectResultEvidence struct {
	BeginTime int            `json:"beginTime"` // 证据开始相对时间，单位为毫秒 如：149000 转换为"00:02:29"
	EndTime   int            `json:"endTime"`   // 证据结束相对时间，单位为毫秒
	Type      int            `json:"type"`      // 1：图片，2：视频
	Url       string         `json:"url"`
	Labels    []*VideoLabels `json:"labels"`
}

type VideoLabels struct {
	Label     int `json:"label"`
	SubLabels []struct {
		SubLabel int               `json:"subLabel"`
		Rate     float64           `json:"rate"`
		Details  *AntiVideoDetails `json:"details"`
	} `json:"subLabels"`
	Level int     `json:"level"` // 0：正常，1：不确定，2：确定
	Rate  float64 `json:"rate"`
}

type AntiVideoDetails struct {
	HitInfos      []string `json:"hitInfos"`
	AnticheatInfo string   `json:"anticheatInfo"`
	ImageTagInfo  []struct {
		TagName  string `json:"tagName"`
		TagGroup string `json:"tagGroup"`
	} `json:"imageTagInfo"`
	ImageListInfo []struct {
		Type     string `json:"type"`
		Url      string `json:"url"`
		HitCount string `json:"hitCount"`
		Word     string `json:"word"`
	} `json:"imageListInfo"`
	HitLocationInfos []struct {
		HitInfo string  `json:"hitInfo"`
		X1      float64 `json:"x1"`
		Y1      float64 `json:"y1"`
		X2      float64 `json:"x2"`
		Y2      float64 `json:"y2"`
	} `json:"hitLocationInfos"`
}

// DemandVideoCallbackResp 点播视频检测结果回调返回值(跟http调用返回的数据结构不一样)
type DemandVideoCallbackResp struct {
	TaskID       string                         `json:"taskId"`
	Callback     string                         `json:"callback"`
	Status       int                            `json:"status"`
	Level        int                            `json:"level"`
	CensorSource int                            `json:"censorSource"`
	CensorTime   int64                          `json:"censorTime"`
	Duration     int                            `json:"duration"`
	Evidences    []*DemandVideoCallbackEvidence `json:"evidences"`
}

type DemandVideoCallbackEvidence struct {
	Type         int                           `json:"type"`
	URL          string                        `json:"url"`
	BeginTime    int                           `json:"beginTime"`
	EndTime      int                           `json:"endTime"`
	Labels       []*DemandVideoCallbackLabel   `json:"labels"`
	CensorSource int                           `json:"censorSource"`
	BackPics     []*DemandVideoCallbackBackPic `json:"backPics"`
}

type DemandVideoCallbackLabel struct {
	Label     int                            `json:"label"`
	Level     int                            `json:"level"`
	Rate      float64                        `json:"rate"`
	SubLabels []*DemandVideoCallbackSubLabel `json:"subLabels"`
}

type DemandVideoCallbackSubLabel struct {
	SubLabel int                        `json:"subLabel"`
	Rate     float64                    `json:"rate"`
	Details  *DemandVideoCallbackDetail `json:"details"`
}

type DemandVideoCallbackDetail struct {
	HitInfos         []string                              `json:"hitInfos"`
	ImageTagInfos    []*DemandVideoCallbackImageTagInfo    `json:"imageTagInfos"`
	HitLocationInfos []*DemandVideoCallbackHitLocationInfo `json:"hitLocationInfos"`
}

type DemandVideoCallbackImageTagInfo struct {
	TagName  string `json:"tagName"`
	TagGroup string `json:"tagGroup"`
}

type DemandVideoCallbackHitLocationInfo struct {
	HitInfo string  `json:"hitInfo"`
	X1      float64 `json:"x1"`
	Y1      float64 `json:"y1"`
	X2      float64 `json:"x2"`
	Y2      float64 `json:"y2"`
}

type DemandVideoCallbackBackPic struct {
	URL string `json:"url"`
}
