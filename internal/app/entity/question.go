package entity

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	Title       string  `gorm:"type:varchar(255);not null" json:"title"`
	Content     string  `gorm:"not null" json:"content"`
	Tag         string  `gorm:"omitempty" json:"tag"`
	Degree      uint    `gorm:"not null" json:"degree"`
	PassingRate float64 `json:"passing_rate"`
	OwnerID     int     `json:"owner_id"`
}

type TestData struct {
	gorm.Model
	QID    uint   `json:"qid" gorm:"qid"`
	Input  string `json:"input" gorm:"input"`
	Output string `json:"output" gorm:"output"`
}

var TagArr []string
