package api

import (
	"github.com/gin-gonic/gin"
	"sylu-oj-gin/internal/app/schema"
	"sylu-oj-gin/internal/app/service"
	"sylu-oj-gin/pkg/error"
)

// @Summary 发布题目接口
// @Description 发布题目
// @Param request body schema.AddQuestion true "question message"
// @Accept       json
// @Produce      json
// @Success      200  {object}  error.ResponseNormal
// @Router       /question/add_question [post]
func AddQuestion(c *gin.Context) {
	var saq schema.AddQuestion
	if err := c.ShouldBindJSON(&saq); err != nil {
		error.Response(c, error.BadRequest, gin.H{}, "输入内容有误！")
		return
	}

	id, _ := c.Get("UserID")
	uid := id.(uint)

	service.AddQuestion(c, saq, uid)
}
