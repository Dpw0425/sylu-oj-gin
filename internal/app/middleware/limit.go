package middleware

import (
	"github.com/gin-gonic/gin"
	"sylu-oj-gin/internal/app/config"
	"sylu-oj-gin/pkg/error"
	"time"
)

func LimitRoute(limit int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.ClientIP() + c.FullPath()
		// 读取请求次数
		value, err := config.REDISDB.Get(key).Int64()
		if err != nil {
			// 初始化请求次数
			config.REDISDB.Set(key, 1, time.Hour)
			c.Next()
		} else if value < limit {
			// 请求次数 +1
			config.REDISDB.Incr(key)
			c.Next()
		} else {
			// 请求次数超过限制
			error.Response(c, error.TooManyRequests, gin.H{}, "请求过于频繁。")
			c.Abort()
		}
	}
}

// 按时间单路由限流
func LimitAverageRoute(limit int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.ClientIP() + c.FullPath()
		// 读取值并验证
		if err := config.REDISDB.Get(key).Err(); err != nil {
			// 初始化请求次数
			config.REDISDB.Set(key, 1, time.Hour/time.Duration(limit))
			c.Next()
		} else {
			error.ErrResponse(c, error.ErrTooManyRequests, error.TooManyRequests, "请求过于频繁，请一小时后再重试。", err)
			c.Abort()
		}
	}
}
