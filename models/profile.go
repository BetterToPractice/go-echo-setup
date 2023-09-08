package models

import "gorm.io/gorm"

type Profile struct {
	gorm.Model

	PhoneNumber string `gorm:"column:phone_number;null;"`
	Gender      string `gorm:"column:gender;not null;"`

	UserID uint `gorm:"column:user_id;not null;"`
}
