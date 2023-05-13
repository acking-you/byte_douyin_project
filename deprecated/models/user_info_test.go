package models

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	InitDB()
	code := m.Run()
	os.Exit(code)
}

func TestUserInfoDAO_GetFollowListByUserId(t *testing.T) {
	var userList []*UserInfo
	err := NewUserInfoDAO().GetFollowListByUserId(1, &userList)
	if err != nil {
		panic(err)
	}
	for _, user := range userList {
		fmt.Printf("%#v\n", *user)
	}
}

func TestUserInfoDAO_GetFollowerListByUserId(t *testing.T) {
	var userList []*UserInfo
	err := NewUserInfoDAO().GetFollowerListByUserId(2, &userList)
	if err != nil {
		panic(err)
	}
	for _, user := range userList {
		fmt.Printf("%#v\n", *user)
	}
}
