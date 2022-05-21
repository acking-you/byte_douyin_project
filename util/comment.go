package util

import (
	"errors"
	"github.com/ACking-you/byte_douyin_project/models"
)

func FillCommentField(comments *[]*models.Comment) error {
	size := len(*comments)
	if comments == nil || size == 0 {
		return errors.New("util.FillCommentField comments为空")
	}
	for i := 0; i < size; i++ {
		(*comments)[i].CreateDate = (*comments)[i].CreatedAt.Format("1-2") //转为前端要求的日期格式
	}
	return nil
}
