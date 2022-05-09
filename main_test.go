package netease

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

var (
	secretID   string
	secretKey  string
	businessID string
)

// 鉴于安全原因, 不便暴露 secretID/secretKey/businessID 等信息
// 所以在测试命令后提供了如下参数.
func TestMain(m *testing.M) {
	flag.StringVar(&secretID, "secretID", "", "secretID")
	flag.StringVar(&secretKey, "secretKey", "", "secretKey")
	flag.StringVar(&businessID, "businessID", "", "businessID")
	if !flag.Parsed() {
		flag.Parse()
	}
	fmt.Printf("-args -secretID %s -secretKey %s -businessID %s\n", secretID, secretKey, businessID)
	m.Run()
	os.Exit(1)
}
