package service

import (
	"github.com/gin-gonic/gin"
	"sylu-oj-gin/internal/app/config"
	"sylu-oj-gin/internal/app/entity"
	"sylu-oj-gin/internal/app/schema"
	"sylu-oj-gin/pkg/error"
	"sylu-oj-gin/pkg/logger"
	"sylu-oj-gin/pkg/utils"
	"time"
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

func Login(c *gin.Context, sul schema.UserLogin) {
	var eu entity.User
	result := config.MYSQLDB.Table("users").Where("username = ?", sul.Username).Or("email = ?", sul.Username).First(&eu)
	if result.Error != nil {
		error.Response(c, error.BadRequest, gin.H{}, "用户未注册！")
		return
	}

	if eu.Password != sul.Password {
		error.Response(c, error.BadRequest, gin.H{}, "密码错误！")
		return
	}

	token, err := utils.ReleaseToken(eu)
	if err != nil {
		logger.Error("Token 发送失败！", err)
		error.Response(c, error.BadRequest, gin.H{}, "token 发送失败！")
		return
	}

	error.Response(c, error.OK, gin.H{"token": token, "username": eu.Username}, "登录成功！")
}

func UserInfo(c *gin.Context, id int) {
	var eu entity.User
	if result := config.MYSQLDB.Table("users").Where("id = ?", id).First(&eu); result.Error != nil {
		error.Response(c, error.BadRequest, gin.H{}, "获取用户信息失败！")
		return
	}

	var sui schema.UserInfo
	if eu.Authority == "admin" {
		sui.Identity = "admin"
		sui.Menu = append(sui.Menu, schema.ResponseMenu{
			ID:    "3",
			Title: "题目管理",
			Path:  "/console/problem",
		})
		sui.Menu = append(sui.Menu, schema.ResponseMenu{
			ID:    "5",
			Title: "添加题目",
			Path:  "/console/addproblem",
		})
	} else {
		sui.Identity = "stu"
	}

	sui.Info.Username = eu.Username
	sui.Info.StartTime = eu.CreatedAt.Format(time.RFC3339)
	config.MYSQLDB.Table("answers").Select("COUNT(DISTINCT question_id)").Where("user_id = ?", eu.ID).Scan(&sui.Info.Submit)
	config.MYSQLDB.Table("answers").Select("COUNT(DISTINCT question_id)").Where("user_id = ? AND status = ?", eu.ID, "Accepted").Scan(&sui.Info.Accept)

	config.MYSQLDB.Table("users").Select("username AS name").Where("authority = ?", "stu").Find(&sui.Students)

	if eu.Authority != "admin" {
		sui.Students = nil
	}

	error.Response(c, error.OK, gin.H{"identity": sui.Identity, "menu": sui.Menu, "info": sui.Info, "students": sui.Students}, "欢迎！")
}
