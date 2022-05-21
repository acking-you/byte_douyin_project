package models

import (
	"errors"
	"time"
)

type Comment struct {
	Id         int64    `json:"id,omitempty"`
	UserInfoId int64    `json:"-"` //用于一对多关系的id
	VideoId    int64    `json:"-"` //一对多，视频对评论
	User       UserInfo `json:"user" gorm:"-"`
	Content    string   `json:"content,omitempty"`
	CreatedAt  time.Time
	CreateDate string `json:"create_date" gorm:"-"`
}

type CommentDAO struct {
}

var (
	commentDao CommentDAO
)

func NewCommentDAO() *CommentDAO {
	return &commentDao
}

func (c *CommentDAO) AddComment(comment *Comment) error {
	if comment == nil {
		return errors.New("AddComment comment空指针")
	}
	return DB.Create(comment).Error
}

func (c *CommentDAO) QueryCommentListByVideoId(videoId int64, comments *[]*Comment) error {
	if comments == nil {
		return errors.New("QueryCommentListByVideoId comments空指针")
	}
	if err := DB.Model(&Comment{}).Where("video_id=?", videoId).Find(comments).Error; err != nil {
		return err
	}
	return nil
}
