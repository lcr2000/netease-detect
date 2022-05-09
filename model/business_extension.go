package model

import (
	"fmt"
	"net/url"
)

/*
	业务扩展参数
*/

type BusinessExtension struct {
	UserExtension   *UserExtension
	DeviceExtension *DeviceExtension
	SceneExtension  *SceneExtension
	OtherExtension  *OtherExtension
}

// UserExtension 用户扩展参数
// 根据用户信息输出用户画像，并与易盾自有画像库比对，辅助反垃圾结果判定
type UserExtension struct {
	Account      string `json:"account"`      // N	128	用户唯一标识，与易盾账号画像库匹配，建议抄送，辅助机审策略精准调优
	Phone        string `json:"phone"`        // N	64	手机号，与易盾风险库匹配。默认国内手机号，如有海外手机，需包含国家地区代码，格式为“+447410xxx186（+44即为国家码）”。如果需要加密，支持传入hash值，hash算法：md5(phone)
	Nickname     string `json:"nickname"`     // N	128	用户昵称，建议抄送，辅助机审策略精准调优
	Gender       int    `json:"gender"`       // N	4	用户性别，0未知，1男，2女，在社交、直播场景建议抄送，辅助策略精准调优
	Age          int    `json:"age"`          // N	4	用户年龄，0未知，在社交场景建议抄送，辅助策略精准调优
	Level        int    `json:"level"`        // N	4	用户等级，0未知，1初级，2中级，3高级，其他值请与易盾策略约定，建议抄送，辅助策略精准调优
	RegisterTime int64  `json:"registerTime"` // N	13	注册时间，UNIX 时间戳(毫秒值)
	FriendNum    int64  `json:"friendNum"`    // N	20	好友数，在社交、直播场景中使用，建议抄送
	FansNum      int64  `json:"fansNum"`      // N	20	粉丝数，在社交、直播场景中使用，建议抄送
	IsPremiumUse int    `json:"isPremiumUse"` // N	4	是否付费用户，0为默认值，1为付费，建议抄送，易盾将结合该信息综合判断
	Role         string `json:"role"`         // N	32	用户类型角色，可针对不同的角色配置不同的策略。角色与易盾策略约定即可
}

// DeviceExtension 设备扩展参数
// 根据设备信息输出设备画像，并与易盾自有设备画像库比对，辅助反垃圾结果判定
type DeviceExtension struct {
	DeviceID   string `json:"deviceId"`   // N	128	用户设备 id，与易盾设备画像库匹配，明文请转大写传入；MD5加密请明文转大写后MD5计算，再转大写传入，建议抄送
	DeviceType int    `json:"deviceType"` // N	4	用户设备id的类型，0:其他，10:IMEI，11:AndroidID，12:IDFA，13:IDFV，14:MAC ，20:IMEI_MD5，21:AndroidID_MD5，22:IDFA_MD5，23:IDFV_MD5，24:MAC_MD5
	Mac        string `json:"mac"`        // N	128	用户设备 id，与易盾设备画像库匹配，明文请转大写传入；MD5加密请明文转大写后MD5计算，再转大写传入，建议抄送 N	64	用户设备mac信息,与易盾设备画像库匹配，建议抄送
	Imei       string `json:"imei"`       // N	128	用户设备 id，与易盾设备画像库匹配，明文请转大写传入；MD5加密请明文转大写后MD5计算，再转大写传入，建议抄送 N	64	国际移动设备识别码，与易盾设备画像库匹配，建议抄送
	Idfa       string `json:"idfa"`       // N	128	用户设备 id，与易盾设备画像库匹配，明文请转大写传入；MD5加密请明文转大写后MD5计算，再转大写传入，建议抄送N	64	iOS设备标识码，与易盾设备画像库匹配，建议抄送
	Idfv       string `json:"idfv"`       // N	128	用户设备 id，与易盾设备画像库匹配，明文请转大写传入；MD5加密请明文转大写后MD5计算，再转大写传入，建议抄送N	64	iOS设备标识码 ，与易盾设备画像库匹配，建议抄送
	AppVersion string `json:"appVersion"` // N	128	用户设备 id，与易盾设备画像库匹配，明文请转大写传入；MD5加密请明文转大写后MD5计算，再转大写传入，建议抄送N	32	APP版本号 ，可根据不同版本号配置不同策略，建议抄送
}

// SceneExtension 场景扩展参数
// 场景扩展参数，有助于通过业务场景维度辅助反垃圾结果判定
type SceneExtension struct {
	ReceiveUID   string `json:"receiveUid"`   // N	64	接受消息的用户标识，私聊/评论回复场景使用，易盾可根据该id关联检测，辅助机审策略精准调优
	Relationship int    `json:"relationship"` // N	11	收发消息者好友关系，1接收人关注发送人，2发送人关注接收人，3互相关注，4互未关注，私聊/评论场景抄送
	GroupID      string `json:"groupId"`      // N	32	群聊id，群聊场景使用，建议抄送，辅助机审策略精准调优
	RoomID       string `json:"roomId"`       // N	32	聊天室/直播/游戏房间，派对房/直播场景使用，可根据不同的房间设置不同策略，建议抄送，辅助机审策略精准调优
	Topic        string `json:"topic"`        // N	128	文章/帖子id，发帖/动态场景使用，易盾可根据该id，关联检测，辅助机审策略精准调优
	CommentID    string `json:"commentId"`    // N	32	主评论id，围绕主评论盖楼场景使用，建议抄送，辅助机审策略精准调优
	CommodityID  string `json:"commodityId"`  // N	32	商品id，直播卖货/商品介绍场景使用，可根据商品类型设置策略，建议抄送，辅助机审策略精准调优
}

