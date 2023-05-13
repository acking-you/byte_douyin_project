package cache

import (
	"fmt"
	"testing"
)

func TestProxyIndexMap_UpdateUserRelation(t *testing.T) {
	NewProxyIndexMap().UpdateUserRelation(1, 2, true)
	fmt.Println(NewProxyIndexMap().GetUserRelation(1, 3))
}

func TestProxyIndexMap_GetVideoFavorState(t *testing.T) {
	fmt.Println(NewProxyIndexMap().GetVideoFavorState(1, 2))
	NewProxyIndexMap().UpdateVideoFavorState(1, 2, true)
	fmt.Println(NewProxyIndexMap().GetVideoFavorState(1, 2))
}
