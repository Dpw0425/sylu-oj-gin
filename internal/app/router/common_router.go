package router

import (
	"github.com/gin-gonic/gin"
	"sylu-oj-gin/internal/app/api"
)

func TestRouters(app *gin.Engine) {
	test := app.Group("/test")
	{
		test.GET("/ping", api.Ping)
	}
}
