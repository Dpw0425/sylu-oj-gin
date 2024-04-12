package router

import (
	"github.com/gin-gonic/gin"
	"sylu-oj-gin/internal/app/api"
	"sylu-oj-gin/internal/app/middleware"
)

func ExamRouter(app *gin.Engine) {
	exam := app.Group("/exam")
	{
		exam.Use(middleware.Auth())
		{
			exam.POST("/add_exam", api.AddExam)
		}
	}
}
