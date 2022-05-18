package handlers

import (
	"encoding/json"
	"github.com/ACking-you/byte_douyin_project/models"
	"testing"
)

func TestFeedHandler(t *testing.T) {
	models.Init()
	data, err := json.Marshal(getFeed())
	if err != nil {
		return
	}
	println(string(data))
}
