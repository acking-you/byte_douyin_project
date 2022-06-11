package video

import (
	"errors"
	"github.com/ACking-you/byte_douyin_project/models"
	"github.com/ACking-you/byte_douyin_project/service/video"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func PostFavorHandler(c *gin.Context) {
	NewProxyPostFavorHandler(c).Do()
}

type ProxyPostFavorHandler struct {
	*gin.Context

	userId     int64
	videoId    int64
	actionType int64
}

func NewProxyPostFavorHandler(c *gin.Context) *ProxyPostFavorHandler {
	return &ProxyPostFavorHandler{Context: c}
}

func (p *ProxyPostFavorHandler) Do() {
	//解析参数
	if err := p.parseNum(); err != nil {
		p.SendError(err.Error())
		return
	}

	//正式调用
	err := video.PostFavorState(p.userId, p.videoId, p.actionType)
	if err != nil {
		p.SendError(err.Error())
		return
	}

	//成功返回
	p.SendOk()
}

func (p *ProxyPostFavorHandler) parseNum() error {
	//解析userId
	rawUserId, _ := p.Get("user_id")
	userId, ok := rawUserId.(int64)
	if !ok {
		return errors.New("userId解析出错")
	}

	rawVideoId := p.Query("video_id")
	videoId, err := strconv.ParseInt(rawVideoId, 10, 64)
	if err != nil {
		return err
	}
	rawActionType := p.Query("action_type")
	actionType, err := strconv.ParseInt(rawActionType, 10, 64)
	if err != nil {
		return err
	}
	p.videoId = videoId
	p.actionType = actionType
	p.userId = userId
	return nil
}

func (p *ProxyPostFavorHandler) SendError(msg string) {
	p.JSON(http.StatusOK, models.CommonResponse{StatusCode: 1, StatusMsg: msg})
}

func (p *ProxyPostFavorHandler) SendOk() {
	p.JSON(http.StatusOK, models.CommonResponse{StatusCode: 0})
}
