package service

import (
	"github.com/gin-gonic/gin"
	"sylu-oj-gin/internal/app/config"
	"sylu-oj-gin/internal/app/entity"
	"sylu-oj-gin/internal/app/schema"
	"sylu-oj-gin/pkg/error"
)

func AddExam(c *gin.Context, sae schema.AddExam) {
	var ee entity.Exam
	// TODO: START TRANSACTION
	for _, student := range sae.Student {
		ee.Name = sae.Name
		ee.Student = student
		result := config.MYSQLDB.Table("exams").Create(&ee)
		if result.Error != nil {
			error.Response(c, error.BadRequest, gin.H{}, "创建失败！")
			// TODO: ROLL BACK
			return
		}
	}

	// TODO: COMMIT TRANSACTION
	error.Response(c, error.OK, gin.H{}, "创建成功！")
}
