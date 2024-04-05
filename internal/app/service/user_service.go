package service

import (
	"github.com/gin-gonic/gin"
	"sylu-oj-gin/internal/app/config"
	"sylu-oj-gin/internal/app/entity"
	"sylu-oj-gin/internal/app/schema"
	"sylu-oj-gin/pkg/error"
	"sylu-oj-gin/pkg/utils"
)

func Register(c *gin.Context, sur schema.UserRegister) {
	var eu entity.User
	result := config.MYSQLDB.Table("users").Where("username = ?", sur.Username).First(&eu)
	if result.RowsAffected != 0 {
		error.Response(c, error.BadRequest, gin.H{}, "用户名已存在！")
		return
	}

	result1 := config.MYSQLDB.Table("users").Where("email = ?", sur.Email).First(&eu)
	if result1.RowsAffected != 0 {
		error.Response(c, error.BadRequest, gin.H{}, "邮箱已被注册！")
		return
	}

	if !utils.VerifyCode(sur.Email, sur.VerifyCode) {
		error.Response(c, error.BadRequest, gin.H{}, "验证码错误！")
		return
	}

	var eu1 entity.User
	eu1.Username = sur.Username
	eu1.Password = sur.Password
	eu1.Email = sur.Email
	result2 := config.MYSQLDB.Table("users").Create(&eu1)
	if result2.Error != nil {
		error.Response(c, error.BadRequest, gin.H{}, "注册失败！")
		return
	}

	error.Response(c, error.OK, gin.H{}, "注册成功！")
}
