package main

import (
	"github.com/ACking-you/byte_douyin_project/router"
)

func main() {
	r := router.InitDouyinRouter()

	err := r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		return
	}
}
