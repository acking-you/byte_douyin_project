package video

import (
	"errors"
	"github.com/ACking-you/byte_douyin_project/models"
	"github.com/ACking-you/byte_douyin_project/service/video"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FavorVideoListResponse struct {
	models.CommonResponse
	*video.FavorList
}

func QueryFavorVideoListHandler(c *gin.Context) {
	NewProxyFavorVideoListHandler(c).Do()
}

type ProxyFavorVideoListHandler struct {
	*gin.Context

	userId int64
}

func NewProxyFavorVideoListHandler(c *gin.Context) *ProxyFavorVideoListHandler {
	return &ProxyFavorVideoListHandler{Context: c}
}

func (p *ProxyFavorVideoListHandler) Do() {
	//解析参数
	if err := p.parseNum(); err != nil {
		p.SendError(err.Error())
		return
	}

	//正式调用
	favorVideoList, err := video.QueryFavorVideoList(p.userId)
	if err != nil {
		p.SendError(err.Error())
		return
	}

	//成功返回
	p.SendOk(favorVideoList)
}

func (p *ProxyFavorVideoListHandler) parseNum() error {
	rawUserId, _ := p.Get("user_id")
	userId, ok := rawUserId.(int64)
	if !ok {
		return errors.New("userId解析出错")
	}
	p.userId = userId
	return nil
}

func (p *ProxyFavorVideoListHandler) SendError(msg string) {
	p.JSON(http.StatusOK, FavorVideoListResponse{
		CommonResponse: models.CommonResponse{StatusCode: 1, StatusMsg: msg}})
}

func (p *ProxyFavorVideoListHandler) SendOk(favorList *video.FavorList) {
	p.JSON(http.StatusOK, FavorVideoListResponse{CommonResponse: models.CommonResponse{StatusCode: 0},
		FavorList: favorList,
	})
}
