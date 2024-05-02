package entity

import "gorm.io/gorm"

type QuestionExam struct {
	gorm.Model
	QuestionID int `json:"question_id" gorm:"question_id"`
	ExamID     int `json:"exam_id" gorm:"exam_id"`
}
