package models

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestSQL(t *testing.T) {
	InitDB()
	err := NewUserInfoDAO().CancelUserFollow(1, 2)
	if err != nil {
		panic(err)
	}
}

func TestAddComment(t *testing.T) {
	InitDB()
	c := Comment{
		Id:         0,
		UserInfoId: 1,
		VideoId:    1,
		User:       UserInfo{},
		Content:    "你好",
	}
	err := NewCommentDAO().AddCommentAndUpdateCount(&c)
	if err != nil {
		panic(err)
	}
}

func TestQueryCommentListByVideoId(t *testing.T) {
	InitDB()
	var comments []*Comment
	err := NewCommentDAO().QueryCommentListByVideoId(1, &comments)
	if err != nil {
		panic(err)
	}
	s, err := json.Marshal(*comments[0])
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v", string(s))
}

func TestUserInfoDAO_QueryUserInfoById(t *testing.T) {
	InitDB()
	var userInfo UserInfo
	err := NewUserInfoDAO().QueryUserInfoById(1, &userInfo)
	if err != nil {
		panic(err)
	}
	s, err := json.Marshal(&userInfo)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", string(s))
}

func TestFormatTime(t *testing.T) {
	tm := time.Now()
	tm = tm.AddDate(0, 6, 10)
	fmt.Println(tm.Format("1-2"))
}
