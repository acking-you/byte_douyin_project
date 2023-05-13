package middleware

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/gin-gonic/gin"
)

func SHA1(s string) string {

	o := sha1.New()

	o.Write([]byte(s))

	return hex.EncodeToString(o.Sum(nil))
}

func SHAMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		password := context.Query("password")
		if password == "" {
			password = context.PostForm("password")
		}
		context.Set("password", SHA1(password))
		context.Next()
	}
}
