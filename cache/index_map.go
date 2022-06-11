package cache

import (
	"context"
	"fmt"
	"github.com/ACking-you/byte_douyin_project/config"
	"github.com/go-redis/redis/v8"
)

// 用户id->被点赞的视频id集合->是否含有该视频id

var ctx = context.Background()
var rdb *redis.Client

const (
	favor    = "favor"
	relation = "relation"
)

func init() {
	rdb = redis.NewClient(
		&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", config.Info.RDB.IP, config.Info.RDB.Port),
			Password: "", //没有设置密码
			DB:       config.Info.RDB.Database,
		})
}

var (
	proxyIndexOperation ProxyIndexMap
)

type ProxyIndexMap struct {
}

func NewProxyIndexMap() *ProxyIndexMap {
	return &proxyIndexOperation
}

// UpdateVideoFavorState 更新点赞状态，state:true为点赞，false为取消点赞
func (i *ProxyIndexMap) UpdateVideoFavorState(userId int64, videoId int64, state bool) {
	key := fmt.Sprintf("%s:%d", favor, userId)
	if state {
		rdb.SAdd(ctx, key, videoId)
		return
	}
	rdb.SRem(ctx, key, videoId)
}

// GetVideoFavorState 得到点赞状态
func (i *ProxyIndexMap) GetVideoFavorState(userId int64, videoId int64) bool {
	key := fmt.Sprintf("%s:%d", favor, userId)
	ret := rdb.SIsMember(ctx, key, videoId)
	return ret.Val()
}

// UpdateUserRelation 更新点赞状态，state:true为点关注，false为取消关注
func (i *ProxyIndexMap) UpdateUserRelation(userId int64, followId int64, state bool) {
	key := fmt.Sprintf("%s:%d", relation, userId)
	if state {
		rdb.SAdd(ctx, key, followId)
		return
	}
	rdb.SRem(ctx, key, followId)
}

// GetUserRelation 得到关注状态
func (i *ProxyIndexMap) GetUserRelation(userId int64, followId int64) bool {
	key := fmt.Sprintf("%s:%d", relation, userId)
	ret := rdb.SIsMember(ctx, key, followId)
	return ret.Val()
}
