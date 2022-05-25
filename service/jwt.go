package service

import (
	"errors"
	"github.com/ACking-you/byte_douyin_project/models"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("acking-you.xyz")

type Claims struct {
	UserId int64
	jwt.StandardClaims
}

// ReleaseToken 颁发token
func ReleaseToken(user models.UserLogin) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: user.UserInfoId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "douyin_pro_131",
			Subject:   "L_B__",
		}}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ParseToken 解析token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, claims, err
}

// JWTAuth 用于验证token，并返回token对应的userid
func JWTAuth(token string) (int64, error) {
	if token == "" {
		return 0, errors.New("token为空")
	}
	_, claim, err := ParseToken(token)
	if err != nil {
		return 0, errors.New("token过期")
	}
	//最后验证这个user是否真的存在
	if !models.NewUserInfoDAO().IsUserExistById(claim.UserId) {
		return 0, errors.New("user不存在")
	}

	return claim.UserId, nil
}
