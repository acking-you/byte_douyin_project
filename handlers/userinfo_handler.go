package handlers

import (
	"github.com/ACking-you/byte_douyin_project/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserResponse struct {
	models.Response
	User *models.UserInfo `json:"user"`
}

func UserInfoHandler(c *gin.Context) {
	rawId := c.Query("user_id")
	userId, err := strconv.ParseInt(rawId, 10, 64)
	if err != nil {
		UserInfoError(c, "id解析错误")
		return
	}
	//TODO 由于传来的直接就有用户的id信息，没必要再鉴权得到id，token可以方便以后进行真正的鉴权系统（用户权限，目前尚未开发
	_ = c.Query("token")

	//由于得到userinfo不需要组装model层的数据，所以直接调用model层的接口
	userinfoDAO := models.NewUserInfoDAO()

	var userInfo models.UserInfo
	err = userinfoDAO.QueryUserInfoById(userId, &userInfo)
	if err != nil {
		UserInfoError(c, err.Error())
		return
	}

	UserInfoOk(c, &userInfo)
}

func UserInfoError(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, UserResponse{
		Response: models.Response{StatusCode: 1, StatusMsg: msg},
	})
}

func UserInfoOk(c *gin.Context, user *models.UserInfo) {
	c.JSON(http.StatusOK, UserResponse{
		Response: models.Response{StatusCode: 0},
		User:     user,
	})
}
