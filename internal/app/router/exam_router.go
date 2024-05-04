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
			exam.POST("/add_exam", api.AddExam)               // 创建实验
			exam.POST("/add_question", api.AddQuestionToExam) // 向实验中添加题目
			exam.GET("/inspect", api.Inspect)                 // 检查进度
			exam.GET("/list", api.ExamList)                   // 实验列表
		}
	}
}
