package model

type LiveAudioDetectSubmitReq struct {
	Version           string `json:"version"`           // 是	4	接口版本号，值为v4
	Url               string `json:"url"`               // 是	1024 使用拉流检测为直播流地址；使用SDK检测请根据不同服务商生成相应的URL
	DataId            string `json:"dataId"`            // 是	128	数据唯一标识，能够根据该值定位到该条数据，如对数据结果有异议，可以发送该值给客户经理查询
	Title             string `json:"title"`             // 否	512	直播语音标题
	Ip                string `json:"ip"`                // 否	128	用户IP地址
	Account           string `json:"account"`           // 否	128	用户唯一标识，如果无需登录则为空
	RoomNo            string `json:"roomNo"`            // 否	128	主播房间号
	AccountLevel      string `json:"accountLevel"`      // 否	10	账号级别/主播级别(大写A-Z)
	AccountName       string `json:"accountName"`       // 否	30	账号名称/主播名称
	DeviceId          string `json:"deviceId"`          // 否	128	用户设备 id，与易盾设备画像库匹配，明文请转大写传入；MD5加密请明文转大写后MD5计算，再转大写传入，建议抄送
	DeviceType        int    `json:"deviceType"`        // 否	4   用户设备id的类型，0:其他，10:IMEI，11:AndroidID，12:IDFA，13:IDFV，14:MAC ，20:IMEI_MD5，21:AndroidID_MD5，22:IDFA_MD5，23:IDFV_MD5，24:MAC_MD5
	Callback          string `json:"callback"`          // 否	1024 据回调参数，调用方根据业务情况自行设计，当调用callback接口获取直播音频结果时，该接口会原样返回该字段，详细见直播音频离线检测结果获取。作为数据处理标识，因此该字段应该设计为能唯一定位到该次请求的数据结构，如对用户的昵称进行检测，dataId可设为用户标识（用户ID），用户修改多次，每次请求数据的dataId可能一致，但是callback参数可以设计成定位该次请求的数据结构，比如callback字段设计成json，包含dataId和请求的时间戳等信息，当然如果不想做区分，也可以直接把callback设置成dataId的值。
	CallbackUrl       string `json:"callbackUrl"`       // 否	512	 离线结果回调通知到客户的URL。主动回调数据接口超时时间设置为2s，为了保证顺利接收数据，需保证接收接口性能稳定并且保证幂等性。
	UniqueKey         string `json:"uniqueKey"`         // 否	256	客户个性化直播流唯一性标识，传入后，将以此值作为重复检测依据，若不传，默认以URL作为查重依据
	CheckLanguageCode string `json:"checkLanguageCode"` // 否	2	指定语言检测音频内容，需以易盾标准填入，可选范围：zh:中文,en:英语,it:意大利语,id:印尼语,es:西班牙语,fr:法语,ms:马来语,tl:菲利宾语,th:泰语,de:德语,hi:印地语,ru:俄语,ar:阿拉伯语,vi:越南语,ja:日语,ko:韩语；不填以后台配置为准
}

type LiveAudioDetectFeedback struct {
	TaskId string `json:"taskId"` // 待停止检测的直播音视频taskId，示例值："38e08da8d2574df4bd2eca9b5153df72"
	Status int    `json:"status"` // 直播音视频检测状态, 100-停止检测
}

// FeedbackReq 反馈接口请求体
type FeedbackReq struct {
	TaskId string `json:"taskId"` // 是	32	音频taskId，示例值："38e08da8d2574df4bd2eca9b5153df72"
	Level  int    `json:"level"`  // 是	3	数据级别，0：正常，2：确定
	Label  int    `json:"label"`  // 否	4	分类信息，100：色情，200：广告，260：广告法，300：暴恐，400：违禁，500：涉政，600：谩骂，700：灌水，900：其他，1100：涉价值观，当返回level是正常时，label字段可不传
}
