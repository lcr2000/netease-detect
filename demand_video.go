package netease_detect

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lcr2000/netease-detect/model"
	"net/url"
)

/*
	视频点播文档地址
	https://support.dun.163.com/documents/2018041903?docId=150440843651764224
*/

// DemandVideoDetectSubmit 点播视频提交检测
func (c *Client) DemandVideoDetectSubmit(req *model.DemandVideoDetectSubmitReq) (rsp *model.VideoCheckSubmitResp, err error) {
	if req == nil || req.Url == "" || req.DataId == "" {
		err = errors.New("params is required")
		return
	}

	params := url.Values{
		"url":     []string{req.Url},
		"dataId":  []string{req.DataId},
		"version": []string{"v3.2"},
	}
	if req.Title != "" {
		params["title"] = []string{req.Title}
	}
	if req.Callback != "" {
		params["callback"] = []string{req.Callback}
	}
	if req.CallbackUrl != "" {
		params["callbackUrl"] = []string{req.CallbackUrl}
	}
	if req.UniqueKey != "" {
		params["uniqueKey"] = []string{req.UniqueKey}
	}
	if req.SubProduct != "" {
		params["subProduct"] = []string{req.SubProduct}
	}
	if req.Account != "" {
		params["account"] = []string{req.Account}
	}
	if req.AdvancedFrequency != "" {
		params["advancedFrequency"] = []string{req.AdvancedFrequency}
	}

	bytes, err := c.Request(DemandVideoDetectSubmitUrl, "v3.2", params)
	if err != nil {
		return
	}

	rsp = &model.VideoCheckSubmitResp{}
	if err = json.Unmarshal(bytes, &rsp); err != nil {
		return
	}
	if rsp.Code != CallSuccessCode {
		err = fmt.Errorf("DemandVideoDetectSubmit fail, code=%v", rsp.Code)
		return
	}

	return
}

// GetDemandVideoDetectResult 获取点播视频检测结果
func (c *Client) GetDemandVideoDetectResult() (rsp *model.VideoDetectResp, err error) {
	params := url.Values{}
	body, err := c.Request(DemandVideoDetectResultUrl, "v3.1", params)
	if err != nil {
		return
	}
	rsp = &model.VideoDetectResp{}
	if err = json.Unmarshal(body, &rsp); err != nil {
		return
	}
	if rsp.Code != CallSuccessCode {
		err = fmt.Errorf("GetDemandVideoDetectResult fail, code=%v", rsp.Code)
		return
	}
	return
}