// OtherExtension 其他拓展参数
type OtherExtension struct {
	IP          string `json:"ip"`          // 否	128	用户IP地址，建议抄送，辅助机审策略精准调优
	RelatedKeys string `json:"relatedKeys"` // 否	512	String数组，多个关联Key以逗号分隔（"xxx,xxx"），最多三个Key，单个Key长度不超过128，适用于私聊/评论/跟帖等情况同一用户或不同用户发送多条违规内容关联检测的场景。如需要检测同一评论下的同一用户或不同用户发送违规内容盖楼场景，Key传值方式可以为（"评论ID,用户ID"）
	ExtStr1     string `json:"extStr1"`     // 否	128	自定义扩展参数
	ExtStr2     string `json:"extStr2"`     // 否	128	自定义扩展参数
	ExtLon1     int64  `json:"extLon1"`     // 否	2^63-1	自定义扩展参数
	ExtLon2     int64  `json:"extLon2"`     // 否	2^63-1	自定义扩展参数
}

// SplicingBusinessExpansion 拼接业务拓展参数
func SplicingBusinessExpansion(businessExtension *BusinessExtension, params url.Values) {
	if businessExtension == nil {
		return
	}

	userExtension := businessExtension.UserExtension
	if userExtension != nil {
		if userExtension.Account != "" {
			params["account"] = []string{userExtension.Account}
		}
		if userExtension.Phone != "" {
			params["phone"] = []string{userExtension.Phone}
		}
		if userExtension.Nickname != "" {
			params["nickname"] = []string{userExtension.Nickname}
		}
		if userExtension.Gender > 0 {
			params["gender"] = []string{fmt.Sprintf("%d", userExtension.Gender)}
		}
		if userExtension.Age > 0 {
			params["age"] = []string{fmt.Sprintf("%d", userExtension.Age)}
		}
		if userExtension.Level > 0 {
			params["level"] = []string{fmt.Sprintf("%d", userExtension.Level)}
		}
		if userExtension.RegisterTime > 0 {
			params["registerTime"] = []string{fmt.Sprintf("%d", userExtension.RegisterTime)}
		}
		if userExtension.FriendNum > 0 {
			params["friendNum"] = []string{fmt.Sprintf("%d", userExtension.FriendNum)}
		}
		if userExtension.FansNum > 0 {
			params["fansNum"] = []string{fmt.Sprintf("%d", userExtension.FansNum)}
		}
		if userExtension.IsPremiumUse > 0 {
			params["isPremiumUse"] = []string{fmt.Sprintf("%d", userExtension.IsPremiumUse)}
		}
		if userExtension.Role != "" {
			params["role"] = []string{userExtension.Role}
		}
	}

	deviceExtension := businessExtension.DeviceExtension
	if deviceExtension != nil {
		if deviceExtension.DeviceID != "" {
			params["deviceId"] = []string{deviceExtension.DeviceID}
		}
		if deviceExtension.DeviceType > 0 {
			params["deviceType"] = []string{fmt.Sprintf("%d", deviceExtension.DeviceType)}
		}
		if deviceExtension.Mac != "" {
			params["mac"] = []string{deviceExtension.Mac}
		}
		if deviceExtension.Imei != "" {
			params["imei"] = []string{deviceExtension.Imei}
		}
		if deviceExtension.Idfa != "" {
			params["idfa"] = []string{deviceExtension.Idfa}
		}
		if deviceExtension.Idfv != "" {
			params["idfv"] = []string{deviceExtension.Idfv}
		}
		if deviceExtension.AppVersion != "" {
			params["appVersion"] = []string{deviceExtension.AppVersion}
		}
	}

	sceneExtension := businessExtension.SceneExtension
	if sceneExtension != nil {
		if sceneExtension.ReceiveUID != "" {
			params["receiveUid"] = []string{sceneExtension.ReceiveUID}
		}
		if sceneExtension.Relationship > 0 {
			params["relationship"] = []string{fmt.Sprintf("%d", sceneExtension.Relationship)}
		}
		if sceneExtension.GroupID != "" {
			params["groupId"] = []string{sceneExtension.GroupID}
		}
		if sceneExtension.RoomID != "" {
			params["roomId"] = []string{sceneExtension.RoomID}
		}
		if sceneExtension.Topic != "" {
			params["topic"] = []string{sceneExtension.Topic}
		}
		if sceneExtension.CommentID != "" {
			params["commentId"] = []string{sceneExtension.CommentID}
		}
		if sceneExtension.CommodityID != "" {
			params["commodityId"] = []string{sceneExtension.CommodityID}
		}
	}

	otherExtension := businessExtension.OtherExtension
	if otherExtension != nil {
		if otherExtension.IP != "" {
			params["ip"] = []string{otherExtension.IP}
		}
		if otherExtension.RelatedKeys != "" {
			params["relatedKeys"] = []string{otherExtension.RelatedKeys}
		}
		if otherExtension.ExtStr1 != "" {
			params["extStr1"] = []string{otherExtension.ExtStr1}
		}
		if otherExtension.ExtStr2 != "" {
			params["extStr2"] = []string{otherExtension.ExtStr2}
		}
		if otherExtension.ExtLon1 > 0 {
			params["extLon1"] = []string{fmt.Sprintf("%d", otherExtension.ExtLon1)}
		}
		if otherExtension.ExtLon2 > 0 {
			params["extLon2"] = []string{fmt.Sprintf("%d", otherExtension.ExtLon2)}
		}
	}
}
