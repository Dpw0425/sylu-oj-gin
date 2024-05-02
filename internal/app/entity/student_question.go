package entity

import "gorm.io/gorm"

type StudentQuestion struct {
	gorm.Model
	ExamID     int    `json:"exam_id" gorm:"exam_id"`
	Username   string `json:"username" gorm:"username"`
	QuestionID int    `json:"question_id" gorm:"question_id"`
	Status     string `json:"status" gorm:"status;default:'incomplete'"`
}
