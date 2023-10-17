package models

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/BetterToPractice/go-echo-setup/models/dto"
	"gorm.io/gorm"
	"unsafe"
)

type User struct {
	gorm.Model
	Username string `gorm:"column:username;size:64;not null;index;" json:"username" validate:"required"`
	Password string `gorm:"column:password;not null;" json:"-" validate:"required"`
	Email    string `gorm:"column:email;not null;" json:"email" validate:"required"`

	Profile Profile `gorm:"foreignKey:UserID;references:ID;"`
	Posts   []Post  `gorm:"foreignKey:UserID;references:ID;"`
}

type Users = []User

func (m *User) BeforeCreate(_ *gorm.DB) (err error) {
	m.Password = HashPassword(m.Password)
	return
}

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

func HashPassword(password string) string {
	sum := sha256.Sum256(*(*[]byte)(unsafe.Pointer(&password)))
	return hex.EncodeToString(sum[:])
}
