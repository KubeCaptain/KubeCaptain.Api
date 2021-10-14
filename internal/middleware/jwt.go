package middleware

import (
	"KubeCaption.Api/pkg/jwt"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := getToken(c)
		if token == "" {
			unauthorized(c, "token不可为空！")
		}
		claims, err := jwt.ParseToken(token, "ku123test")
		if err != nil {
			unauthorized(c, "token有误，请登陆后访问！")
		}
		if time.Now().Unix() > claims.ExpiresAt {
			unauthorized(c, "token已过期，请重新登陆！")
		}
		c.Next()
	}
}

func getToken(c *gin.Context) string  {
	return c.Query("token")
}

func unauthorized(c *gin.Context, err string)  {
	c.JSON(http.StatusUnauthorized, errors.New(err))
	c.Abort()
	return
}