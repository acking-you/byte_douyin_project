package util

import (
	"errors"
	"github.com/ACking-you/byte_douyin_project/models"
)

func FillCommentListFields(comments *[]*models.Comment) error {
	size := len(*comments)
	if comments == nil || size == 0 {
		return errors.New("util.FillCommentListFields comments为空")
	}
	for i := 0; i < size; i++ {
		(*comments)[i].CreateDate = (*comments)[i].CreatedAt.Format("1-2") //转为前端要求的日期格式
	}
	return nil
}

func FillCommentFields(comment *models.Comment) error {
	if comment == nil {
		return errors.New("FillCommentFields comments为空")
	}
	comment.CreateDate = comment.CreatedAt.Format("1-2") //转为前端要求的日期格式
	return nil
}
