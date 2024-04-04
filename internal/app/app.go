package app

import (
	"github.com/gin-gonic/gin"
	"sylu-oj-gin/internal/app/config"
	"sylu-oj-gin/internal/app/db"
	"sylu-oj-gin/pkg/logger"
)

func Init() *gin.Engine {
	logger.InitZap()
	config.InitConfig()
	db.InitMysql()
	db.InitRedis()

	r := InitGinEngine()

	return r
}
