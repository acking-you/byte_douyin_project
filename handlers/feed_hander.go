package handlers

import (
	"github.com/ACking-you/byte_douyin_project/models"
	"github.com/ACking-you/byte_douyin_project/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type FeedResponse struct {
	models.Response
	*service.VideoList
	NextTime int64 `json:"next_time,omitempty"`
}

//getFeed 获取返回内容
func getFeed() *FeedResponse {
	videos, err := service.QueryVideoListDemo()
	curTime := time.Now().Unix()
	if err != nil {
		return &FeedResponse{
			Response:  models.Response{StatusCode: -1, StatusMsg: err.Error()},
			VideoList: videos,
			NextTime:  curTime,
		}
	}
	return &FeedResponse{
		Response:  models.Response{StatusCode: 0, StatusMsg: "ok"},
		VideoList: videos,
		NextTime:  curTime,
	}
}

func FeedHandler(c *gin.Context) {
	res := getFeed()
	c.JSON(http.StatusOK, res)
}
