package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"sylu-oj-gin/internal/app/config"
	"sylu-oj-gin/internal/app/middleware"
)

func Register(app *gin.Engine) error {
	RegisterAPI(app)
	return nil
}

func RegisterAPI(app *gin.Engine) {
	app.Use(middleware.LimitRoute(config.CONFIG.Limit.Limit))
	{
		swagAny := func(c *gin.Context) { ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "SWAGGER")(c) }
		app.GET("/swagger/*any", swagAny)

		TestRouters(app)
		UserRouter(app)
		QuestionRouter(app)
		ExamRouter(app)
	}
}
