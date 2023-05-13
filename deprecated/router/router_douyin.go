package router

import (
	"github.com/ACking-you/byte_douyin_project/config"
	comment2 "github.com/ACking-you/byte_douyin_project/handlers/comment"
	user_info2 "github.com/ACking-you/byte_douyin_project/handlers/user_info"
	user_login2 "github.com/ACking-you/byte_douyin_project/handlers/user_login"
	video2 "github.com/ACking-you/byte_douyin_project/handlers/video"
	middleware2 "github.com/ACking-you/byte_douyin_project/middleware"
	"github.com/ACking-you/byte_douyin_project/models"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	models.InitDB()
	r := gin.Default()

	r.Static("static", config.Global.StaticSourcePath)

	baseGroup := r.Group("/douyin")
	//根据灵活性考虑是否加入JWT中间件来进行鉴权，还是在之后再做鉴权
	// basic apis
	baseGroup.GET("/feed/", video2.FeedVideoListHandler)
	baseGroup.GET("/user/", middleware2.JWTMiddleWare(), user_info2.UserInfoHandler)
	baseGroup.POST("/user/login/", middleware2.SHAMiddleWare(), user_login2.UserLoginHandler)
	baseGroup.POST("/user/register/", middleware2.SHAMiddleWare(), user_login2.UserRegisterHandler)
	baseGroup.POST("/publish/action/", middleware2.JWTMiddleWare(), video2.PublishVideoHandler)
	baseGroup.GET("/publish/list/", middleware2.NoAuthToGetUserId(), video2.QueryVideoListHandler)

	//extend 1
	baseGroup.POST("/favorite/action/", middleware2.JWTMiddleWare(), video2.PostFavorHandler)
	baseGroup.GET("/favorite/list/", middleware2.NoAuthToGetUserId(), video2.QueryFavorVideoListHandler)
	baseGroup.POST("/comment/action/", middleware2.JWTMiddleWare(), comment2.PostCommentHandler)
	baseGroup.GET("/comment/list/", middleware2.JWTMiddleWare(), comment2.QueryCommentListHandler)

	//extend 2
	baseGroup.POST("/relation/action/", middleware2.JWTMiddleWare(), user_info2.PostFollowActionHandler)
	baseGroup.GET("/relation/follow/list/", middleware2.NoAuthToGetUserId(), user_info2.QueryFollowListHandler)
	baseGroup.GET("/relation/follower/list/", middleware2.NoAuthToGetUserId(), user_info2.QueryFollowerHandler)
	return r
}
