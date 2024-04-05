package api

import (
	"github.com/gin-gonic/gin"
	"sylu-oj-gin/internal/app/service"
	"sylu-oj-gin/pkg/error"
)

// @Summary 测试接口
// @Description 测试连接
// @Accept       json
// @Produce      json
// @Success      200  {object}  error.ResponseNormal
// @Router       /test/ping [get]
func Ping(c *gin.Context) {
	error.Response(c, error.OK, gin.H{}, "pong")
}

// @Summary 邮箱验证接口
// @Description 发送邮箱验证码
// @Param email query string true "user email"
// @Accept       json
// @Produce      json
// @Success      200  {object}  error.ResponseNormal
// @Router       /utils/email_verify [get]
func SendEmailCode(c *gin.Context) {
	email := c.Query("email")

	service.SendEmailCode(c, email)
}
