package api

import (
	"github.com/gin-gonic/gin"
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
