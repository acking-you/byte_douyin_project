package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	DB2 "simpleTikTok/DB"
	"simpleTikTok/common"
	"simpleTikTok/model"
)

// UserLoginResponse 返回的注册和登录的Response
type UserLoginResponse struct {
	model.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

// UserResponse 拉取user信息的Response
type UserResponse struct {
	model.Response
	User model.User `json:"user"`
}

// Register 注册逻辑
func Register(ctx *gin.Context) {
	DB := DB2.GetDB()
	username := ctx.Query("username")
	password := ctx.Query("password")

	if isUsernameExist(DB, username) {
		ctx.JSON(http.StatusUnprocessableEntity, UserLoginResponse{
			Response: model.Response{
				StatusCode: 422,
				StatusMsg:  "用户已存在",
			},
		})
		return
	}

	// 创建用户
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(500, UserLoginResponse{
			Response: model.Response{StatusCode: 500, StatusMsg: "系统错误"},
		})
		return
	}

	newUser := model.UserLogin{
		Username: username,
		Password: string(hashedPassword),
	}

	DB.Create(&newUser)

	// 发放token
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, UserLoginResponse{
			Response: model.Response{StatusCode: 500, StatusMsg: "系统错误"},
		})
		log.Printf("token generate error: %v\n", err)
		return
	}

	// User表中的数据
	DB.Create(&model.User{
		Id:            newUser.ID,
		Name:          newUser.Username,
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      false,
	})

	ctx.JSON(http.StatusOK, UserLoginResponse{
		Response: model.Response{StatusCode: 0},
		UserId:   newUser.ID,
		Token:    token,
	})
}

func Login(ctx *gin.Context) {
	DB := DB2.GetDB()
	username := ctx.Query("username")
	password := ctx.Query("password")

	var user model.UserLogin
	DB.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, UserLoginResponse{
			Response: model.Response{
				StatusCode: 1,
				StatusMsg:  "用户不存在",
			},
		})
	}
	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status_code": 400, "status_msg": "密码错误!"})
		return
	}

	// 发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, UserLoginResponse{
			Response: model.Response{StatusCode: 500, StatusMsg: "系统错误"},
		})
		log.Printf("token generate error: %v\n", err)
		return
	}

	ctx.JSON(http.StatusOK, UserLoginResponse{
		Response: model.Response{StatusCode: 0},
		UserId:   user.ID,
		Token:    token,
	})

}

// UserInfo 这里会传ID和token，就不解析token了，后续可以解析token
func UserInfo(ctx *gin.Context) {
	DB := DB2.GetDB()

	userid := ctx.Query("user_id")
	var user model.User

	DB.Where("Id = ?", userid).First(&user)

	ctx.JSON(http.StatusOK, UserResponse{
		Response: model.Response{StatusCode: 0},
		User:     user,
	})
}

// 判断username是否存在
func isUsernameExist(db *gorm.DB, username string) bool {
	var user model.UserLogin
	db.Where("username = ?", username).First(&user)
	return user.ID != 0
}
