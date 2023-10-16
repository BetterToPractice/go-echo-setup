package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model

	Title  string `gorm:"column:title;size:200;not null;"`
	Body   string `gorm:"column:body;not null;"`
	UserID uint   `gorm:"column:user_id;not null;"`
	User   User   `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Posts []Post
