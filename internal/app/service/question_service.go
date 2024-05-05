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

func AddQuestion(c *gin.Context, saq schema.AddQuestion, uid int) {
	newDB := config.MYSQLDB.Session(&gorm.Session{NewDB: true})
	tx := newDB.Begin()

	var eq entity.Question
	eq.Title = saq.Title
	eq.Content = saq.Content
	eq.Tag = utils.ArrToString(saq.Tag)
	eq.Degree = saq.Degree
	eq.OwnerID = uid
	result := tx.Table("questions").Create(&eq)
	if result.Error != nil {
		error.Response(c, error.BadRequest, gin.H{}, "发布失败！")
		tx.Rollback()
		return
	}

	var etd entity.TestData
	for _, value := range saq.IO {
		etd.QID = eq.ID
		etd.Input = value.Input
		etd.Output = value.Output
		result := tx.Table("test_data").Create(&etd)
		if result.Error != nil {
			error.Response(c, error.BadRequest, gin.H{}, "发布失败！")
			tx.Rollback()
			return
		}
	}

	error.Response(c, error.OK, gin.H{}, "发布成功！")
	tx.Commit()
}

func QuestionList(c *gin.Context, page int, number int, searchTitle string, searchTag string, searchDegree int, order int) {
	var sqml = make([]schema.QuestionMsg, 0)
	var eql = make([]entity.Question, 0)
	var result *gorm.DB

	query := config.MYSQLDB.Model(&entity.Question{})

	if searchTitle != "" {
		query = query.Where("title LIKE ?", "%"+searchTitle+"%")
	}

	if searchTag != "" {
		query = query.Where("tag LIKE ?", "%"+searchTag+"%")
	}

	if searchDegree != 0 {
		query = query.Where("degree = ?", searchDegree)
	}

	if order != 0 {
		if order == 1 {
			query = query.Order("degree ASC")
		} else {
			query = query.Order("degree DESC")
		}
	}

	result = query.Limit(number).Offset((page - 1) * number).Find(&eql)

	if result.Error != nil {
		error.Response(c, error.BadRequest, gin.H{}, "查询失败！")
		return
	}

	if result.RowsAffected == 0 {
		error.Response(c, error.BadRequest, gin.H{}, "暂无符合条件的记录！")
		return
	}

	var sqm schema.QuestionMsg
	for _, eq := range eql {
		sqm.ID = eq.ID
		sqm.Title = eq.Title
		sqm.Tag = utils.StringToArr(eq.Tag)
		sqm.Degree = eq.Degree
		sqml = append(sqml, sqm)
	}

	error.Response(c, error.OK, gin.H{"question_list": sqml}, "查询成功！")
}

func GetQuestionMsg(c *gin.Context, id int) {
	var sq schema.Question
	var eq entity.Question
	result := config.MYSQLDB.Table("questions").Where("id = ?", id).First(&eq)
	if result.Error != nil {
		error.Response(c, error.BadRequest, gin.H{}, "查询失败！")
		return
	}

	sq.ID = int(eq.ID)
	sq.Title = eq.Title
	sq.Content = eq.Content
	sq.Tag = utils.StringToArr(eq.Tag)
	sq.Degree = int(eq.Degree)

	error.Response(c, error.OK, gin.H{"question_msg": sq}, "查询成功！")
}

func CommitAnswer(c *gin.Context, sa schema.Answer, uid int) {
	newDB := config.MYSQLDB.Session(&gorm.Session{NewDB: true})
	tx := newDB.Begin()

	var eq entity.Question
	result := tx.Table("questions").Where("id = ?", sa.ID).First(&eq)
	if result.Error != nil {
		error.Response(c, error.BadRequest, gin.H{}, "提交失败！")
		tx.Rollback()
		return
	}

	var ea entity.Answer
	ea.UserID = uid
	ea.Answer = sa.Answer
	ea.QuestionID = eq.ID
	result1 := tx.Table("answers").Create(&ea)
	if result1.Error != nil {
		error.Response(c, error.BadRequest, gin.H{}, "提交失败！")
		tx.Rollback()
		return
	}
	// TODO: Using a thread pool to start the question machine
	var result2 string
	var etdl = make([]entity.TestData, 0)
	result3 := tx.Table("test_data").Where("q_id = ?", sa.ID).Find(&etdl)
	if result3.Error != nil {
		error.Response(c, error.BadRequest, gin.H{}, "提交失败！")
		tx.Rollback()
		return
	}

	tx.Commit()

	for _, value := range etdl {
		if flag := utils.Judge(sa.Answer, value.Input, value.Output); flag != "Accepted" {
			result2 = flag
			break
		}
	}

	if result2 == "" {
		result2 = "Accepted"
	}

	if result2 == "Accepted" {
		config.MYSQLDB.Table("student_questions").Where("question_id = ? AND username IN (SELECT username FROM users WHERE users.id = ?)", sa.ID, uid).Update("status", "pass")
	} else {
		config.MYSQLDB.Table("student_questions").Where("question_id = ? AND username IN (SELECT username FROM users WHERE users.id = ?)", sa.ID, uid).Update("status", "fail")
	}

	error.Response(c, error.OK, gin.H{"result": result2}, "提交成功！")
}

func DelQuestion(c *gin.Context, qid int, uid int) {
	var eq entity.Question
	result := config.MYSQLDB.Table("questions").Where("id = ?", qid).First(&eq)
	if result.Error != nil {
		error.Response(c, error.BadRequest, gin.H{}, "获取题目失败！")
		return
	}

	if eq.OwnerID != uid {
		error.Response(c, error.BadRequest, gin.H{}, "您无权删除！")
		return
	}

	tx := config.MYSQLDB.Begin()
	result1 := tx.Table("questions").Where("id = ?", qid).Delete(&eq)
	if result1.Error != nil {
		error.Response(c, error.BadRequest, gin.H{}, "删除失败！")
		tx.Rollback()
		return
	}
	result2 := tx.Table("answers").Where("question_id = ?", qid).Delete(&entity.Answer{})
	if result2.Error != nil {
		error.Response(c, error.BadRequest, gin.H{}, "删除失败！")
		tx.Rollback()
		return
	}

	tx.Commit()
	error.Response(c, error.OK, gin.H{}, "删除成功！")
}
