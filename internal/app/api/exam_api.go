package api

import (
	"github.com/gin-gonic/gin"
	"strconv"
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

	id, _ := c.Get("UserID")
	uid := id.(int)

	service.AddExam(c, sae, uid)
}

// AddQuestionToExam @Summary 添加题目接口
// @Description 添加题目到实验
// @Param request body schema.AddQuestionToExam true "exam id and question id"
// @Accept       json
// @Produce      json
// @Success      200  {object}  error.ResponseNormal
// @Router       /exam/add_question [post]
func AddQuestionToExam(c *gin.Context) {
	var saq schema.AddQuestionToExam
	if err := c.ShouldBindJSON(&saq); err != nil {
		error.Response(c, error.BadRequest, gin.H{}, "输入内容有误！")
		return
	}

	service.AddQuestionToExam(c, saq)
}

// Inspect @Summary 检查进度接口
// @Description 检查实验内题目的完成度
// @Param id query int true "exam id"
// @Accept       json
// @Produce      json
// @Success      200  {object}  error.ResponseNormal
// @Router       /exam/inspect [get]
func Inspect(c *gin.Context) {
	eid, _ := strconv.Atoi(c.Query("eid"))
	qid, _ := strconv.Atoi(c.Query("qid"))

	service.Inspect(c, eid, qid)
}

// ExamList @Summary 获取实验列表
// @Description 查看有关当前用户的全部实验
// @Param page query int true "page number"
// @Param number query int true "number per page"
// @Accept       json
// @Produce      json
// @Success      200  {object}  error.ResponseNormal
// @Router       /exam/list [get]
func ExamList(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	number, _ := strconv.Atoi(c.Query("number"))

	id, _ := c.Get("UserID")
	uid := id.(int)

	service.ExamList(c, uid, page, number)
}
