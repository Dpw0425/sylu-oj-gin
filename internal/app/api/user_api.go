package api

import (
	"github.com/gin-gonic/gin"
	"sylu-oj-gin/internal/app/schema"
	"sylu-oj-gin/internal/app/service"
	"sylu-oj-gin/pkg/error"
)

// Register @Summary 注册接口
// @Description 用户注册
// @Param request body schema.UserRegister true "user message"
// @Accept       json
// @Produce      json
// @Success      200  {object}  error.ResponseNormal
// @Router       /user/register [post]
func Register(c *gin.Context) {
	var sur schema.UserRegister
	if err := c.ShouldBindJSON(&sur); err != nil {
		error.Response(c, error.BadRequest, gin.H{}, "输入内容有误！")
		return
	}

	service.Register(c, sur)
}

// Login @Summary 登录接口
// @Description 用户登录
// @Param request body schema.UserLogin true "user message"
// @Accept       json
// @Produce      json
// @Success      200  {object}  error.ResponseNormal
// @Router       /user/login [post]
func Login(c *gin.Context) {
	var sul schema.UserLogin
	if err := c.ShouldBindJSON(&sul); err != nil {
		error.Response(c, error.BadRequest, gin.H{}, "输入内容有误！")
		return
	}

	service.Login(c, sul)
}

// Logout @Summary 登出接口
// @Description 退出登录
// @Accept       json
// @Produce      json
// @Success      200  {object}  error.ResponseNormal
// @Router       /user/logout [get]
func Logout(c *gin.Context) {
	c.Set("UserID", "")
	c.Header("token", "")

	error.Response(c, error.BadRequest, gin.H{}, "退出成功！")
}

func UserInfo(c *gin.Context) {
	id, _ := c.Get("UserID")
	uid := id.(int)

	service.UserInfo(c, uid)
}
