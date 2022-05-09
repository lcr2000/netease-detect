package model

type AudioDetectSubmitReq struct {
	Url         string // Y	512	语音文件url
	Version     string // Y	4	接口版本号，可选值 v3.5
	Title       string // N	512	文件标题
	Ip          string // N	32	用户IP地址
	Account     string // N	128	用户唯一标识，如果无需登录则为空
	DeviceId    string // N	128	用户设备 id，与易盾设备画像库匹配，明文请转大写传入；MD5加密请明文转大写后MD5计算，再转大写传入，建议抄送
	DeviceType  int    // N	4	用户设备id的类型，0:其他，10:IMEI，11:AndroidID，12:IDFA，13:IDFV，14:MAC ，20:IMEI_MD5，21:AndroidID_MD5，22:IDFA_MD5，23:IDFV_MD5，24:MAC_MD5
	Callback    string // N	2^16-1	数据回调参数，调用方根据业务情况自行设计，当调用离线结果获取接口时，该接口会原样返回该字段，详细见音频离线检测结果获取。作为数据处理标识，因此该字段应该设计为能唯一定位到该次请求的数据结构，如对用户的昵称进行检测，dataId可设为用户标识（用户ID），用户修改多次，每次请求数据的dataId可能一致，但是callback参数可以设计成定位该次请求的数据结构，比如callback字段设计成json，包含dataId和请求的时间戳等信息，当然如果不想做区分，也可以直接把callback设置成dataId的值。
	CallbackUrl string // N	256	离线结果回调通知到客户的URL。主动回调数据接口超时时间设置为2s，为了保证顺利接收数据，需保证接收接口性能稳定并且保证幂等性。
	SubProduct  string // N	32	业务结算id(自定义),业务方传入参数，用于资源账单统计,如需开启请联系易盾客户经理。开启之后资源账单按套餐结算sheet增加subProduct维度的数据请求量、检测量、消耗金额的统计
}

type AudioDetectReq struct {
	DataCheckType int      // 否	1	检测类型，默认0-url，1-语音内容base64
	Url           string   // 否	1024	音频下载url，当检测类型是url时，url必填
	Data          string   // 否	2560000	语音内容base64，当检测类型是base64时，data必填
	Title         string   // 否	512	音频标题
	DataId        string   // 否	128	数据唯一标识，能够根据该值定位到该条数据，如对数据结果有异议，可以发送该值给客户经理查询
	Callback      string   // 否	65535	数据回调参数，调用方根据业务情况自行设计，当调用离线结果获取接口或查询接口时，该接口会原样返回该字段，详细见音频离线检测结果获取音频轮询模式结果获取。作为数据处理标识，因此该字段应该设计为能唯一定位到该次请求的数据结构，如对用户的昵称进行检测，dataId可设为用户标识（用户ID），用户修改多次，每次请求数据的dataId可能一致，但是callback参数可以设计成定位该次请求的数据结构，比如callback字段设计成json，包含dataId和请求的时间戳等信息，当然如果不想做区分，也可以直接把callback设置成dataId的值。
	CallbackUrl   string   // 否	512	离线结果回调通知到客户的URL。主动回调数据接口超时时间默认设置为2s，为了保证顺利接收数据，需保证接收接口性能稳定并且保证幂等性。
	PublishTime   int64    // 否	13	发表时间，UNIX 时间戳(毫秒值)
	Nickname      string   // 否	128	用户昵称，建议抄送，辅助机审策略精准调优
	Ip            string   // 否	128	用户IP地址，建议抄送，辅助机审策略精准调优
	Account       string   // 否	128	用户唯一标识，与易盾账号画像库匹配，建议抄送，辅助机审策略精准调优
	DeviceId      string   // 否	128	用户设备id，与易盾设备画像库匹配，明文请转大写传入；MD5加密请明文转大写后MD5计算，再转大写传入，建议抄送
	DeviceType    int      // 否	4	用户设备id的类型，0:其他，10:IMEI，11:AndroidID，12:IDFA，13:IDFV，14:MAC ，20:IMEI_MD5，21:AndroidID_MD5，22:IDFA_MD5，23:IDFV_MD5，24:MAC_MD5，可选值-10000~10000
	Extension     string   // 否	35000	自定义扩展参数，json字符串
	SubProduct    string   // 否	32	业务结算id(自定义),业务方传入参数，用于资源账单统计,如需开启请联系易盾客户经理。开启之后资源账单按套餐结算sheet增加subProduct维度的数据请求量、检测量、消耗金额的统计
	UniqueKey     string   // 否	256	检测去重字段，填入后使用此字段去重，如果没填入则默认使用url进行去重
	RelatedKeys   []string // 否	3个且单个长度最大128	上下文关联检测key
	// 业务拓展参数
	*BusinessExtension
}
