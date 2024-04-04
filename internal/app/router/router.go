package router

import (
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Register(app *gin.Engine) error {
	RegisterAPI(app)
	return nil
}

func RegisterAPI(app *gin.Engine) {
	// TODO: limit route

	swagAny := func(c *gin.Context) { ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "SWAGGER")(c) }
	app.GET("/swagger/*any", swagAny)

	TestRouters(app)
}
