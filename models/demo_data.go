package models

const PageCount = 15

//用于简单测试的demo数据

var VideoIndexMap map[int64]*Video

func Init() {
	initVideoIndexMap()
}

var DemoUser = UserInfo{
	Id:            1,
	Name:          "TestUser",
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      false,
}

var DemoVideos = []*Video{
	{
		Id:            1,
		Author:        DemoUser,
		PlayUrl:       "https://www.w3schools.com/html/movie.mp4",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	},
	{
		Id:            2,
		Author:        DemoUser,
		PlayUrl:       "https://media.w3.org/2010/05/sintel/trailer.mp4",
		CoverUrl:      "https://img-blog.csdnimg.cn/4192402992cc4f7294516943d696b1a4.png",
		FavoriteCount: 1,
		CommentCount:  1,
		IsFavorite:    true,
	},
	{
		Id:            3,
		Author:        DemoUser,
		PlayUrl:       "http://192.168.142.251:8080/static/BV1aS4y1j7Hs-rw978CGKX0RWi9zL.mp4",
		CoverUrl:      "https://img-blog.csdnimg.cn/img_convert/c644e86b92d93fcc4173c23ce127e625.png#pic_center",
		FavoriteCount: 1,
		CommentCount:  1,
		IsFavorite:    true,
	},
}

var DemoUsersLoginInfo = map[string]UserInfo{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

func initVideoIndexMap() {
	VideoIndexMap = make(map[int64]*Video, len(DemoVideos))
	for _, v := range DemoVideos {
		VideoIndexMap[v.Id] = v
	}
}
