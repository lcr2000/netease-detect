package netease

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	"github.com/lcr2000/goutils"
	"github.com/lcr2000/netease-detect/model"
)

// LiveDetectSubmit 直播音视频提交检测
func (c *Client) LiveDetectSubmit(req *model.LiveDetectSubmitReq) (rsp *model.LiveDetectSubmitResp, err error) {
	if req == nil || req.Url == "" || req.DataId == "" {
		err = errors.New("params is required")
		return
	}

	params := url.Values{
		"url":    []string{req.Url},
		"dataId": []string{req.DataId},
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
	if req.RoomNo != "" {
		params["roomNo"] = []string{req.RoomNo}
	}
	if req.Age != "" {
		params["age"] = []string{req.Age}
	}
	if req.AccountLevel != "" {
		params["accountLevel"] = []string{req.AccountLevel}
	}
	if req.AccountName != "" {
		params["accountName"] = []string{req.AccountName}
	}
	if req.DeviceId != "" {
		params["deviceId"] = []string{req.DeviceId}
	}
	if req.DeviceType != 0 {
		params["deviceType"] = []string{fmt.Sprintf("%d", req.DeviceType)}
	}
	if req.Livelink != "" {
		params["livelink"] = []string{req.Livelink}
	}
	if req.ScreenMode != 0 {
		params["ScreenMode"] = []string{fmt.Sprintf("%d", req.ScreenMode)}
	}
	if req.DetectType != 0 {
		params["detectType"] = []string{fmt.Sprintf("%d", req.DetectType)}
	}
	if req.LabourUnion != "" {
		params["labourUnion"] = []string{req.LabourUnion}
	}
	if req.OperationManager != "" {
		params["operationManager"] = []string{req.OperationManager}
	}
	if req.ScFrequency != 0 {
		params["scFrequency"] = []string{fmt.Sprintf("%d", req.ScFrequency)}
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
	if req.WallHidden != 0 {
		params["wallHidden"] = []string{fmt.Sprintf("%d", req.WallHidden)}
	}
	if req.CheckLanguageCode != "" {
		params["checkLanguageCode"] = []string{req.CheckLanguageCode}
	}

	body, err := c.Request(LiveSubmitURL, APIVersionV3, params)
	if err != nil {
		return
	}

	rsp = &model.LiveDetectSubmitResp{}
	if err = json.Unmarshal(body, &rsp); err != nil {
		return
	}
	if rsp.Code != CallSuccessCode {
		err = fmt.Errorf("LiveDetectSubmit fail, code=%v, errMsg=%s", rsp.Code, rsp.Msg)
		return
	}

	return
}

// GetLiveDetectResult 获取直播音视频检测结果
func (c *Client) GetLiveDetectResult() (rsp *model.LiveDetectResultResp, err error) {
	params := url.Values{}
	body, err := c.Request(LiveResultURL, APIVersionV3, params)
	if err != nil {
		return
	}
	rsp = &model.LiveDetectResultResp{}
	if err = json.Unmarshal(body, &rsp); err != nil {
		return
	}
	if rsp.Code != CallSuccessCode {
		err = fmt.Errorf("GetLiveDetectResult fail, code=%v", rsp.Code)
		return
	}
	return
}

// LiveDetectStop 停止直播音视频检测
func (c *Client) LiveDetectStop(taskIds []string) (rsp *model.LiveDetectStopResp, err error) {
	if len(taskIds) == 0 {
		err = errors.New("params is required")
		return
	}

	realTimeInfoList := make([]*model.LiveDetectTaskInfo, 0)
	for _, taskID := range taskIds {
		realTimeInfoList = append(realTimeInfoList, &model.LiveDetectTaskInfo{
			TaskId: taskID,
			Status: StopDetectStatus,
		})
	}
	params := url.Values{
		"realTimeInfoList": []string{goutils.JsonMarshalNoError(realTimeInfoList)},
	}

	bytes, err := c.Request(LiveStopURL, APIVersionV1, params)
	if err != nil {
		return
	}

	rsp = &model.LiveDetectStopResp{}
	if err = json.Unmarshal(bytes, &rsp); err != nil {
		return
	}
	if rsp.Code != CallSuccessCode {
		err = fmt.Errorf("LiveDetectStop fail, code=%v, errMsg=%s", rsp.Code, rsp.Msg)
		return
	}

	return
}
