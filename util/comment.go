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
	dao := models.NewUserInfoDAO()
	for _, v := range *comments {
		_ = dao.QueryUserInfoById(v.UserInfoId, &v.User) //填充这条评论的作者信息
		v.CreateDate = v.CreatedAt.Format("1-2")         //转为前端要求的日期格式
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
