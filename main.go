package main

import (
	"fmt"
	"github.com/ACking-you/byte_douyin_project/config"
	"github.com/ACking-you/byte_douyin_project/router"
)

func main() {
	r := router.InitDouyinRouter()
	err := r.Run(fmt.Sprintf(":%d", config.Info.Port)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		return
	}
}
