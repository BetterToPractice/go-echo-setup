package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"column:username;size:64;not null;index;" json:"username" validate:"required"`
	Password string `gorm:"column:password;not null;" json:"password" validate:"required"`
	Email    string `gorm:"column:email;not null;" json:"email" validate:"required"`
}

type UserIndexResult struct {
	Results []User `json:"list"`
}
