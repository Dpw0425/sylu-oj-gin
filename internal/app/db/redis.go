package db

import (
	"github.com/go-redis/redis"
	"sylu-oj-gin/internal/app/config"
	"sylu-oj-gin/pkg/logger"
)

func InitRedis() {
	// 获取 redis 配置
	r := config.CONFIG.Redis
	// 建立连接
	client := redis.NewClient(&redis.Options{
		Addr:     r.Addr,
		Password: r.Password,
	})
	// 测试连接
	pong, err := client.Ping().Result()
	if err != nil {
		logger.Error("连接 redis 失败: %v", err)
	} else {
		config.REDISDB = client
		logger.Info("连接 redis 成功: %v", pong)
	}
}
