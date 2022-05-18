package util

import (
	"fmt"
	"github.com/ACking-you/byte_douyin_project/config"
	"github.com/ACking-you/byte_douyin_project/models"
	"log"
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
