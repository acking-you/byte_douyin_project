package service

import (
	"fmt"
	"github.com/ACking-you/byte_douyin_project/models"
	"testing"
	"time"
)

func TestJwt(t *testing.T) {
	user := models.UserLogin{
		Id:         1,
		UserInfoId: 1,
		Username:   "L_B__",
		Password:   "08898247",
	}
	token, err := ReleaseToken(user)
	time.Sleep(time.Second * 2)
	if err != nil {
		panic(err)
	}
	_, c, err := ParseToken(token)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v", c)
}
