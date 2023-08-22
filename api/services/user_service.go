package services

import (
	"github.com/BetterToPractice/go-echo-setup/lib"
	"github.com/BetterToPractice/go-echo-setup/models"
)

type UserService struct {
	logger lib.Logger
	db     lib.Database
}

func NewUserService(logger lib.Logger, db lib.Database) UserService {
	return UserService{
		logger: logger,
		db:     db,
	}
}

func (c *UserService) Query() (models.UserIndexResult, error) {
	var result models.UserIndexResult
	err := c.db.ORM.Model(&models.User{}).Find(&result.Results).Error
	return result, err
}
