package handlers

import (
	"github.com/ACking-you/byte_douyin_project/models"
	"github.com/ACking-you/byte_douyin_project/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type FeedResponse struct {
	models.Response
	*service.FeedVideoList
}

func FeedVideoListHandler(c *gin.Context) {
	p := NewProxyFeedVideoList(c)
	token, ok := c.GetQuery("token")
	//无登录状态
	if !ok {
		rawTime := c.Query("latest_time")
		err := p.DoNoToken(rawTime)
		if err != nil {
			p.FeedVideoListError(err.Error())
		}
		return
	}

	//有登录状态 TODO 暂未实现
	err := p.DoHasToken(token)
	if err != nil {
		p.FeedVideoListError(err.Error())
	}
}

type ProxyFeedVideoList struct {
	*gin.Context
}

func NewProxyFeedVideoList(c *gin.Context) *ProxyFeedVideoList {
	return &ProxyFeedVideoList{Context: c}
}

// DoNoToken 未登录的视频流推送处理
func (p *ProxyFeedVideoList) DoNoToken(rawTimestamp string) error {
	var latestTime time.Time
	intTime, err := strconv.ParseInt(rawTimestamp, 10, 64)
	if err == nil {
		latestTime = time.Unix(0, intTime*1e6) //注意：前端传来的时间戳是以ms为单位的
	}
	videoList, err := service.QueryFeedVideoList(0, latestTime)
	if err != nil {
		return err
	}
	p.FeedVideoListOk(videoList)
	return nil
}

// DoHasToken TODO 如果有登录状态可以根据兴趣做一些推荐
func (p *ProxyFeedVideoList) DoHasToken(token string) error {

	return nil
}

func (p *ProxyFeedVideoList) FeedVideoListError(msg string) {
	p.JSON(http.StatusOK, FeedResponse{Response: models.Response{
		StatusCode: 1,
		StatusMsg:  msg,
	}})
}

func (p *ProxyFeedVideoList) FeedVideoListOk(videoList *service.FeedVideoList) {
	p.JSON(http.StatusOK, FeedResponse{
		Response: models.Response{
			StatusCode: 0,
		},
		FeedVideoList: videoList,
	},
	)
}
