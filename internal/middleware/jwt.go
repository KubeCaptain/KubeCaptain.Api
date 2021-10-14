package middleware

import (
	"KubeCaption.Api/pkg/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := getToken(c)
		if token == "" {
			unauthorized(c, "token不可为空！")
			return
		}
		claims, err := jwt.ParseToken(token, "ku123test")
		if err != nil {
			unauthorized(c, "token有误，请登陆后访问！")
			return
		}
		if time.Now().Unix() > claims.ExpiresAt {
			unauthorized(c, "token已过期，请重新登陆！")
			return
		}
		c.Next()
	}
}

func getToken(c *gin.Context) (token string)  {
	token = c.Query("token")
	if token != "" {
		return
	}
	headerAuth := strings.Split(c.GetHeader("Authorization")," ")
	if len(headerAuth) >= 2 || headerAuth[0] == "Bearer"  {
		token = headerAuth[1]
	}
	return
}

func unauthorized(c *gin.Context, err string)  {
	c.JSON(http.StatusUnauthorized, gin.H{ "message": err})
	c.Abort()
}