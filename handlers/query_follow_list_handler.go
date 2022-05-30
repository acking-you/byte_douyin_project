package handlers

import (
	"errors"
	"github.com/ACking-you/byte_douyin_project/models"
	"github.com/ACking-you/byte_douyin_project/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FollowListResponse struct {
	models.Response
	*service.FollowList
}

func QueryFollowListHandler(c *gin.Context) {
	NewProxyQueryFollowList(c).Do()
}

type ProxyQueryFollowList struct {
	*gin.Context

	userId int64

	*service.FollowList
}

func NewProxyQueryFollowList(context *gin.Context) *ProxyQueryFollowList {
	return &ProxyQueryFollowList{Context: context}
}

func (p *ProxyQueryFollowList) Do() {
	var err error
	if err = p.parseNum(); err != nil {
		p.SendError("token解析错误")
		return
	}
	if err = p.prepareData(); err != nil {
		if errors.Is(err, service.ErrUserNotExist) {
			p.SendError(err.Error())
		} else {
			p.SendError("数据库访问出错")
		}
		return
	}
	p.SendOk("请求成功")
}

func (p *ProxyQueryFollowList) parseNum() error {
	token := p.Query("token")
	userId, err := service.JWTAuth(token)
	if err != nil {
		return err
	}
	p.userId = userId
	return nil
}

func (p *ProxyQueryFollowList) prepareData() error {
	list, err := service.QueryFollowList(p.userId)
	if err != nil {
		return err
	}
	p.FollowList = list
	return nil
}
func (p *ProxyQueryFollowList) SendError(msg string) {
	p.JSON(http.StatusOK, FollowListResponse{
		Response: models.Response{StatusCode: 1, StatusMsg: msg},
	})
}

func (p *ProxyQueryFollowList) SendOk(msg string) {
	p.JSON(http.StatusOK, FollowListResponse{
		Response:   models.Response{StatusCode: 0, StatusMsg: msg},
		FollowList: p.FollowList,
	})
}
