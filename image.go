package netease_detect

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lcr2000/goutils"
	"github.com/lcr2000/netease-detect/model"
	"net/url"
)

/*
	图片在线检测文档地址
	https://support.dun.163.com/documents/2018041902?docId=424387300808773632
*/

// ImageDetect 图片检测
func (c *Client) ImageDetect(req *model.ImageDetectReq) (rsp *model.ImageDetectResp, err error) {
	if req == nil || len(req.Images) == 0 {
		err = errors.New("params is required")
		return
	}
	// 单次请求最多支持32张图片
	if len(req.Images) > 32 {
		err = errors.New("support up to 32 images")
		return
	}

	params := url.Values{
		"images":  []string{goutils.JsonMarshalNoError(req.Images)},
		"version": []string{"v4.1"},
	}
	// 业务参数
	if req.Ip != "" {
		params["ip"] = []string{req.Ip}
	}
	if len(req.CheckLabels) > 0 {
		params["checkLabels"] = req.CheckLabels
	}
	if req.SubProduct != "" {
		params["subProduct"] = []string{req.SubProduct}
	}
	if req.Extension != "" {
		params["extension"] = []string{req.Extension}
	}
	// 业务拓展参数
	model.SplicingBusinessExpansion(req.BusinessExtension, params)

	bytes, err := c.Request(ImageDetectUrl, "v4.1", params)
	if err != nil {
		return
	}

	rsp = &model.ImageDetectResp{}
	if err = json.Unmarshal(bytes, &rsp); err != nil {
		return
	}
	if rsp.Code != CallSuccessCode {
		err = fmt.Errorf("ImageDetect fail, code=%v", rsp.Code)
		return
	}

	return
}
