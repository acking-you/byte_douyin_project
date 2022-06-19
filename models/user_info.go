package models

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"sync"
)

var (
	ErrIvdPtr        = errors.New("空指针错误")
	ErrEmptyUserList = errors.New("用户列表为空")
)

type UserInfo struct {
	Id            int64       `json:"id" gorm:"id,omitempty"`
	Name          string      `json:"name" gorm:"name,omitempty"`
	FollowCount   int64       `json:"follow_count" gorm:"follow_count,omitempty"`
	FollowerCount int64       `json:"follower_count" gorm:"follower_count,omitempty"`
	IsFollow      bool        `json:"is_follow" gorm:"is_follow,omitempty"`
	User          *UserLogin  `json:"-"`                                     //用户与账号密码之间的一对一
	Videos        []*Video    `json:"-"`                                     //用户与投稿视频的一对多
	Follows       []*UserInfo `json:"-" gorm:"many2many:user_relations;"`    //用户之间的多对多
	FavorVideos   []*Video    `json:"-" gorm:"many2many:user_favor_videos;"` //用户与点赞视频之间的多对多
	Comments      []*Comment  `json:"-"`                                     //用户与评论的一对多
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
		return ErrIvdPtr
	}
	//DB.Where("id=?",userId).First(userinfo)
	DB.Where("id=?", userId).Select([]string{"id", "name", "follow_count", "follower_count", "is_follow"}).First(userinfo)
	//id为零值，说明sql执行失败
	if userinfo.Id == 0 {
		return errors.New("该用户不存在")
	}
	return nil
}

func (u *UserInfoDAO) AddUserInfo(userinfo *UserInfo) error {
	if userinfo == nil {
		return ErrIvdPtr
	}
	return DB.Create(userinfo).Error
}

func (u *UserInfoDAO) IsUserExistById(id int64) bool {
	var userinfo UserInfo
	if err := DB.Where("id=?", id).Select("id").First(&userinfo).Error; err != nil {
		log.Println(err)
	}
	if userinfo.Id == 0 {
		return false
	}
	return true
}
func (u *UserInfoDAO) AddUserFollow(userId, userToId int64) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("UPDATE user_infos SET follow_count=follow_count+1 WHERE id = ?", userId).Error; err != nil {
			return err
		}
		if err := tx.Exec("UPDATE user_infos SET follower_count=follower_count+1 WHERE id = ?", userToId).Error; err != nil {
			return err
		}
		if err := tx.Exec("INSERT INTO `user_relations` (`user_info_id`,`follow_id`) VALUES (?,?)", userId, userToId).Error; err != nil {
			return err
		}
		return nil
	})
}

func (u *UserInfoDAO) CancelUserFollow(userId, userToId int64) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("UPDATE user_infos SET follow_count=follow_count-1 WHERE id = ? AND follow_count>0", userId).Error; err != nil {
			return err
		}
		if err := tx.Exec("UPDATE user_infos SET follower_count=follower_count-1 WHERE id = ? AND follower_count>0", userToId).Error; err != nil {
			return err
		}
		if err := tx.Exec("DELETE FROM `user_relations` WHERE user_info_id=? AND follow_id=?", userId, userToId).Error; err != nil {
			return err
		}
		return nil
	})
}

func (u *UserInfoDAO) GetFollowListByUserId(userId int64, userList *[]*UserInfo) error {
	if userList == nil {
		return ErrIvdPtr
	}
	var err error
	if err = DB.Raw("SELECT u.* FROM user_relations r, user_infos u WHERE r.user_info_id = ? AND r.follow_id = u.id", userId).Scan(userList).Error; err != nil {
		return err
	}
	if len(*userList) == 0 || (*userList)[0].Id == 0 {
		return ErrEmptyUserList
	}
	return nil
}

func (u *UserInfoDAO) GetFollowerListByUserId(userId int64, userList *[]*UserInfo) error {
	if userList == nil {
		return ErrIvdPtr
	}
	var err error
	if err = DB.Raw("SELECT u.* FROM user_relations r, user_infos u WHERE r.follow_id = ? AND r.user_info_id = u.id", userId).Scan(userList).Error; err != nil {
		return err
	}
	//if len(*userList) == 0 || (*userList)[0].Id == 0 {
	//	return ErrEmptyUserList
	//}
	return nil
}
