package handlers

import (
	"errors"
	"github.com/ACking-you/byte_douyin_project/models"
	"github.com/ACking-you/byte_douyin_project/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func PostFollowActionHandler(c *gin.Context) {
	NewProxyPostFollowAction(c).Do()
}

type ProxyPostFollowAction struct {
	*gin.Context

	userId     int64
	followId   int64
	actionType int
}

func NewProxyPostFollowAction(context *gin.Context) *ProxyPostFollowAction {
	return &ProxyPostFollowAction{Context: context}
}

func (p *ProxyPostFollowAction) Do() {
	var err error
	if err = p.prepareNum(); err != nil {
		p.SendError("用户鉴权失败")
		return
	}
	if err = p.startAction(); err != nil {
		//当错误为model层发生的，那么就是重复键值的插入了
		if errors.Is(err, service.ErrIvdAct) || errors.Is(err, service.ErrIvdFolUsr) {
			p.SendError(err.Error())
		} else {
			p.SendError("请勿重复关注")
		}
		return
	}
	p.SendOk("操作成功")
}

func (p *ProxyPostFollowAction) prepareNum() error {
	token := p.Query("token")
	//必须包含鉴权token字段
	userId, err := service.JWTAuth(token)
	if err != nil {
		return err
	}
	p.userId = userId

	//解析需要关注的id
	followId := p.Query("to_user_id")
	parseInt, err := strconv.ParseInt(followId, 10, 64)
	if err != nil {
		return err
	}
	p.followId = parseInt

	//解析action_type
	actionType := p.Query("action_type")
	parseInt, err = strconv.ParseInt(actionType, 10, 32)
	if err != nil {
		return err
	}
	p.actionType = int(parseInt)
	return nil
}

func (p *ProxyPostFollowAction) startAction() error {
	err := service.PostFollowAction(p.userId, p.followId, p.actionType)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProxyPostFollowAction) SendError(msg string) {
	p.JSON(http.StatusOK, models.Response{StatusCode: 1, StatusMsg: msg})
}

func (p *ProxyPostFollowAction) SendOk(msg string) {
	p.JSON(http.StatusOK, models.Response{StatusCode: 1, StatusMsg: msg})
}
