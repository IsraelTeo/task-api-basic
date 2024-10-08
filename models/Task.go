package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model

	Title       string `gorm:"not null;unique" json:"title"`
	Description string
	Done        bool `gorm:"default:false"`
	UserID      uint
}
