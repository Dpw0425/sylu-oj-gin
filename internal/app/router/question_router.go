package router

import (
	"github.com/gin-gonic/gin"
	"sylu-oj-gin/internal/app/api"
	"sylu-oj-gin/internal/app/middleware"
)

func QuestionRouter(app *gin.Engine) {
	question := app.Group("/question")
	{
		question.Use(middleware.Auth())
		{
			question.POST("/add_question", api.AddQuestion)
			question.GET("/list", api.QuestionList)
			question.GET("/get_question_msg", api.GetQuestionMsg)
			question.POST("/commit_answer", api.CommitAnswer)
			question.DELETE("/del", api.DelQuestion)
		}
	}
}
