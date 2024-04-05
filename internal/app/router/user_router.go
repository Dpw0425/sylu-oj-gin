package router

import (
	"github.com/gin-gonic/gin"
	"sylu-oj-gin/internal/app/api"
)

func UserRouter(app *gin.Engine) {
	user := app.Group("/user")
	{
		user.POST("register", api.Register)
	}
}
