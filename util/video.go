package util

import (
	"errors"
	"fmt"
	"github.com/ACking-you/byte_douyin_project/cache"
	"github.com/ACking-you/byte_douyin_project/config"
	"github.com/ACking-you/byte_douyin_project/models"
	"log"
	"path/filepath"
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

func FillVideoListFields(userId int64, videos *[]*models.Video) (*time.Time, error) {
	size := len(*videos)
	if videos == nil || size == 0 {
		return nil, errors.New("util.FillVideoListFields videos为空")
	}
	dao := models.NewUserInfoDAO()
	p := cache.NewProxyIndexMap()

	latestTime := (*videos)[size-1].CreatedAt //获取最近的投稿时间
	//添加作者信息（后续通过NoSQL优化？
	for i := 0; i < size; i++ {
		var userInfo models.UserInfo
		err := dao.QueryUserInfoById((*videos)[i].UserInfoId, &userInfo)
		if err != nil {
			continue
		}
		(*videos)[i].Author = userInfo
		//填充有登录信息的点赞状态
		if userId > 0 {
			(*videos)[i].IsFavorite = p.GetVideoFavorState(userId, (*videos)[i].Id)
		}
	}
	return &latestTime, nil
}

func SaveImageFromVideo(name string, isDebug bool) error {
	v2i := NewVideo2Image()
	if isDebug {
		v2i.Debug()
	}
	v2i.InputPath = filepath.Join(config.Info.StaticSourcePath, name+defaultVideoSuffix)
	v2i.OutputPath = filepath.Join(config.Info.StaticSourcePath, name+defaultImageSuffix)
	v2i.FrameCount = 1
	queryString, err := v2i.GetQueryString()
	if err != nil {
		return err
	}
	return v2i.ExecCommand(queryString)
}
