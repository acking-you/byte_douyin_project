package user_login

import (
	"github.com/ACking-you/byte_douyin_project/models"
	"github.com/ACking-you/byte_douyin_project/service/user_login"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserLoginResponse struct {
	models.CommonResponse
	*user_login.LoginResponse
}

func UserLoginHandler(c *gin.Context) {
	username := c.Query("username")
	raw, _ := c.Get("password")
	password, ok := raw.(string)
	if !ok {
		c.JSON(http.StatusOK, UserLoginResponse{
			CommonResponse: models.CommonResponse{
				StatusCode: 1,
				StatusMsg:  "密码解析错误",
			},
		})
	}
	userLoginResponse, err := user_login.QueryUserLogin(username, password)

	//用户不存在返回对应的错误
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			CommonResponse: models.CommonResponse{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}

	//用户存在，返回相应的id和token
	c.JSON(http.StatusOK, UserLoginResponse{
		CommonResponse: models.CommonResponse{StatusCode: 0},
		LoginResponse:  userLoginResponse,
	})
}
