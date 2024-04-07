package router

import (
	"github.com/gin-gonic/gin"
	"sylu-oj-gin/internal/app/api"
)

func UserRouter(app *gin.Engine) {
	user := app.Group("/user")
	{
		user.POST("/register", api.Register) // 注册接口
		user.POST("/login", api.Login)       // 登录接口
		user.GET("/logout", api.Logout)      // 登出接口
	}
}