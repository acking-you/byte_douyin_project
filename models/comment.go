package models

type Comment struct {
	Id         int64    `json:"id,omitempty"`
	UserInfoId uint     `json:"-"` //用于一对多关系的id
	User       UserInfo `json:"user" gorm:"-"`
	Content    string   `json:"content,omitempty"`
	CreateDate string   `json:"create_date,omitempty"`
}
