package service

import (
	"github.com/gin-gonic/gin"
	"sylu-oj-gin/internal/app/config"
	"sylu-oj-gin/internal/app/entity"
	"sylu-oj-gin/internal/app/schema"
	"sylu-oj-gin/pkg/error"
	"sylu-oj-gin/pkg/utils"
)

func AddQuestion(c *gin.Context, saq schema.AddQuestion, uid uint) {
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
