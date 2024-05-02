package service

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"sylu-oj-gin/internal/app/config"
	"sylu-oj-gin/internal/app/entity"
	"sylu-oj-gin/internal/app/schema"
	"sylu-oj-gin/pkg/error"
)

func AddExam(c *gin.Context, sae schema.AddExam) {
	var ee entity.Exam

	newDB := config.MYSQLDB.Session(&gorm.Session{NewDB: true})
	tx := newDB.Begin()

	for _, student := range sae.Student {
		ee.Name = sae.Name
		ee.Student = student
		result := tx.Table("exams").Create(&ee)
		if result.Error != nil {
			error.Response(c, error.BadRequest, gin.H{}, "创建失败！")
			tx.Rollback()
			return
		}
	}

	tx.Commit()
	error.Response(c, error.OK, gin.H{}, "创建成功！")
}

func AddQuestionToExam(c *gin.Context, saq schema.AddQuestionToExam) {
	var eqe entity.QuestionExam
	var esq entity.StudentQuestion
	var eel = make([]entity.Exam, 0)

	newDB := config.MYSQLDB.Session(&gorm.Session{NewDB: true})
	tx := newDB.Begin()

	result := tx.Table("exams").Where("id = ?", saq.ExamID).Find(&eel)
	if result.Error != nil {
		error.Response(c, error.BadRequest, gin.H{}, "添加失败！")
		tx.Rollback()
		return
	}

	eqe.ExamID = saq.ExamID
	esq.ExamID = saq.ExamID
	for _, qid := range saq.ID {
		eqe.QuestionID = qid
		esq.QuestionID = qid
		result := tx.Table("question_exams").Create(&eqe)
		if result.Error != nil {
			error.Response(c, error.BadRequest, gin.H{}, "添加失败！")
			tx.Rollback()
			return
		}

		for _, stu := range eel {
			esq.Username = stu.Student
			result1 := tx.Table("student_questions").Create(&esq)
			if result1.Error != nil {
				error.Response(c, error.BadRequest, gin.H{}, "添加失败！")
				tx.Rollback()
				return
			}
		}
	}

	tx.Commit()
	error.Response(c, error.OK, gin.H{}, "添加成功！")
}

func Inspect(c *gin.Context, eid int) {
	var esql = make([]entity.StudentQuestion, 0)
	var sesrl = make([]schema.ExamStatusResp, 0)
	var sesr schema.ExamStatusResp
	result := config.MYSQLDB.Table("student_questions").Where("exam_id = ?", eid).Find(&esql)
	if result.Error != nil {
		error.Response(c, error.BadRequest, gin.H{}, "查询失败！")
		return
	}

	for _, esql := range esql {
		sesr.Username = esql.Username
		sesr.Status = esql.Status
		result1 := config.MYSQLDB.Table("users").Where("username = ?", esql.Username).First(&sesr.ID)
		if result1.Error != nil {
			error.Response(c, error.BadRequest, gin.H{}, "查询失败！")
			return
		}
		sesrl = append(sesrl, sesr)
	}

	error.Response(c, error.OK, gin.H{"student": sesrl}, "查询成功！")
}
