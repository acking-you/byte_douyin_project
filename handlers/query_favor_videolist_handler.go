package handlers

import (
	"github.com/ACking-you/byte_douyin_project/models"
	"github.com/ACking-you/byte_douyin_project/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type FavorVideoListResponse struct {
	models.Response
	*service.FavorVideoList
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
	favorVideoList, err := service.QueryFavorVideoList(p.userId)
	if err != nil {
		p.SendError(err.Error())
		return
	}

	//成功返回
	p.SendOk(favorVideoList)
}

func (p *ProxyFavorVideoListHandler) parseNum() error {
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

func (p *ProxyFavorVideoListHandler) SendError(msg string) {
	p.JSON(http.StatusOK, FavorVideoListResponse{
		Response: models.Response{StatusCode: 1, StatusMsg: msg}})
}

func (p *ProxyFavorVideoListHandler) SendOk(favorList *service.FavorVideoList) {
	p.JSON(http.StatusOK, FavorVideoListResponse{Response: models.Response{StatusCode: 0},
		FavorVideoList: favorList,
	})
}
