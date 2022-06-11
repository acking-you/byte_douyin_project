package comment

import (
	"errors"
	"github.com/ACking-you/byte_douyin_project/handlers/video"
	"github.com/ACking-you/byte_douyin_project/models"
	"github.com/ACking-you/byte_douyin_project/service/comment"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ListResponse struct {
	models.CommonResponse
	*comment.List
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
	commentList, err := comment.QueryCommentList(p.userId, p.videoId)
	if err != nil {
		p.SendError(err.Error())
		return
	}

	//成功返回
	p.SendOk(commentList)
}

func (p *ProxyCommentListHandler) parseNum() error {
	rawUserId, _ := p.Get("user_id")
	userId, ok := rawUserId.(int64)
	if !ok {
		return errors.New("userId解析出错")
	}
	p.userId = userId

	rawVideoId := p.Query("video_id")
	videoId, err := strconv.ParseInt(rawVideoId, 10, 64)
	if err != nil {
		return err
	}
	p.videoId = videoId

	return nil
}

func (p *ProxyCommentListHandler) SendError(msg string) {
	p.JSON(http.StatusOK, video.FavorVideoListResponse{
		CommonResponse: models.CommonResponse{StatusCode: 1, StatusMsg: msg}})
}

func (p *ProxyCommentListHandler) SendOk(commentList *comment.List) {
	p.JSON(http.StatusOK, ListResponse{CommonResponse: models.CommonResponse{StatusCode: 0},
		List: commentList,
	})
}
