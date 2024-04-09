package entity

import "gorm.io/gorm"

type Answer struct {
	gorm.Model
	QuestionID uint   `json:"question_id"`
	Answer     string `json:"answer"`
	UserID     int    `json:"user_id"`
}
