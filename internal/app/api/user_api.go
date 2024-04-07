package api

import (
	"github.com/gin-gonic/gin"
	"sylu-oj-gin/internal/app/schema"
	"sylu-oj-gin/internal/app/service"
	"sylu-oj-gin/pkg/error"
)

// @Summary 注册接口
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

// @Summary 登录接口
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
