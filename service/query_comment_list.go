package service

import (
	"fmt"
	"github.com/ACking-you/byte_douyin_project/models"
	"github.com/ACking-you/byte_douyin_project/util"
)

type CommentList struct {
	Comments []*models.Comment `json:"comment_list"`
}

func QueryCommentList(userId, videoId int64) (*CommentList, error) {
	return NewQueryCommentListFlow(userId, videoId).Do()
}

type QueryCommentListFlow struct {
	userId  int64
	videoId int64

	comments []*models.Comment

	commentList *CommentList
}

func NewQueryCommentListFlow(userId, videoId int64) *QueryCommentListFlow {
	return &QueryCommentListFlow{userId: userId, videoId: videoId}
}

func (q *QueryCommentListFlow) Do() (*CommentList, error) {
	if err := q.checkNum(); err != nil {
		return nil, err
	}
	if err := q.prepareData(); err != nil {
		return nil, err
	}
	if err := q.packData(); err != nil {
		return nil, err
	}
	return q.commentList, nil
}

func (q *QueryCommentListFlow) checkNum() error {
	if !models.NewUserInfoDAO().IsUserExistById(q.userId) {
		return fmt.Errorf("用户%d处于登出状态", q.userId)
	}
	if !models.NewVideoDAO().IsVideoExistById(q.videoId) {
		return fmt.Errorf("视频%d不存在或已经被删除", q.videoId)
	}
	return nil
}

func (q *QueryCommentListFlow) prepareData() error {
	err := models.NewCommentDAO().QueryCommentListByVideoId(q.videoId, &q.comments)
	if err != nil {
		return err
	}
	//根据前端的要求填充正确的时间格式
	err = util.FillCommentListFields(&q.comments)
	if err != nil {
		return err
	}
	return nil
}

func (q *QueryCommentListFlow) packData() error {
	q.commentList = &CommentList{Comments: q.comments}
	return nil
}
