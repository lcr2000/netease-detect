package netease

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	"github.com/lcr2000/netease-detect/model"
)

/*
	点播视频检测文档地址
	https://support.dun.163.com/documents/2018041903?docId=150440843651764224
*/

// VideoDetectSubmit 提交点播视频异步检测
func (c *Client) VideoDetectSubmit(req *model.VideoDetectSubmitReq) (rsp *model.VideoCheckSubmitResp, err error) {
	if req == nil || req.URL == "" || req.DataID == "" {
		err = errors.New("params is required")
		return
	}

	params := url.Values{
		"url":     []string{req.URL},
		"dataId":  []string{req.DataID},
		"version": []string{"v3.2"},
	}
	if req.Title != "" {
		params["title"] = []string{req.Title}
	}
	if req.Callback != "" {
		params["callback"] = []string{req.Callback}
	}
	if req.CallbackURL != "" {
		params["callbackUrl"] = []string{req.CallbackURL}
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

	bytes, err := c.Request(VideoSubmitURL, "v3.2", params)
	if err != nil {
		return
	}

	rsp = &model.VideoCheckSubmitResp{}
	if err = json.Unmarshal(bytes, &rsp); err != nil {
		return
	}
	if rsp.Code != CallSuccessCode {
		err = fmt.Errorf("VideoDetectSubmit fail, code=%v", rsp.Code)
		return
	}

	return
}

// GetVideoDetectResult 获取点播视频检测结果
func (c *Client) GetVideoDetectResult() (rsp *model.VideoDetectResp, err error) {
	params := url.Values{}
	body, err := c.Request(VideoResultURL, "v3.1", params)
	if err != nil {
		return
	}
	rsp = &model.VideoDetectResp{}
	if err = json.Unmarshal(body, &rsp); err != nil {
		return
	}
	if rsp.Code != CallSuccessCode {
		err = fmt.Errorf("GetVideoDetectResult fail, code=%v", rsp.Code)
		return
	}
	return
}
