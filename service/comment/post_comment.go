package comment

import (
	"errors"
	"fmt"
	"github.com/ACking-you/byte_douyin_project/models"
	"github.com/ACking-you/byte_douyin_project/util"
)

const (
	CREATE = 1
	DELETE = 2
)

type Response struct {
	MyComment *models.Comment `json:"comment"`
}

func PostComment(userId int64, videoId int64, commentId int64, actionType int64, commentText string) (*Response, error) {
	return NewPostCommentFlow(userId, videoId, commentId, actionType, commentText).Do()
}

type PostCommentFlow struct {
	userId      int64
	videoId     int64
	commentId   int64
	actionType  int64
	commentText string

	comment *models.Comment

	*Response
}

func NewPostCommentFlow(userId int64, videoId int64, commentId int64, actionType int64, commentText string) *PostCommentFlow {
	return &PostCommentFlow{userId: userId, videoId: videoId, commentId: commentId, actionType: actionType, commentText: commentText}
}

func (p *PostCommentFlow) Do() (*Response, error) {
	var err error
	if err = p.checkNum(); err != nil {
		return nil, err
	}
	if err = p.prepareData(); err != nil {
		return nil, err
	}
	if err = p.packData(); err != nil {
		return nil, err
	}
	return p.Response, err
}

// CreateComment 增加评论
func (p *PostCommentFlow) CreateComment() (*models.Comment, error) {
	comment := models.Comment{UserInfoId: p.userId, VideoId: p.videoId, Content: p.commentText}
	err := models.NewCommentDAO().AddCommentAndUpdateCount(&comment)
	if err != nil {
		return nil, err
	}

	return &comment, nil
}

// DeleteComment 删除评论
func (p *PostCommentFlow) DeleteComment() (*models.Comment, error) {
	//获取comment
	var comment models.Comment
	err := models.NewCommentDAO().QueryCommentById(p.commentId, &comment)
	if err != nil {
		return nil, err
	}
	//删除comment
	err = models.NewCommentDAO().DeleteCommentAndUpdateCountById(p.commentId, p.videoId)
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

func (p *PostCommentFlow) checkNum() error {
	if !models.NewUserInfoDAO().IsUserExistById(p.userId) {
		return fmt.Errorf("用户%d不存在", p.userId)
	}
	if !models.NewVideoDAO().IsVideoExistById(p.videoId) {
		return fmt.Errorf("视频%d不存在", p.videoId)
	}
	if p.actionType != CREATE && p.actionType != DELETE {
		return errors.New("未定义的行为")
	}
	return nil
}

func (p *PostCommentFlow) prepareData() error {
	var err error
	switch p.actionType {
	case CREATE:
		p.comment, err = p.CreateComment()
	case DELETE:
		p.comment, err = p.DeleteComment()
	default:
		return errors.New("未定义的操作")
	}
	return err
}

func (p *PostCommentFlow) packData() error {
	//填充字段
	userInfo := models.UserInfo{}
	_ = models.NewUserInfoDAO().QueryUserInfoById(p.comment.UserInfoId, &userInfo)
	p.comment.User = userInfo
	_ = util.FillCommentFields(p.comment)

	p.Response = &Response{MyComment: p.comment}

	return nil
}
