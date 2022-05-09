package netease

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"math/rand"
	"net/url"
	"sort"
	"strconv"
	"time"

	"github.com/lcr2000/goutils"
)

// Client 实例
type Client struct {
	secretID   string // 必填 产品秘钥 id,由易盾内容安全服务分配,产品标识
	secretKey  string // 必填 key
	businessID string // 非必填 业务id,由易盾内容安全服务分配,业务标识
}

// NewClient 初始化网易易盾客户端实例,一般在程序启动的时候调用进行初始化
// secretID、secretKey是必填的,传入空值将panic; businessID为可选值
func NewClient(secretID, secretKey string, businessID ...string) *Client {
	if secretID == "" || secretKey == "" {
		panic("secretID and secretKey is required")
	}
	var bid string
	if len(businessID) > 0 {
		bid = businessID[0]
	}
	return &Client{
		secretID:   secretID,
		secretKey:  secretKey,
		businessID: bid,
	}
}

// Request 通用的请求
func (c *Client) Request(apiURL, version string, params url.Values) (resp []byte, err error) {
	if c.businessID != "" {
		params["businessId"] = []string{c.businessID}
	}
	params["secretId"] = []string{c.secretID}
	params["version"] = []string{version}
	params["timestamp"] = []string{strconv.FormatInt(time.Now().UnixNano()/1000000, 10)}
	params["nonce"] = []string{strconv.FormatInt(rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(10000000000), 10)}
	params["signature"] = []string{GenSignature(params, c.secretKey)}
	resp, err = goutils.HttpPost(apiURL, goutils.HttpContentTypeForm, params.Encode())
	log.Printf("Request apiURL= %s, req= %v, resp= %s", apiURL, params, string(resp))
	if err != nil {
		return
	}
	return
}

// GenSignature 生成签名信息
func GenSignature(params url.Values, secretKey string) string {
	var paramStr string
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		paramStr += key + params[key][0]
	}
	paramStr += secretKey
	md5Reader := md5.New()
	md5Reader.Write([]byte(paramStr))
	return hex.EncodeToString(md5Reader.Sum(nil))
}
