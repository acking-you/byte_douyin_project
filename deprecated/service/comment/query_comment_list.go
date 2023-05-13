package comment

import (
	"errors"
	"fmt"
	models2 "github.com/ACking-you/byte_douyin_project/models"
	"github.com/ACking-you/byte_douyin_project/util"
)

type List struct {
	Comments []*models2.Comment `json:"comment_list"`
}

func QueryCommentList(userId, videoId int64) (*List, error) {
	return NewQueryCommentListFlow(userId, videoId).Do()
}

type QueryCommentListFlow struct {
	userId  int64
	videoId int64

	comments []*models2.Comment

	commentList *List
}

func NewQueryCommentListFlow(userId, videoId int64) *QueryCommentListFlow {
	return &QueryCommentListFlow{userId: userId, videoId: videoId}
}

func (q *QueryCommentListFlow) Do() (*List, error) {
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
	if !models2.NewUserInfoDAO().IsUserExistById(q.userId) {
		return fmt.Errorf("用户%d处于登出状态", q.userId)
	}
	if !models2.NewVideoDAO().IsVideoExistById(q.videoId) {
		return fmt.Errorf("视频%d不存在或已经被删除", q.videoId)
	}
	return nil
}

func (q *QueryCommentListFlow) prepareData() error {
	err := models2.NewCommentDAO().QueryCommentListByVideoId(q.videoId, &q.comments)
	if err != nil {
		return err
	}
	//根据前端的要求填充正确的时间格式
	err = util.FillCommentListFields(&q.comments)
	if err != nil {
		return errors.New("暂时还没有人评论")
	}
	return nil
}

func (q *QueryCommentListFlow) packData() error {
	q.commentList = &List{Comments: q.comments}
	return nil
}
