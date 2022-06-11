package user_info

import (
	"errors"
	"github.com/ACking-you/byte_douyin_project/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserResponse struct {
	models.CommonResponse
	User *models.UserInfo `json:"user"`
}

func UserInfoHandler(c *gin.Context) {
	p := NewProxyUserInfo(c)
	//得到上层中间件根据token解析的userId
	rawId, ok := c.Get("user_id")
	if !ok {
		p.UserInfoError("解析userId出错")
		return
	}
	err := p.DoQueryUserInfoByUserId(rawId)
	if err != nil {
		p.UserInfoError(err.Error())
	}
}

type ProxyUserInfo struct {
	c *gin.Context
}

func NewProxyUserInfo(c *gin.Context) *ProxyUserInfo {
	return &ProxyUserInfo{c: c}
}

func (p *ProxyUserInfo) DoQueryUserInfoByUserId(rawId interface{}) error {
	userId, ok := rawId.(int64)
	if !ok {
		return errors.New("解析userId失败")
	}
	//由于得到userinfo不需要组装model层的数据，所以直接调用model层的接口
	userinfoDAO := models.NewUserInfoDAO()

	var userInfo models.UserInfo
	err := userinfoDAO.QueryUserInfoById(userId, &userInfo)
	if err != nil {
		return err
	}
	p.UserInfoOk(&userInfo)
	return nil
}

func (p *ProxyUserInfo) UserInfoError(msg string) {
	p.c.JSON(http.StatusOK, UserResponse{
		CommonResponse: models.CommonResponse{StatusCode: 1, StatusMsg: msg},
	})
}

func (p *ProxyUserInfo) UserInfoOk(user *models.UserInfo) {
	p.c.JSON(http.StatusOK, UserResponse{
		CommonResponse: models.CommonResponse{StatusCode: 0},
		User:           user,
	})
}
