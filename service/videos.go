package service

import (
	"github.com/ACking-you/byte_douyin_project/models"
)

type VideoList struct {
	Videos []*models.Video `json:"video_list,omitempty"`
}

func QueryVideoListDemo() (*VideoList, error) {
	return &VideoList{Videos: models.DemoVideos}, nil
	//videos := &VideoList{}
	//v1 := models.GetVideoById(1)
	//if v1 == nil {
	//	return videos,errors.New("video1 get error")
	//}
	//v2 := models.GetVideoById(2)
	//if v2 == nil {
	//	return videos,errors.New("video2 get error")
	//}
	//videos.Videos = append(videos.Videos,v1)
	//videos.Videos = append(videos.Videos,v2)
	//return videos,nil
}
