package model

// UserLogin 用户登录表，链接User表
type UserLogin struct {
	ID       int64  `gorm:"primary_key;AUTO_INCREMENT;notnull"`
	Username string `gorm:"notnull;unique"`
	Password string `gorm:"size:255;notnull"`
}
