package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string `gorm:"type:varchar(255);not null" json:"username"`
	Password  string `gorm:"type:varchar(255);not null" json:"password"`
	Email     string `gorm:"type:varchar(255);unique;not null" json:"email"`
	Authority string `gorm:"type:enum('admin', 'stu');default:'stu'" json:"authority"`
	Level     uint   `gorm:"default:0" json:"level"`
}
