package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"regexp"
	"sylu-oj-gin/internal/app/config"
)

func Cors() gin.HandlerFunc {
	c := cors.DefaultConfig()
	c.AllowMethods = []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"}
	c.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "token"}
	if gin.Mode() == gin.ReleaseMode {
		// 生产环境配置跨域域名，否则 403
		c.AllowOrigins = []string{config.CONFIG.Http.Host}
	} else {
		// 测试环境允许域名的请求
		c.AllowOrigins = []string{config.CONFIG.Http.Host}
		// 测试环境下允许 localhost 的请求
		c.AllowOriginFunc = func(origin string) bool {
			if regexp.MustCompile(`^http://127\.0\.0\.1:\d+$`).MatchString(origin) {
				return true
			}
			if regexp.MustCompile(`^http://localhost:\d+$`).MatchString(origin) {
				return true
			}
			return false
		}
	}
	return cors.New(c)
}
