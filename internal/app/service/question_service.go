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
	var eq entity.Question
	eq.Title = saq.Title
	eq.Content = saq.Content
	eq.Tag = utils.ArrToString(saq.Tag)
	eq.Degree = saq.Degree
	eq.OwnerID = uid
	result := config.MYSQLDB.Table("questions").Create(&eq)
	if result.Error != nil {
		error.Response(c, error.BadRequest, gin.H{}, "发布失败！")
		return
	}

	error.Response(c, error.OK, gin.H{}, "发布成功！")
}

func QuestionList(c *gin.Context, page int, number int, searchTitle string, searchTag string, searchDegree int, order int) {
	var sqml = make([]schema.QuestionMsg, 0)
	var eql = make([]entity.Question, 0)
	var sql = "SELECT * FROM questions "
	var flag = 0
	var result *gorm.DB
	if searchTitle != "" {
		sql = sql + "WHERE title LIKE '" + searchTitle + "' "
		flag++
	}
	if searchTag != "" {
		if flag != 0 {
			sql = sql + "AND tag LIKE '" + searchTag + "' "
		} else {
			sql = sql + "WHERE tag LIKE '" + searchTag + "' "
			flag++
		}
	}
	if searchDegree != 0 {
		if flag != 0 {
			result = config.MYSQLDB.Limit(number).Offset((page-1)*number).Raw(sql+"AND degree = ?", searchDegree).Find(&eql)
		} else {
			result = config.MYSQLDB.Limit(number).Offset((page-1)*number).Raw(sql+"WHERE degree = ?", searchDegree).Find(&eql)
		}
		if result.Error != nil {
			error.Response(c, error.BadRequest, gin.H{}, "查询失败！")
			return
		}
		if result.RowsAffected == 0 {
			error.Response(c, error.BadRequest, gin.H{}, "暂无符合条件的记录！")
			return
		}
		flag = 10
	}

	if order != 0 {
		if order == 1 {
			sql = sql + "ORDER BY degree ASC"
		} else {
			sql = sql + "ORDER BY degree DESC"
		}
	}

	sql = sql + ";"

	if flag != 10 {
		result = config.MYSQLDB.Limit(number).Offset((page - 1) * number).Raw(sql).Find(&eql)
		if result.Error != nil {
			error.Response(c, error.BadRequest, gin.H{}, "查询失败！")
			return
		}
		if result.RowsAffected == 0 {
			error.Response(c, error.BadRequest, gin.H{}, "暂无符合条件的记录！")
			return
		}
	}

	var sqm schema.QuestionMsg
	for _, eq := range eql {
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

func CommitAnswer(c *gin.Context, sa schema.Answer) {

}
