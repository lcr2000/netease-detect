package model

type VideoDetectSubmitReq struct {
	URL         string // Y	512	点播视频地址
	DataID      string // Y	128	点播视频唯一标识
	Version     string // Y	4	接口版本号，可选值 v3.2
	Title       string // N	512	视频名称
	Callback    string // N	512	数据回调参数，产品根据业务情况自行设计，当获取离线检测结果时，易盾内容安全服务会返回该字段
	CallbackURL string // N	256	离线结果回调通知到客户的URL。主动回调数据接口超时时间设置为2s，为了保证顺利接收数据，需保证接收接口性能稳定并且保证幂等性。
	UniqueKey   string // N	64	客户个性化视频唯一性标识，传入后，将以此值作为重复检测依据，若不传，默认以URL作为查重依据,如果重复提交会被拒绝，返回报错信息请求重复，以及原提交taskID值，具体返回请查看响应示例
	SubProduct  string // N	32	业务结算id(自定义),业务方传入参数，用于资源账单统计,如需开启请联系易盾客户经理。开启之后资源账单按套餐结算sheet增加subProduct维度的数据请求量、检测量、消耗金额的统计
	Account     string // 用户扩展参数 N	128	用户唯一标识，与易盾账号画像库匹配，建议抄送，辅助机审策略精准调优

	// ----------------截帧频率指定参数---------------
	AdvancedFrequency string // json字符串	N	128	高级截帧设置，此项填写，默认截帧策略失效
}

// AdvancedFrequency advancedFrequency字符串数据结构
type AdvancedFrequency struct {
	DurationPoints []int `json:"durationPoints"` //  数组  Y    5   视频时长区间分割，单位为秒
	Frequencies    []int `json:"frequencies"`    // 数组 Y    6   视频时长区间对应的截帧频率，可设置范围为0.5~600秒
}
