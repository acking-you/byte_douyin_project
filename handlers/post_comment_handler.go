package handlers

import (
	"fmt"
	"github.com/ACking-you/byte_douyin_project/models"
	"github.com/ACking-you/byte_douyin_project/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type PostCommentResponse struct {
	models.Response
	*service.CommentResponse
}

func PostCommentHandler(c *gin.Context) {
	NewProxyPostCommentHandler(c).Do()
}

type ProxyPostCommentHandler struct {
	*gin.Context

	videoId     int64
	userId      int64
	commentId   int64
	actionType  int64
	commentText string
}

func NewProxyPostCommentHandler(context *gin.Context) *ProxyPostCommentHandler {
	return &ProxyPostCommentHandler{Context: context}
}

func (p *ProxyPostCommentHandler) Do() {
	//解析参数
	if err := p.parseNum(); err != nil {
		p.SendError(err.Error())
		return
	}

	//正式调用
	commentRes, err := service.PostComment(p.userId, p.videoId, p.commentId, p.actionType, p.commentText)
	if err != nil {
		p.SendError(err.Error())
		return
	}

	//成功返回
	p.SendOk(commentRes)
}

func (p *ProxyPostCommentHandler) parseNum() error {
	rawVideoId := p.Query("video_id")
	videoId, err := strconv.ParseInt(rawVideoId, 10, 64)
	if err != nil {
		return err
	}
	p.videoId = videoId

	//根据actionType解析对应的可选参数
	rawActionType := p.Query("action_type")
	actionType, err := strconv.ParseInt(rawActionType, 10, 64)
	switch actionType {
	case service.CREATE:
		p.commentText = p.Query("comment_text")
	case service.DELETE:
		p.commentId, err = strconv.ParseInt(p.Query("comment_id"), 10, 64)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("未定义的行为%d", actionType)
	}
	p.actionType = actionType

	//解析userId，如果userId出错再根据token解析
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

func (p *ProxyPostCommentHandler) SendError(msg string) {
	p.JSON(http.StatusOK, PostCommentResponse{
		Response: models.Response{StatusCode: 1, StatusMsg: msg}, CommentResponse: &service.CommentResponse{}})
}

func (p *ProxyPostCommentHandler) SendOk(comment *service.CommentResponse) {
	p.JSON(http.StatusOK, PostCommentResponse{
		Response:        models.Response{StatusCode: 0},
		CommentResponse: comment,
	})
}
