package models

import (
	"testing"
)

func TestInitDB(t *testing.T) {
	//InitDB()

	//v := &Video{
	//	Title: "hhh",
	//}
	var userInfo = UserInfo{
		//Videos: []*Video{v},
	}
	DB.Where("id=?", 1).Save(&userInfo)
	//for _,video := range videos{
	//	fmt.Println(*video)
	//}
}
