package netease

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	"github.com/lcr2000/goutils"
	"github.com/lcr2000/netease-detect/model"
)

/*
  音频检测文档地址
  https://support.dun.163.com/documents/588434426518708224?docId=600817227615772672
*/

// LiveAudioDetectSubmit 直播音频提交检测
func (c *Client) LiveAudioDetectSubmit(req *model.LiveAudioDetectSubmitReq) (rsp *model.LiveAudioDetectSubmitResp, err error) {
	if req == nil || req.URL == "" || req.DataID == "" {
		err = errors.New("params is required")
		return
	}

	params := url.Values{
		"url":    []string{req.URL},
		"dataId": []string{req.DataID},
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
	if req.RoomNo != "" {
		params["roomNo"] = []string{req.RoomNo}
	}
	if req.AccountLevel != "" {
		params["accountLevel"] = []string{req.AccountLevel}
	}
	if req.AccountName != "" {
		params["accountName"] = []string{req.AccountName}
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
	if req.UniqueKey != "" {
		params["uniqueKey"] = []string{req.UniqueKey}
	}
	if req.CheckLanguageCode != "" {
		params["checkLanguageCode"] = []string{req.CheckLanguageCode}
	}

	bytes, err := c.Request(LiveAudioSubmitURL, APIVersionV4, params)
	if err != nil {
		return
	}

	rsp = &model.LiveAudioDetectSubmitResp{}
	if err = json.Unmarshal(bytes, &rsp); err != nil {
		return
	}
	if rsp.Code != CallSuccessCode {
		err = fmt.Errorf("LiveAudioDetectSubmit fail, code=%v", rsp.Code)
		return
	}

	return
}

// GetLiveAudioDetectResult 获取直播音频检测结果
func (c *Client) GetLiveAudioDetectResult() (rsp *model.LiveAudioDetectResultResp, err error) {
	params := url.Values{}
	body, err := c.Request(LiveAudioResultURL, APIVersionV4, params)
	if err != nil {
		return
	}
	rsp = &model.LiveAudioDetectResultResp{}
	if err = json.Unmarshal(body, &rsp); err != nil {
		return
	}
	if rsp.Code != CallSuccessCode {
		err = fmt.Errorf("GetLiveAudioDetectResult fail, code=%v", rsp.Code)
		return
	}
	return
}

// LiveAudioDetectStop 停止直播音频检测
func (c *Client) LiveAudioDetectStop(taskIds []string) (rsp *model.LiveAudioDetectStopResp, err error) {
	if len(taskIds) == 0 {
		err = errors.New("params is required")
		return
	}
	if len(taskIds) > 100 {
		err = errors.New("submit up to 100")
		return
	}

	// 直播信息更新数据(Json数组),提交时转换为string格式，数组最多100个
	feedback := make([]*model.LiveAudioDetectFeedback, 0, len(taskIds))
	for _, taskID := range taskIds {
		feedback = append(feedback, &model.LiveAudioDetectFeedback{
			TaskID: taskID,
			Status: StopDetectStatus,
		})
	}
	params := url.Values{
		"feedbacks": []string{goutils.JsonMarshalNoError(feedback)},
	}

	bytes, err := c.Request(LiveAudioStopURL, APIVersionV1, params)
	if err != nil {
		return
	}

	rsp = &model.LiveAudioDetectStopResp{}
	if err = json.Unmarshal(bytes, &rsp); err != nil {
		return
	}
	if rsp.Code != CallSuccessCode {
		err = fmt.Errorf("LiveAudioDetectStop fail, code=%v", rsp.Code)
		return
	}

	return
}

// LiveAudioDetectFeedback 直播音频检测反馈
func (c *Client) LiveAudioDetectFeedback(req *model.FeedbackReq) (rsp *model.LiveAudioDetectFeedbackResp, err error) {
	if req == nil || req.Level < 0 || req.TaskID == "" {
		err = errors.New("params is required")
		return
	}

	params := url.Values{
		"feedbacks": []string{goutils.JsonMarshalNoError(req)},
	}

	bytes, err := c.Request(LiveAudioFeedbackURL, APIVersionV1, params)
	if err != nil {
		return
	}

	rsp = &model.LiveAudioDetectFeedbackResp{}
	if err = json.Unmarshal(bytes, &rsp); err != nil {
		return
	}
	if rsp.Code != CallSuccessCode {
		err = fmt.Errorf("LiveAudioDetectFeedback fail, code=%v", rsp.Code)
		return
	}

	return
}
