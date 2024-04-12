package entity

import "gorm.io/gorm"

type Exam struct {
	gorm.Model
	Name    string `json:"name"`
	Student string `json:"student"`
}
