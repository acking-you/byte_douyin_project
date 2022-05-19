package util

import (
	"errors"
	"fmt"
	"github.com/ACking-you/byte_douyin_project/config"
	"github.com/ACking-you/byte_douyin_project/models"
	"log"
	"time"
)

func GetFileUrl(fileName string) string {
	base := fmt.Sprintf(`http://%s:%d/static/%s`, config.Info.IP, config.Info.Port, fileName)
	return base
}

// NewFileName 根据userId+用户发布的视频数量连接成独一无二的文件名
func NewFileName(userId int64) string {
	var count int64

	err := models.NewVideoDAO().QueryVideoCountByUserId(userId, &count)
	if err != nil {
		log.Println(err)
	}
	return fmt.Sprintf("%d-%d", userId, count)
}

func FillVideoAuthor(videos *[]*models.Video) (*time.Time, error) {
	size := len(*videos)
	if videos == nil || size == 0 {
		return nil, errors.New("util.FillVideoAuthor videos为空")
	}
	dao := models.NewUserInfoDAO()
	var userInfo models.UserInfo
	latestTime := (*videos)[size-1].CreatedAt //获取最近的投稿时间
	//添加作者信息（后续通过NoSQL优化？
	for i := 0; i < size; i++ {
		err := dao.QueryUserInfoById((*videos)[i].UserInfoId, &userInfo)
		if err != nil {
			continue
		}
		(*videos)[i].Author = userInfo
	}
	return &latestTime, nil
}
