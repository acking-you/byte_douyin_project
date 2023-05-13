package user_info

import (
	"errors"
	"github.com/ACking-you/byte_douyin_project/models"
)

var (
	ErrUserNotExist = errors.New("用户不存在或已注销")
)

type FollowList struct {
	UserList []*models.UserInfo `json:"user_list"`
}

func QueryFollowList(userId int64) (*FollowList, error) {
	return NewQueryFollowListFlow(userId).Do()
}

type QueryFollowListFlow struct {
	userId int64

	userList []*models.UserInfo

	*FollowList
}

func NewQueryFollowListFlow(userId int64) *QueryFollowListFlow {
	return &QueryFollowListFlow{userId: userId}
}

func (q *QueryFollowListFlow) Do() (*FollowList, error) {
	var err error
	if err = q.checkNum(); err != nil {
		return nil, err
	}
	if err = q.prepareData(); err != nil {
		return nil, err
	}
	if err = q.packData(); err != nil {
		return nil, err
	}

	return q.FollowList, nil
}

func (q *QueryFollowListFlow) checkNum() error {
	if !models.NewUserInfoDAO().IsUserExistById(q.userId) {
		return ErrUserNotExist
	}
	return nil
}

func (q *QueryFollowListFlow) prepareData() error {
	var userList []*models.UserInfo
	err := models.NewUserInfoDAO().GetFollowListByUserId(q.userId, &userList)
	if err != nil {
		return err
	}
	for i, _ := range userList {
		userList[i].IsFollow = true //当前用户的关注列表，故isFollow定为true
	}
	q.userList = userList
	return nil
}

func (q *QueryFollowListFlow) packData() error {
	q.FollowList = &FollowList{UserList: q.userList}

	return nil
}
