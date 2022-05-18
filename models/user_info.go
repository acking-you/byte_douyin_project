package models

import (
	"errors"
	"log"
	"sync"
)

type UserInfo struct {
	Id            int64      `json:"id,omitempty"`
	Name          string     `json:"name,omitempty"`
	FollowCount   int64      `json:"follow_count,omitempty"`
	FollowerCount int64      `json:"follower_count,omitempty"`
	IsFollow      bool       `json:"is_follow,omitempty"`
	User          *UserLogin `json:"-"` //数据库中一对一的关系
	Videos        []*Video   `json:"-"` //数据库中一对多关系
	Comments      []*Comment `json:"-"` //数据库中一对多关系
}

type UserInfoDAO struct {
}

var (
	userInfoDAO  *UserInfoDAO
	userInfoOnce sync.Once
)

func NewUserInfoDAO() *UserInfoDAO {
	userInfoOnce.Do(func() {
		userInfoDAO = new(UserInfoDAO)
	})
	return userInfoDAO
}

func (u *UserInfoDAO) QueryUserInfoById(userId int64, userinfo *UserInfo) error {
	if userinfo == nil {
		return errors.New("userinfo 指针为空")
	}
	DB.Where("id=?", userId).First(userinfo)
	//id为零值，说明sql执行失败
	if userinfo.Id == 0 {
		return errors.New("该用户不存在")
	}
	return nil
}

func (u *UserInfoDAO) AddUserInfo(userinfo *UserInfo) error {
	if userinfo == nil {
		return errors.New("userinfo 空指针")
	}
	return DB.Create(userinfo).Error
}

func (u *UserInfoDAO) IsUserExistById(id int64) bool {
	var userinfo UserInfo
	if err := DB.Where("id=?", id).First(&userinfo).Error; err != nil {
		log.Println(err)
	}
	if userinfo.Id == 0 {
		return false
	}
	return true
}
