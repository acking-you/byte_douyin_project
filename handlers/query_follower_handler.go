package handlers

import (
	"errors"
	"github.com/ACking-you/byte_douyin_project/models"
	"github.com/ACking-you/byte_douyin_project/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FollowerListResponse struct {
	models.Response
	*service.FollowerList
}

func QueryFollowerHandler(c *gin.Context) {
	NewProxyQueryFollowerHandler(c).Do()
}

type ProxyQueryFollowerHandler struct {
	*gin.Context

	userId int64

	*service.FollowerList
}

func NewProxyQueryFollowerHandler(context *gin.Context) *ProxyQueryFollowerHandler {
	return &ProxyQueryFollowerHandler{Context: context}
}

func (p *ProxyQueryFollowerHandler) Do() {
	var err error
	if err = p.parseNum(); err != nil {
		p.SendError("token解析出错")
		return
	}
	if err = p.prepareData(); err != nil {
		if errors.Is(err, service.ErrUserNotExist) {
			p.SendError(err.Error())
		} else {
			p.SendError("准备数据出错")
		}
		return
	}
	p.SendOk("成功")
}

func (p *ProxyQueryFollowerHandler) parseNum() error {
	token := p.Query("token")
	userId, err := service.JWTAuth(token)
	if err != nil {
		return err
	}
	p.userId = userId
	return nil
}

func (p *ProxyQueryFollowerHandler) prepareData() error {
	list, err := service.QueryFollowerList(p.userId)
	if err != nil {
		return err
	}
	p.FollowerList = list
	return nil
}

func (p *ProxyQueryFollowerHandler) SendError(msg string) {
	p.JSON(http.StatusOK, FollowerListResponse{
		Response: models.Response{
			StatusCode: 1,
			StatusMsg:  msg,
		},
	})
}

func (p *ProxyQueryFollowerHandler) SendOk(msg string) {
	p.JSON(http.StatusOK, FollowerListResponse{
		Response: models.Response{
			StatusCode: 1,
			StatusMsg:  msg,
		},
		FollowerList: p.FollowerList,
	})
}
