package video

import (
	"github.com/ACking-you/byte_douyin_project/models"
	"github.com/ACking-you/byte_douyin_project/util"
)

// PostVideo 投稿视频
func PostVideo(userId int64, videoName, coverName, title string) error {
	return NewPostVideoFlow(userId, videoName, coverName, title).Do()
}

func NewPostVideoFlow(userId int64, videoName, coverName, title string) *PostVideoFlow {
	return &PostVideoFlow{
		videoName: videoName,
		coverName: coverName,
		userId:    userId,
		title:     title,
	}
}

type PostVideoFlow struct {
	videoName string
	coverName string
	title     string
	userId    int64

	video *models.Video
}

func (f *PostVideoFlow) Do() error {
	f.prepareParam()

	if err := f.publish(); err != nil {
		return err
	}
	return nil
}

//准备好参数
func (f *PostVideoFlow) prepareParam() {
	f.videoName = util.GetFileUrl(f.videoName)
	f.coverName = util.GetFileUrl(f.coverName)
}

//组合并添加到数据库
func (f *PostVideoFlow) publish() error {
	video := &models.Video{
		UserInfoId: f.userId,
		PlayUrl:    f.videoName,
		CoverUrl:   f.coverName,
		Title:      f.title,
	}
	return models.NewVideoDAO().AddVideo(video)
}
