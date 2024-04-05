package app

import (
	"github.com/gin-gonic/gin"
	"sylu-oj-gin/internal/app/config"
	"sylu-oj-gin/internal/app/middleware"
	"sylu-oj-gin/internal/app/router"
)

func InitGinEngine() *gin.Engine {
	gin.SetMode(config.CONFIG.Mode.RunMode)

	app := gin.Default()

	if config.CONFIG.Cors.Enable {
		app.Use(middleware.Cors())
	}

	router.Register(app)

	return app
}
