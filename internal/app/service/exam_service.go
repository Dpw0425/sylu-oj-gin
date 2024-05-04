package service

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"sylu-oj-gin/internal/app/config"
	"sylu-oj-gin/internal/app/entity"
	"sylu-oj-gin/internal/app/schema"
	"sylu-oj-gin/pkg/error"
	"sylu-oj-gin/pkg/utils"
)

func AddExam(c *gin.Context, sae schema.AddExam, uid int) {
	var ee entity.Exam

	ee.Student = utils.ArrToString(sae.Student)
	ee.Name = sae.Name
	ee.OwnerID = uid
	result := config.MYSQLDB.Table("exams").Create(&ee)
	if result.Error != nil {
		error.Response(c, error.BadRequest, gin.H{}, "创建失败！")
		return
	}

	error.Response(c, error.OK, gin.H{}, "创建成功！")
}

func AddQuestionToExam(c *gin.Context, saq schema.AddQuestionToExam) {
	var eqe entity.QuestionExam
	var esq entity.StudentQuestion
	var ee entity.Exam

	newDB := config.MYSQLDB.Session(&gorm.Session{NewDB: true})
	tx := newDB.Begin()

	result := tx.Table("exams").Where("id = ?", saq.ExamID).First(&ee)
	if result.Error != nil {
		error.Response(c, error.BadRequest, gin.H{}, "添加失败！")
		tx.Rollback()
		return
	}

	eqe.ExamID = saq.ExamID
	esq.ExamID = saq.ExamID
	stuList := utils.StringToArr(ee.Student)
	for _, qid := range saq.ID {
		eqe.QuestionID = qid
		esq.QuestionID = qid
		result := tx.Table("question_exams").Create(&eqe)
		if result.Error != nil {
			error.Response(c, error.BadRequest, gin.H{}, "添加失败！")
			tx.Rollback()
			return
		}

		for _, stu := range stuList {
			esq.Username = stu
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
	var sesrl = make([]schema.ExamStatusResp, 0)
	var sesr schema.ExamStatusResp

	var passList []string
	// 获取实验内的全部学生
	result := config.MYSQLDB.Table("student_questions").
		Select("username").
		Where("status = ? AND exam_id = ?", "pass", eid).
		Group("username").
		Having("COUNT(DISTINCT question_id) = ?", config.MYSQLDB.Table("student_questions").Where("status = ?", "pass").Select("COUNT(DISTINCT question_id)")).
		Pluck("username", &passList)
	if result.Error != nil {
		error.Response(c, error.BadRequest, gin.H{}, "查询失败！")
		return
	}

	var failList []string
	result1 := config.MYSQLDB.Table("student_questions").
		Select("username").
		Where("status != ? AND exam_id = ?", "pass", eid).
		Not(config.MYSQLDB.Table("student_questions").Where("status = ?", "pass").Select("username")).
		Pluck("username", &failList)
	if result1.Error != nil {
		error.Response(c, error.BadRequest, gin.H{}, "查询失败！")
		return
	}

	for _, stu := range passList {
		sesr.Username = stu
		sesr.Status = "pass"
		sesrl = append(sesrl, sesr)
	}

	for _, stu := range failList {
		sesr.Username = stu
		sesr.Status = "incomplete"
		sesrl = append(sesrl, sesr)
	}

	error.Response(c, error.OK, gin.H{"student": sesrl}, "查询成功！")
}

func ExamList(c *gin.Context, uid int, page int, number int) {
	var eu entity.User
	result := config.MYSQLDB.Table("users").Where("id = ?", uid).First(&eu)
	if result.Error != nil {
		error.Response(c, error.BadRequest, gin.H{}, "查询失败！")
		return
	}

	var eel = make([]entity.Exam, 0)
	var sel = make([]schema.ExamSummary, 0)
	if eu.Authority == "admin" {
		result := config.MYSQLDB.Table("exams").Find(&eel).Limit(number).Offset((page - 1) * number)
		if result.Error != nil {
			error.Response(c, error.BadRequest, gin.H{}, "查询失败！")
			return
		}

		if result.RowsAffected == 0 {
			error.Response(c, error.OK, gin.H{}, "暂无实验！")
			return
		}

		var se schema.ExamSummary
		for _, ee := range eel {
			se.ID = ee.ID
			se.Name = ee.Name
			config.MYSQLDB.Table("question_exams").Where("exam_id = ?", ee.ID).Count(&se.QuestionNum)
			se.StudentNum = int64(len(utils.StringToArr(ee.Student)))
			sel = append(sel, se)
		}

		error.Response(c, error.OK, gin.H{"exam_list": sel}, "查询成功！")
	} else {
		result := config.MYSQLDB.Table("exams").Where("student LIKE ?", eu.Username).Find(&eel).Limit(number).Offset((page - 1) * number)
		if result.Error != nil {
			error.Response(c, error.BadRequest, gin.H{}, "查询失败！")
			return
		}

		if result.RowsAffected == 0 {
			error.Response(c, error.OK, gin.H{}, "暂无与您有关的实验！")
			return
		}

		var se schema.ExamSummary
		for _, ee := range eel {
			se.ID = ee.ID
			se.Name = ee.Name
			config.MYSQLDB.Table("question_exams").Where("exam_id = ?", ee.ID).Count(&se.QuestionNum)
			se.StudentNum = int64(len(utils.StringToArr(ee.Student)))
			sel = append(sel, se)
		}

		error.Response(c, error.OK, gin.H{"exam_list": sel}, "查询成功！")
	}
}
