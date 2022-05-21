package cache

// 用户id->视频id->是否被点赞
var userFavorIndexMap map[int64]map[int64]bool

func init() {
	userFavorIndexMap = make(map[int64]map[int64]bool)
}

var (
	proxyIndexOperation ProxyIndexMap
)

type ProxyIndexMap struct {
}

func NewProxyIndexMap() *ProxyIndexMap {
	return &proxyIndexOperation
}

// UpdateVideoFavorState 更新点赞状态，注意go里面的map是类似于指针的结构，需要手动申请内存
func (i *ProxyIndexMap) UpdateVideoFavorState(userId int64, videoId int64, state bool) {
	//如果未初始化内存，则进行内存的初始化
	if _, ok := userFavorIndexMap[userId]; !ok {
		userFavorIndexMap[userId] = make(map[int64]bool)
	}
	userFavorIndexMap[userId][videoId] = state
}

func (i *ProxyIndexMap) GetVideoFavorState(userId int64, videoId int64) bool {
	f, ok := userFavorIndexMap[userId]
	if !ok {
		return false
	}
	state, ok := f[videoId]
	if !ok {
		return false
	}
	return state
}
