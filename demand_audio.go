package netease_detect

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lcr2000/netease-detect/model"
	"net/url"
)

/*
	点播语音文档地址
	https://support.dun.163.com/documents/2018082201?docId=191344157972942848
*/

// DemandAudioDetectSubmit 点播音频提交检测
func (c *Client) DemandAudioDetectSubmit(req *model.DemandAudioDetectSubmitReq) (rsp *model.AudioCheckSubmitResponse, err error) {
	if req == nil || req.Url == "" {
		err = errors.New("params is required")
		return
	}

	params := url.Values{
		"url":     []string{req.Url},
		"version": []string{"v3.5"},
	}
	if req.Title != "" {
		params["title"] = []string{req.Title}
	}
	if req.Ip != "" {
		params["ip"] = []string{req.Ip}
	}
	if req.Account != "" {
		params["account"] = []string{req.Account}
	}
	if req.DeviceId != "" {
		params["deviceId"] = []string{req.DeviceId}
	}
	if req.DeviceType != 0 {
		params["deviceType"] = []string{fmt.Sprintf("%d", req.DeviceType)}
	}
	if req.Callback != "" {
		params["callback"] = []string{req.Callback}
	}
	if req.CallbackUrl != "" {
		params["callbackUrl"] = []string{req.CallbackUrl}
	}
	if req.SubProduct != "" {
		params["subProduct"] = []string{req.SubProduct}
	}

	bytes, err := c.Request(DemandAudioDetectSubmitUrl, "v3.5", params)
	if err != nil {
		return
	}

	rsp = &model.AudioCheckSubmitResponse{}
	if err = json.Unmarshal(bytes, &rsp); err != nil {
		return
	}
	if rsp.Code != CallSuccessCode {
		err = fmt.Errorf("DemandAudioDetectSubmit fail, code=%v", rsp.Code)
		return
	}

	return
}

// GetDemandAudioDetectResult 获取点播音频检测结果
func (c *Client) GetDemandAudioDetectResult() (rsp *model.AudioDetectResp, err error) {
	params := url.Values{}
	body, err := c.Request(DemandAudioDetectResultUrl, "v3.5", params)
	if err != nil {
		return
	}
	rsp = &model.AudioDetectResp{}
	if err = json.Unmarshal(body, &rsp); err != nil {
		return
	}
	if rsp.Code != CallSuccessCode {
		err = fmt.Errorf("GetDemandAudioDetectResult fail, code=%v", rsp.Code)
		return
	}
	return
}
