package controller

import (
	"TiktokServer/middleware"
	jwt2 "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

//创建一个token
func generateToken(c *gin.Context, username string) (token string, err error) {
	// 构造SignKey: 签名和解签名需要使用一个值
	jwt := middleware.NewJWT()
	// 构造用户claims信息(负荷)
	claims := middleware.CustomClaims{
		Name: username,
		StandardClaims: jwt2.StandardClaims{
			NotBefore: time.Now().Unix() - 1000, // 签名生效时间
			ExpiresAt: time.Now().Unix() + 3600, // 签名过期时间
			Issuer:    "wxf.top",                // 签名颁发者
		},
	}
	// 根据claims生成token对象
	token, err = jwt.CreateToken(claims)
	if err != nil {
		return
	}
	log.Println("create token", token)
	return
}
