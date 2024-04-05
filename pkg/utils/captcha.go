package utils

import (
	"sylu-oj-gin/internal/app/config"
	"sylu-oj-gin/pkg/logger"
	"time"
)

func GenerateCode(email string) bool {
	// 生成验证码
	code := Random(6)
	// 存储至 Redis
	err := config.REDISDB.Set(email, code, time.Minute*time.Duration(config.CONFIG.Email.Expires)).Err()
	if err != nil {
		logger.Error("redis 验证码存储失败.", err)
		return false
	}
	return SendEmail(email, code)
}

func VerifyCode(email string, code string) bool {
	// 获取 Redis 中的验证码
	redisCode, err := config.REDISDB.Get(email).Result()
	if err != nil {
		logger.Error("redis 验证码获取失败.", err)
		return false
	}
	// 比较验证码
	if code != redisCode {
		logger.Error("redis 验证码错误.", err)
		return false
	}
	return true
}
