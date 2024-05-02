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

	newDB := config.MYSQLDB.Session(&gorm.Session{NewDB: true})
	tx := newDB.Begin()

	eqe.ExamID = saq.ExamID
	for _, qid := range saq.ID {
		eqe.QuestionID = qid
		result := tx.Table("question_exams").Create(&eqe)
		if result.Error != nil {
			error.Response(c, error.BadRequest, gin.H{}, "添加失败！")
			tx.Rollback()
			return
		}
	}

	tx.Commit()
	error.Response(c, error.OK, gin.H{}, "添加成功！")
}
