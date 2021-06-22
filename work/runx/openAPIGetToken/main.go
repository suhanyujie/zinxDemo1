package main

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// 只用于内部测试使用。给测试同事提供接口，获取三方组织的 token，以进行进一步的接口测试。
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// http://192.168.88.40:8081/getToken/c10ec58587dxxxxxkey/4f6d623f7e9243dd9570f0xxxsecret
	r.GET("/getToken/:key/:secret", func(c *gin.Context) {
		key := c.Param("key")
		if key == "" {
			c.JSON(200, gin.H{
				"message": "请传入 key",
			})
			return
		}
		secret := c.Param("secret")
		if secret == "" {
			c.JSON(200, gin.H{
				"message": "请传入 secret",
			})
			return
		}
		newToken := GenToken(key, secret)
		c.JSON(200, gin.H{
			"message": "success!",
			"data":    newToken,
		})
		return
	})
	r.Run(":8081")
}

func GenToken(key, secret string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["_key"] = key
	token.Claims = claims
	accessToken, _ := token.SignedString([]byte(secret))
	// fmt.Println(accessToken)
	return accessToken
}
