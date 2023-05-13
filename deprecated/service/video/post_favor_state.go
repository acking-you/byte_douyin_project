package video

import (
	"errors"
	"github.com/ACking-you/byte_douyin_project/cache"
	models2 "github.com/ACking-you/byte_douyin_project/models"
)

const (
	PLUS  = 1
	MINUS = 2
)

func PostFavorState(userId, videoId, actionType int64) error {
	return NewPostFavorStateFlow(userId, videoId, actionType).Do()
}

type PostFavorStateFlow struct {
	userId     int64
	videoId    int64
	actionType int64
}

func NewPostFavorStateFlow(userId, videoId, action int64) *PostFavorStateFlow {
	return &PostFavorStateFlow{
		userId:     userId,
		videoId:    videoId,
		actionType: action,
	}
}

func (p *PostFavorStateFlow) Do() error {
	var err error
	if err = p.checkNum(); err != nil {
		return err
	}

	switch p.actionType {
	case PLUS:
		err = p.PlusOperation()
	case MINUS:
		err = p.MinusOperation()
	default:
		return errors.New("未定义的操作")
	}
	return err
}

// PlusOperation 点赞操作
func (p *PostFavorStateFlow) PlusOperation() error {
	//视频点赞数目+1
	err := models2.NewVideoDAO().PlusOneFavorByUserIdAndVideoId(p.userId, p.videoId)
	if err != nil {
		return errors.New("不要重复点赞")
	}
	//对应的用户是否点赞的映射状态更新
	cache.NewProxyIndexMap().UpdateVideoFavorState(p.userId, p.videoId, true)
	return nil
}

// MinusOperation 取消点赞
func (p *PostFavorStateFlow) MinusOperation() error {
	//视频点赞数目-1
	err := models2.NewVideoDAO().MinusOneFavorByUserIdAndVideoId(p.userId, p.videoId)
	if err != nil {
		return errors.New("点赞数目已经为0")
	}
	//对应的用户是否点赞的映射状态更新
	cache.NewProxyIndexMap().UpdateVideoFavorState(p.userId, p.videoId, false)
	return nil
}

func (p *PostFavorStateFlow) checkNum() error {
	if !models2.NewUserInfoDAO().IsUserExistById(p.userId) {
		return errors.New("用户不存在")
	}
	if p.actionType != PLUS && p.actionType != MINUS {
		return errors.New("未定义的行为")
	}
	return nil
}
