package model

type ImageDetectReq struct {
	Images []*ImageInfo `json:"images"` // String(ImageInfo json数组)	Y	32张或10MB	images为json数组，支持批量检测
	// 业务拓展参数
	*BusinessExtension
	Ip          string   `json:"ip"`          // N	128	用户IP地址
	CheckLabels []string `json:"checkLabels"` // String数组 // N	64	接口指定过检分类，可多选，过检分类列表：100：色情，110：性感低俗，200：广告，210：二维码，260：广告法，300：暴恐，400：违禁，500：涉政，800：恶心类，900：其他，1100：涉价值观
	SubProduct  string   `json:"subProduct"`  // N	32	业务结算id(自定义),业务方传入参数，用于资源账单统计,如需开启请联系易盾客户经理。开启之后资源账单按套餐结算sheet增加subProduct维度的数据请求量、检测量、消耗金额的统计
	Extension   string   `json:"extension"`   // N	30000	自定义扩展参数
}

type ImageInfo struct {
	Name        string `json:"name"`        // Y	1024	图片名称(或图片标识)， 该字段为回调信号字段，产品可以根据业务情况自行设计，如json结构、或者为图片url均可
	ImageType   int    `json:"type"`        // Y	4	类型，分别为1：图片URL，2:图片BASE64值
	Data        string `json:"data"`        // Y	32张或者10MB	图片内容，如type=1，则该值为图片URL，图片URL检测单次请求最多支持32张。如type=2，则该值为图片Base64值，转换后的base64内容以及所有请求参数大小不超过10M
	CallbackUrl string `json:"callbackUrl"` // N	256	离线结果回调通知到客户的URL。主动回调数据接口超时时间设置为2s，为了保证顺利接收数据，需保证接收接口性能稳定并且保证幂等性。
}
