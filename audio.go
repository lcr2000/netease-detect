package netease

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	"github.com/lcr2000/netease-detect/model"
	"github.com/spf13/cast"
)

/*
	点播音频检测文档地址
	https://support.dun.163.com/documents/2018082201?docId=191344157972942848
*/

// AudioDetectSubmit 提交点播音频异步检测
func (c *Client) AudioDetectSubmit(req *model.AudioDetectSubmitReq) (rsp *model.AudioCheckSubmitResponse, err error) {
	if req == nil || req.URL == "" {
		err = errors.New("params is required")
		return
	}

	params := url.Values{
		"url":     []string{req.URL},
		"version": []string{"v3.5"},
	}
	if req.Title != "" {
		params["title"] = []string{req.Title}
	}
	if req.IP != "" {
		params["ip"] = []string{req.IP}
	}
	if req.Account != "" {
		params["account"] = []string{req.Account}
	}
	if req.DeviceID != "" {
		params["deviceId"] = []string{req.DeviceID}
	}
	if req.DeviceType != 0 {
		params["deviceType"] = []string{fmt.Sprintf("%d", req.DeviceType)}
	}
	if req.Callback != "" {
		params["callback"] = []string{req.Callback}
	}
	if req.CallbackURL != "" {
		params["callbackUrl"] = []string{req.CallbackURL}
	}
	if req.SubProduct != "" {
		params["subProduct"] = []string{req.SubProduct}
	}

	bytes, err := c.Request(AudioSubmitURL, "v3.5", params)
	if err != nil {
		return
	}

	rsp = &model.AudioCheckSubmitResponse{}
	if err = json.Unmarshal(bytes, &rsp); err != nil {
		return
	}
	if rsp.Code != CallSuccessCode {
		err = fmt.Errorf("AudioDetectSubmit fail, code=%v", rsp.Code)
		return
	}

	return
}

// GetAudioDetectResult 获取点播音频异步检测结果
func (c *Client) GetAudioDetectResult() (rsp *model.AudioDetectResp, err error) {
	params := url.Values{}
	body, err := c.Request(AudioResultURL, "v3.5", params)
	if err != nil {
		return
	}
	rsp = &model.AudioDetectResp{}
	if err = json.Unmarshal(body, &rsp); err != nil {
		return
	}
	if rsp.Code != CallSuccessCode {
		err = fmt.Errorf("GetAudioDetectResult fail, code=%v", rsp.Code)
		return
	}
	return
}

// AudioDetect 点播音频同步检测
func (c *Client) AudioDetect(req *model.AudioDetectReq) (rsp *model.AudioSyncDetectResp, err error) {
	if req == nil {
		err = errors.New("params is required")
		return
	}
	if req.URL == "" && req.Data == "" {
		err = errors.New("params is required")
		return
	}
	if req.DataCheckType != 0 && req.DataCheckType != 1 {
		err = errors.New("data check type params error")
		return
	}

	params := url.Values{
		"version":       []string{APIVersionV2},
		"dataCheckType": []string{cast.ToString(req.DataCheckType)},
	}
	if req.URL != "" {
		params["url"] = []string{req.URL}
	}
	if req.Data != "" {
		params["data"] = []string{req.Data}
	}
	if req.Title != "" {
		params["title"] = []string{req.Title}
	}
	if req.DataID != "" {
		params["dataId"] = []string{req.DataID}
	}
	if req.Callback != "" {
		params["callback"] = []string{req.Callback}
	}
	if req.CallbackURL != "" {
		params["callbackUrl"] = []string{req.CallbackURL}
	}
	if req.PublishTime > 0 {
		params["publishTime"] = []string{cast.ToString(req.PublishTime)}
	}
	if req.Nickname != "" {
		params["nickname"] = []string{req.Nickname}
	}
	if req.IP != "" {
		params["ip"] = []string{req.IP}
	}
	if req.Account != "" {
		params["account"] = []string{req.Account}
	}
	if req.DeviceID != "" {
		params["deviceId"] = []string{req.DeviceID}
	}
	if req.DeviceType > 0 {
		params["deviceType"] = []string{cast.ToString(req.DeviceType)}
	}
	if req.Extension != "" {
		params["extension"] = []string{req.Extension}
	}
	if req.SubProduct != "" {
		params["subProduct"] = []string{req.SubProduct}
	}
	if req.UniqueKey != "" {
		params["uniqueKey"] = []string{req.UniqueKey}
	}
	if len(req.RelatedKeys) > 0 {
		params["relatedKeys"] = req.RelatedKeys
	}
	// 业务拓展参数
	model.SplicingBusinessExpansion(req.BusinessExtension, params)

	bytes, err := c.Request(AudioURL, APIVersionV2, params)
	if err != nil {
		return
	}

	rsp = &model.AudioSyncDetectResp{}
	if err = json.Unmarshal(bytes, &rsp); err != nil {
		return
	}
	if rsp.Code != CallSuccessCode {
		err = fmt.Errorf("AudioDetect fail, code=%v", rsp.Code)
		return
	}

	return
}
