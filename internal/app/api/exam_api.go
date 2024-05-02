package api

import (
	"github.com/gin-gonic/gin"
	"sylu-oj-gin/internal/app/schema"
	"sylu-oj-gin/internal/app/service"
	"sylu-oj-gin/pkg/error"
)

// AddExam @Summary 创建实验接口
// @Description 创建实验
// @Param request body schema.AddExam true "exam message"
// @Accept       json
// @Produce      json
// @Success      200  {object}  error.ResponseNormal
// @Router       /exam/add_exam [post]
func AddExam(c *gin.Context) {
	var sae schema.AddExam
	if err := c.ShouldBindJSON(&sae); err != nil {
		error.Response(c, error.BadRequest, gin.H{}, "输入内容有误！")
		return
	}

	service.AddExam(c, sae)
}

func AddQuestionToExam(c *gin.Context) {
	var saq schema.AddQuestionToExam
	if err := c.ShouldBindJSON(&saq); err != nil {
		error.Response(c, error.BadRequest, gin.H{}, "输入内容有误！")
		return
	}

	service.AddQuestionToExam(c, saq)
}
