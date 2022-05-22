package handlers

import (
	"github.com/ACking-you/byte_douyin_project/models"
	"github.com/ACking-you/byte_douyin_project/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CommentListResponse struct {
	models.Response
	*service.CommentList
}

func QueryCommentListHandler(c *gin.Context) {
	NewProxyCommentListHandler(c).Do()
}

type ProxyCommentListHandler struct {
	*gin.Context

	videoId int64
	userId  int64
}

func NewProxyCommentListHandler(context *gin.Context) *ProxyCommentListHandler {
	return &ProxyCommentListHandler{Context: context}
}

func (p *ProxyCommentListHandler) Do() {
	//解析参数
	if err := p.parseNum(); err != nil {
		p.SendError(err.Error())
		return
	}

	//正式调用
	commentList, err := service.QueryCommentList(p.userId, p.videoId)
	if err != nil {
		p.SendError(err.Error())
		return
	}

	//成功返回
	p.SendOk(commentList)
}

func (p *ProxyCommentListHandler) parseNum() error {
	rawVideoId := p.Query("video_id")
	videoId, err := strconv.ParseInt(rawVideoId, 10, 64)
	if err != nil {
		return err
	}
	p.videoId = videoId

	rawUserId := p.Query("user_id")
	userId, err := strconv.ParseInt(rawUserId, 10, 64)
	if err == nil {
		p.userId = userId
		return nil
	}
	//如果userId解析有问题，才换token
	token := p.Query("token")
	userId, err = service.JWTAuth(token)
	if err != nil {
		return err
	}
	p.userId = userId
	return nil
}

func (p *ProxyCommentListHandler) SendError(msg string) {
	p.JSON(http.StatusOK, FavorVideoListResponse{
		Response: models.Response{StatusCode: 1, StatusMsg: msg}})
}

func (p *ProxyCommentListHandler) SendOk(commentList *service.CommentList) {
	p.JSON(http.StatusOK, CommentListResponse{Response: models.Response{StatusCode: 0},
		CommentList: commentList,
	})
}
