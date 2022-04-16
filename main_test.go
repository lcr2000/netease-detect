package netease_detect

import (
	"flag"
	"fmt"
	"testing"
)

var (
	secretId   string
	secretKey  string
	businessId string
)

// 鉴于安全原因, 不便暴露 secretId/secretKey/businessId 等信息
// 所以在测试命令后提供了如下参数.
func TestMain(m *testing.M) {
	flag.StringVar(&secretId, "secretId", "", "secretId")
	flag.StringVar(&secretKey, "secretKey", "", "secretKey")
	flag.StringVar(&businessId, "businessId", "", "businessId")
	if !flag.Parsed() {
		flag.Parse()
	}
	fmt.Printf("-args -secretId %s -secretKey %s -businessId %s\n", secretId, secretKey, businessId)
	m.Run()
}
