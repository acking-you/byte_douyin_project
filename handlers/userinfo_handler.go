package handlers

import (
	"github.com/ACking-you/byte_douyin_project/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserResponse struct {
	repository.Response
	User repository.User `json:"user"`
}

func UserInfoHandler(c *gin.Context) {
	token := c.Query("token")

	if user, exist := repository.DemoUsersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserResponse{
			Response: repository.Response{StatusCode: 0},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: repository.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}