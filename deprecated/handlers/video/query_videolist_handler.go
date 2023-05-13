package video

import (
	"errors"
	"github.com/ACking-you/byte_douyin_project/models"
	"github.com/ACking-you/byte_douyin_project/service/video"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ListResponse struct {
	models.CommonResponse
	*video.List
}

func QueryVideoListHandler(c *gin.Context) {
	p := NewProxyQueryVideoList(c)
	rawId, _ := c.Get("user_id")
	err := p.DoQueryVideoListByUserId(rawId)
	if err != nil {
		p.QueryVideoListError(err.Error())
	}
}

// ProxyQueryVideoList 代理类
type ProxyQueryVideoList struct {
	c *gin.Context
}

func NewProxyQueryVideoList(c *gin.Context) *ProxyQueryVideoList {
	return &ProxyQueryVideoList{c: c}
}

// DoQueryVideoListByUserId 根据userId字段进行查询
func (p *ProxyQueryVideoList) DoQueryVideoListByUserId(rawId interface{}) error {
	userId, ok := rawId.(int64)
	if !ok {
		return errors.New("userId解析出错")
	}

	videoList, err := video.QueryVideoListByUserId(userId)
	if err != nil {
		return err
	}

	p.QueryVideoListOk(videoList)
	return nil
}

func (p *ProxyQueryVideoList) QueryVideoListError(msg string) {
	p.c.JSON(http.StatusOK, ListResponse{CommonResponse: models.CommonResponse{
		StatusCode: 1,
		StatusMsg:  msg,
	}})
}

func (p *ProxyQueryVideoList) QueryVideoListOk(videoList *video.List) {
	p.c.JSON(http.StatusOK, ListResponse{
		CommonResponse: models.CommonResponse{
			StatusCode: 0,
		},
		List: videoList,
	})
}
