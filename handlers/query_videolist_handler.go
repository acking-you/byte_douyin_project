package handlers

import (
	"github.com/ACking-you/byte_douyin_project/models"
	"github.com/ACking-you/byte_douyin_project/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type VideoListResponse struct {
	models.Response
	*service.VideoList
}

func QueryVideoListHandler(c *gin.Context) {
	p := NewProxyQueryVideoList(c)
	rawId := c.Query("user_id")
	err := p.DoQueryVideoListByUserId(rawId)
	if err == nil {
		return
	}

	//先查看user_id如果user_id出错，则再看看token
	token := c.Query("token")
	err = p.DoQueryVideoListByToken(token)
	if err != nil {
		p.QueryVideoListError(err.Error())
	}
}

// ProxyQueryVideoList 代理类
type ProxyQueryVideoList struct {
	c *gin.Context
}

func NewProxyQueryVideoList(c *gin.Context) *ProxyQueryVideoList {
	return &ProxyQueryVideoList{c: c}
}

// DoQueryVideoListByUserId 根据userId字段进行查询
func (p *ProxyQueryVideoList) DoQueryVideoListByUserId(rawId string) error {
	userId, err := strconv.ParseInt(rawId, 10, 64)
	if err != nil {
		return err
	}

	videoList, err := service.QueryVideoListByUserId(userId)
	if err != nil {
		return err
	}

	p.QueryVideoListOk(videoList)
	return nil
}

// DoQueryVideoListByToken 根据token查询
func (p *ProxyQueryVideoList) DoQueryVideoListByToken(token string) error {
	userId, err := service.JWTAuth(token)
	if err != nil {
		return err
	}

	videoList, err := service.QueryVideoListByUserId(userId)
	if err != nil {
		return err
	}

	p.QueryVideoListOk(videoList)
	return nil
}

func (p *ProxyQueryVideoList) QueryVideoListError(msg string) {
	p.c.JSON(http.StatusOK, VideoListResponse{Response: models.Response{
		StatusCode: 1,
		StatusMsg:  msg,
	}})
}

func (p *ProxyQueryVideoList) QueryVideoListOk(videoList *service.VideoList) {
	p.c.JSON(http.StatusOK, VideoListResponse{
		Response: models.Response{
			StatusCode: 0,
		},
		VideoList: videoList,
	})
}
