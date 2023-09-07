package models

import (
	"github.com/BetterToPractice/go-echo-setup/models/dto"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"column:username;size:64;not null;index;" json:"username" validate:"required"`
	Password string `gorm:"column:password;not null;" json:"password" validate:"required"`
	Email    string `gorm:"column:email;not null;" json:"email" validate:"required"`

	posts []Post `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Users = []User

type UserIndexResult struct {
	Results []User `json:"list"`
}

type UserQueryParams struct {
	dto.PaginationParam
}

type UserPaginationResult struct {
	List       Users           `json:"list"`
	Pagination *dto.Pagination `json:"pagination"`
}
