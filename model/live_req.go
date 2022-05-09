package model

type LiveDetectSubmitReq struct {
	Version           string `json:"version"`           // 是	4	接口版本号，值为v3
	URL               string `json:"url"`               // 是	1024 直播流地址
	DataID            string `json:"dataId"`            // 是	128	数据唯一标识，能够根据该值定位到该条数据，如对数据结果有异议，可以发送该值给客户经理查询
	Title             string `json:"title"`             // 否	512	直播语音标题
	IP                string `json:"ip"`                // 否	128	用户IP地址
	Account           string `json:"account"`           // 否	128	账号ID/主播ID
	RoomNo            string `json:"roomNo"`            // 否	128	主播房间号
	Age               string `json:"age"`               // 否	64	主播年龄
	AccountLevel      string `json:"accountLevel"`      // 否	10	账号级别/主播级别(大写A-Z)
	AccountName       string `json:"accountName"`       // 否	30	账号名称/主播名称
	DeviceID          string `json:"deviceId"`          // 否	128	用户设备 id，与易盾设备画像库匹配，明文请转大写传入；MD5加密请明文转大写后MD5计算，再转大写传入，建议抄送
	DeviceType        int    `json:"deviceType"`        // 否	4	用户设备id的类型，0:其他，10:IMEI，11:AndroidID，12:IDFA，13:IDFV，14:MAC ，20:IMEI_MD5，21:AndroidID_MD5，22:IDFA_MD5，23:IDFV_MD5，24:MAC_MD5
	Livelink          string `json:"livelink"`          // 否	128	传入该路直播的前台直播路径(必须使用https协议)，主要用于SaaS审核系统-直播音视频查询页面和审核页面增加前台跳转链接功能
	ScreenMode        int    `json:"ScreenMode"`        // 否	1	0-竖屏，1-横屏，默认竖屏
	DetectType        int    `json:"detectType"`        // 否	1	机器过检类型，0：直播视频与直播音频同时过检，1：仅过检直播视频，2：仅过检音频，不填写默认为0，请注意
	LabourUnion       string `json:"labourUnion"`       // 否	256	主播所属工会
	OperationManager  string `json:"operationManager"`  // 否	128	运营管理者
	ScFrequency       int    `json:"scFrequency"`       // 否	4	截图检测频率，默认5秒截图检测一次，可设置范围为0.5~600秒
	Callback          string `json:"callback"`          // 否	1024	数据回调参数，调用方根据业务情况自行设计，当调用callback接口获取直播音视频结果时，该接口会原样返回该字段，详细见直播音视频离线检测结果获取。作为数据处理标识，因此该字段应该设计为能唯一定位到该次请求的数据结构，如对用户的昵称进行检测，dataId可设为用户标识（用户ID），用户修改多次，每次请求数据的dataId可能一致，但是callback参数可以设计成定位该次请求的数据结构，比如callback字段设计成json，包含dataId和请求的时间戳等信息，当然如果不想做区分，也可以直接把callback设置成dataId的值。
	CallbackURL       string `json:"callbackUrl"`       // 否	512	离线结果回调通知到客户的URL。主动回调数据接口超时时间设置为2s，为了保证顺利接收数据，需保证接收接口性能稳定并且保证幂等性。
	UniqueKey         string `json:"uniqueKey"`         // 否	256	客户个性化直播流唯一性标识，传入后，将以此值作为重复检测依据，若不传，默认以URL作为查重依据
	WallHidden        int    `json:"wallHidden"`        // 否	1	是否从审核墙上隐藏，0-否, 1-是
	CheckLanguageCode string `json:"checkLanguageCode"` // 否	2	指定语言检测语音内容，不填以后台配置为准；zh:中文,en:英语,it:意大利语,id:印尼语,es:西班牙语,fr:法语,ms:马来语,tl:菲利宾语,th:泰语,de:德语,hi:印地语,ru:俄语,ar:阿拉伯语,vi:越南语
}

type LiveDetectTaskInfo struct {
	TaskID string `json:"taskId"` // 待停止检测的直播音视频taskId，示例值："38e08da8d2574df4bd2eca9b5153df72"
	Status int    `json:"status"` // 直播音视频检测状态, 100-停止检测
}
