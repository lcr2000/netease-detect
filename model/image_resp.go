package model

type ImageDetectResp struct {
	BaseResp
	AntiSpam []AntiSpam `json:"antispam"`
	Ocr      []Ocr      `json:"ocr"`
	Face     []Face     `json:"face"`
	Quality  []Quality  `json:"quty"`
	Logo     []Logo     `json:"logo"`
	Scene    []Scene    `json:"scene"`
}

type ImageDetectResult struct {
	AntiSpam []AntiSpam `json:"antispam"`
	Face     []Face     `json:"face"`
	Scene    []Scene    `json:"scene"`
}

type AntiSpam struct {
	TaskID           string        `json:"taskId"`
	Status           int           `json:"status"` // 0：检测成功，610：图片下载失败，620：图片格式错误，630：其它
	Action           int           `json:"action"`
	CensorType       int           `json:"censorType"`
	StrategyVersions string        `json:"strategyVersions"`
	Name             string        `json:"name"`
	ImageLabels      []ImageLabels `json:"labels"`
}

type ImageLabels struct {
	Label     int              `json:"label"`
	SubLabels []*ImageSubLabel `json:"subLabels"`
	Level     int              `json:"level"` // 0：正常，1：不确定，2：确定
	Rate      float64          `json:"rate"`
}

type ImageSubLabel struct {
	SubLabel int               `json:"subLabel"`
	Rate     float64           `json:"rate"`
	Details  *AntiImageDetails `json:"details"`
}

type AntiImageDetails struct {
	HitInfos      []string `json:"hitInfos"`
	AnticheatInfo string   `json:"anticheatInfo"`
	ImageTagInfos []struct {
		TagName  string `json:"tagName"`
		TagGroup string `json:"tagGroup"`
	} `json:"imageTagInfos"`
	ImageListInfos []struct {
		Type     string `json:"type"`
		Url      string `json:"url"`
		HitCount string `json:"hitCount"`
		Word     string `json:"word"`
	} `json:"imageListInfos"`
	HitLocationInfos []struct {
		HitInfo string  `json:"hitInfo"`
		X1      float64 `json:"x1"`
		Y1      float64 `json:"y1"`
		X2      float64 `json:"x2"`
		Y2      float64 `json:"y2"`
	} `json:"hitLocationInfos"`
}

type Ocr struct {
	Name    string       `json:"name"`
	Height  int          `json:"height"`
	Width   int          `json:"width"`
	TaskID  string       `json:"taskId"`
	Details []OcrDetails `json:"details"`
}

type OcrDetails struct {
	Content      string         `json:"content"`
	LineContents []LineContents `json:"lineContents"`
}

type LineContents struct {
	LineContent string    `json:"lineContent"`
	Polygon     []float64 `json:"polygon"`
}

type Face struct {
	Name    string       `json:"name"`
	TaskID  string       `json:"taskId"`
	Details []FaceDetail `json:"details"`
}

type FaceDetail struct {
	FaceNumber   int            `json:"faceNumber"`
	FaceContents []FaceContents `json:"faceContents"`
}

type FaceContents struct {
	Name     string  `json:"name"`
	Gender   string  `json:"gender"` // 男（male）、女（female）；不可识别则为空
	Age      int     `json:"age"`
	Type     string  `json:"type"`
	Category string  `json:"category"`
	X1       float64 `json:"x1"`
	Y1       float64 `json:"y1"`
	X2       float64 `json:"x2"`
	Y2       float64 `json:"y2"`
}

type Quality struct {
	Name    string           `json:"name"`
	TaskID  string           `json:"taskId"`
	Details []QualityDetails `json:"details"`
}

type BoarderInfo struct {
	Hit    bool `json:"hit"`
	Top    bool `json:"top"`
	Right  bool `json:"right"`
	Bottom bool `json:"bottom"`
	Left   bool `json:"left"`
}

type BackgroundInfo struct {
	PureBackground bool `json:"pureBackground"`
}

type QualityDetails struct {
	AestheticsRate float64        `json:"aestheticsRate"`
	MetaInfo       MetaInfo       `json:"metaInfo"`
	BoarderInfo    BoarderInfo    `json:"boarderInfo"`
	BackgroundInfo BackgroundInfo `json:"backgroundInfo"`
}

type MetaInfo struct {
	ByteSize int    `json:"byteSize"`
	Height   int    `json:"height"`
	Width    int    `json:"width"`
	Format   string `json:"format"`
}

type Logo struct {
	TaskID  string        `json:"taskId"`
	Name    string        `json:"name"`
	Details []interface{} `json:"details"`
}

type Scene struct {
	TaskID  string `json:"taskId"`
	Name    string `json:"name"`
	Details []struct {
		SceneName string  `json:"sceneName"`
		Rate      float64 `json:"rate"`
	} `json:"details"`
}

type ImageHitInfo struct {
	Labels    map[int]*ImageHitLabelInfo `json:"labels"`
	TaskId    string                     `json:"task_id"`
	ImageName string                     `json:"image_name"`
	Data      string                     `json:"data"`   // 图片url
	Action    int                        `json:"action"` // 图片url
}

type ImageHitLabelInfo struct {
	Label     int                           `json:"label"`
	LabelName string                        `json:"label_name"`
	Level     int                           `json:"level"`
	SubLabels map[int]*ImageHitSubLabelInfo `json:"sub_labels"`
}

type ImageHitSubLabelInfo struct {
	SubLabel        int               `json:"label"`
	SubLabelDetails *AntiImageDetails `json:"sub_details"`
}
