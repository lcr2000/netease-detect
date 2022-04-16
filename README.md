# netease-detect
网易易盾 Golang 版本 SDK

 
 ## 单元测试
 鉴于安全原因, 不便暴露 secretId/secretKey/businessId 等信息.
 所以在测试命令后提供了如下参数.
 ```go
 go test -v -args -secretId xxx -secretKey xxx -businessId xxx
 ```

## 快速开始
 ```go
import (
	"fmt"
	netease_detect "github.com/lcr2000/netease-detect"
	"github.com/lcr2000/netease-detect/model"
)

func main() {
	client := netease_detect.NewClient("secretId", "secretKey", "businessId")
	rsp, err := client.ImageDetect(&model.ImageDetectReq{
		Images: []*model.ImageInfo{{
			Name:      "123456789",
			ImageType: 1,
			Data:      "https://img2.baidu.com/it/u=564570846,288904720&fm=253&fmt=auto&app=138&f=PNG?w=889&h=500",
		}},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rsp)
}
 ```