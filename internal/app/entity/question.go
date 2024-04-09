package entity

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	Title          string  `gorm:"type:varchar(255);not null" json:"title"`
	Content        string  `gorm:"not null" json:"content"`
	Tag            string  `gorm:"omitempty" json:"tag"`
	Degree         uint    `gorm:"not null" json:"degree"`
	PassingRate    float64 `json:"passing_rate"`
	OwnerID        int     `json:"owner_id"`
	InputTest      string  `json:"input_test"`
	ExpectedOutput string  `json:"expected_output"`
}

var TagArr []string
