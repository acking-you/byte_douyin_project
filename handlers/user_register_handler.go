package handlers

import (
	"github.com/ACking-you/byte_douyin_project/models"
	"github.com/ACking-you/byte_douyin_project/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserRegisterResponse struct {
	models.Response
	*service.UserLoginResponse
}

func UserRegisterHandler(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	registerResponse, err := service.PostUserLogin(username, password)

	if err != nil {
		c.JSON(http.StatusOK, UserRegisterResponse{
			Response: models.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	c.JSON(http.StatusOK, UserRegisterResponse{
		Response:          models.Response{StatusCode: 0},
		UserLoginResponse: registerResponse,
	})
}
