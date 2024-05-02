package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sylu-oj-gin/internal/app/config"
	"sylu-oj-gin/internal/app/entity"
	"sylu-oj-gin/pkg/logger"
)

func InitMysql() {
	m := config.CONFIG.Mysql
	db, err := gorm.Open(mysql.Open(config.DSN(m)), &gorm.Config{})

	dbConf, _ := db.DB()
	dbConf.SetMaxIdleConns(10)
	dbConf.SetMaxOpenConns(100)

	if err != nil {
		logger.Error("连接 Mysql 失败: %v", err)
	} else {
		// TODO: auto migrate
		db.AutoMigrate(&entity.User{})
		db.AutoMigrate(&entity.Question{})
		db.AutoMigrate(&entity.Answer{})
		db.AutoMigrate(&entity.Exam{})
		db.AutoMigrate(&entity.TestData{})
		db.AutoMigrate(&entity.QuestionExam{})
		db.AutoMigrate(&entity.StudentQuestion{})

		config.MYSQLDB = db
		logger.Info("连接 Mysql 成功: %v", config.DSN(m))
	}
}
