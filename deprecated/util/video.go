package util

import (
	"errors"
	"fmt"
	"github.com/ACking-you/byte_douyin_project/cache"
	"github.com/ACking-you/byte_douyin_project/config"
	models2 "github.com/ACking-you/byte_douyin_project/models"
	"log"
	"path/filepath"
	"time"
)

func GetFileUrl(fileName string) string {
	base := fmt.Sprintf("http://%s:%d/static/%s", config.Global.IP, config.Global.Port, fileName)
	return base
}

// NewFileName 根据userId+用户发布的视频数量连接成独一无二的文件名
func NewFileName(userId int64) string {
	var count int64

	err := models2.NewVideoDAO().QueryVideoCountByUserId(userId, &count)
	if err != nil {
		log.Println(err)
	}
	return fmt.Sprintf("%d-%d", userId, count)
}

// FillVideoListFields 填充每个视频的作者信息（因为作者与视频的一对多关系，数据库中存下的是作者的id
// 当userId>0时，我们判断当前为登录状态，其余情况为未登录状态，则不需要填充IsFavorite字段
func FillVideoListFields(userId int64, videos *[]*models2.Video) (*time.Time, error) {
	if videos == nil || (len(*videos) == 0) {
		return nil, errors.New("util.FillVideoListFields videos为空")
	}
	size := len(*videos)
	dao := models2.NewUserInfoDAO()
	p := cache.NewProxyIndexMap()

	latestTime := (*videos)[size-1].CreatedAt //获取最近的投稿时间
	//添加作者信息，以及is_follow状态
	for i := 0; i < size; i++ {
		var userInfo models2.UserInfo
		err := dao.QueryUserInfoById((*videos)[i].UserInfoId, &userInfo)
		if err != nil {
			continue
		}
		userInfo.IsFollow = p.GetUserRelation(userId, userInfo.Id) //根据cache更新是否被点赞
		(*videos)[i].Author = userInfo
		//填充有登录信息的点赞状态
		if userId > 0 {
			(*videos)[i].IsFavorite = p.GetVideoFavorState(userId, (*videos)[i].Id)
		}
	}
	return &latestTime, nil
}

// SaveImageFromVideo 将视频切一帧保存到本地
// isDebug用于控制是否打印出执行的ffmepg命令
func SaveImageFromVideo(name string, isDebug bool) error {
	return NewVideo2Image().
		SetInputPath(filepath.Join(config.Global.StaticSourcePath, name+GetDefaultVideoSuffix())).
		SetOutputPath(filepath.Join(config.Global.StaticSourcePath, name+GetDefaultImageSuffix())).
		SetFrameCount(1).
		SetDebug(isDebug).Execute()
}
