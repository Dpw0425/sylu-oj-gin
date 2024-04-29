package middleware

import (
	"github.com/gin-gonic/gin"
	"sylu-oj-gin/pkg/error"
	"sylu-oj-gin/pkg/utils"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 获取 Token
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			error.Response(c, error.Unauth, gin.H{}, "未登录！")
			c.Abort()
			return
		}

		// 解析 Token
		token, claims, err := utils.ParseToken(tokenString)
		if err != nil || !token.Valid {
			error.Response(c, error.Unauth, gin.H{}, "登录已过期！")
			c.Abort()
			return
		}

		// 通过验证, 获取 claims 中的 UserID
		uid := claims.UID

		c.Set("UserID", uid)
		c.Next()
	}
}
