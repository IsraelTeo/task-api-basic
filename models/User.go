package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	FirstName string `gorm:"not null" json:"name"`
	LastName  string `gorm:"not null" json:"last_name"`
	Email     string `gorm:"not null;unique" json: "email"`
	Tasks     []Task `json:"tasks"`
}
