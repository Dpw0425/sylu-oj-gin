package api

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"sylu-oj-gin/internal/app/schema"
	"sylu-oj-gin/internal/app/service"
	"sylu-oj-gin/pkg/error"
)

// AddQuestion @Summary 发布题目接口
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
	uid := id.(int)

	service.AddQuestion(c, saq, uid)
}

// QuestionList @Summary 查看接口
// @Description 查看题目列表
// @Param page query number true "page number"
// @Param number query number true "rows per page"
// @Param search_title query string false "question title"
// @Param search_tag query string false "question tag"
// @Param search_degree query number false "question degree"
// @Param order query number false "reorder"
// @Accept       json
// @Produce      json
// @Success      200  {object}  error.ResponseNormal
// @Router       /question/list [get]
func QuestionList(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	number, _ := strconv.Atoi(c.Query("number"))
	searchTitle := c.Query("search_title")
	searchTag := c.Query("search_tag")
	searchDegree, _ := strconv.Atoi(c.Query("search_degree"))
	order, _ := strconv.Atoi(c.Query("order"))

	service.QuestionList(c, page, number, searchTitle, searchTag, searchDegree, order)
}

// GetQuestionMsg @Summary 查看题目接口
// @Description 查看题目详情
// @Param id query int true "get question message"
// @Accept       json
// @Produce      json
// @Success      200  {object}  error.ResponseNormal
// @Router       /question/get_question_msg [get]
func GetQuestionMsg(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))

	service.GetQuestionMsg(c, id)
}

// CommitAnswer @Summary 查看题目接口
// @Description 查看题目详情
// @Param answer body schema.Answer true "commit your answer of question"
// @Accept       json
// @Produce      json
// @Success      200  {object}  error.ResponseNormal
// @Router       /question/commit_answer [post]
func CommitAnswer(c *gin.Context) {
	var sa schema.Answer
	if err := c.ShouldBindJSON(&sa); err != nil {
		error.Response(c, error.BadRequest, gin.H{}, "输入内容有误！")
		return
	}

	id, _ := c.Get("UserID")
	uid := id.(int)

	service.CommitAnswer(c, sa, uid)
}

// DelQuestion @Summary 删除题目接口
// @Description 删除题目
// @Param id query int true "question id"
// @Accept       json
// @Produce      json
// @Success      200  {object}  error.ResponseNormal
// @Router       /question/del [delete]
func DelQuestion(c *gin.Context) {
	qid, _ := strconv.Atoi(c.Query("id"))

	id, _ := c.Get("UserID")
	uid := id.(int)

	service.DelQuestion(c, qid, uid)
}
