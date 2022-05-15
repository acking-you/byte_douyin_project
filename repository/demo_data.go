package repository

const PageCount = 15


type categoryWeightItem struct {
	Weight int
	Category string
}

var VideoIndexMap map[int64]*Video

func Init()  {
	initVideoIndexMap()
}

var DemoUser = User{
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
		Id: 2,
		Author: DemoUser,
		PlayUrl: "https://media.w3.org/2010/05/sintel/trailer.mp4",
		CoverUrl: "https://w.wallhaven.cc/full/g7/wallhaven-g7gj3d.jpg",
		FavoriteCount: 1,
		CommentCount: 1,
		IsFavorite: true,
	},
}

func initVideoIndexMap() {
	VideoIndexMap = make(map[int64]*Video,len(DemoVideos))
	for _,v := range DemoVideos{
		VideoIndexMap[v.Id] = v
	}
}