package handlers

import (
	"github.com/ACking-you/byte_douyin_project/repository"
	"github.com/ACking-you/byte_douyin_project/service"
	"github.com/gin-gonic/gin"
	"time"
)

type FeedResponse struct {
	StatusCode repository.Response
	VideoList *service.VideoList
	NextTime int64 `json:"next_time,omitempty"`
}

func feed() *FeedResponse{
	return &FeedResponse{
		StatusCode: repository.Response{StatusCode: 0,StatusMsg: "ok"},
		VideoList: service.QueryVideoList(),
		NextTime: time.Now().Unix(),
	}
}

func FeedHandler(c *gin.Context) {

}