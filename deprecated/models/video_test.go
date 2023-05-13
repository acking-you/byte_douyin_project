package models

import (
	"fmt"
	"testing"
	"time"
)

func TestVideoDAO_QueryVideoListByUserId(t *testing.T) {
	InitDB()
	s := make([]*Video, 8)
	err := NewVideoDAO().QueryVideoListByUserId(1, &s)
	if err != nil {
		panic(err)
	}
	for _, v := range s {
		fmt.Printf("%#v\n", *v)
	}
}

func TestVideoDAO_QueryVideoListByLimit(t *testing.T) {
	InitDB()
	s := make([]*Video, 8)
	err := NewVideoDAO().QueryVideoListByLimitAndTime(2, time.Unix(1652895580, 0), &s)
	if err != nil {
		panic(err)
	}
	for _, v := range s {
		fmt.Printf("%#v\n", *v)
	}
}

func TestTime(t *testing.T) {
	println(time.Now().UnixNano())
}
