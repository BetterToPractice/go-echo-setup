package services

import (
	"github.com/BetterToPractice/go-echo-setup/api/repositories"
	"github.com/BetterToPractice/go-echo-setup/lib"
	"github.com/BetterToPractice/go-echo-setup/models"
)

type UserService struct {
	logger         lib.Logger
	userRepository repositories.UserRepository
}

func NewUserService(logger lib.Logger, userRepository repositories.UserRepository) UserService {
	return UserService{
		logger:         logger,
		userRepository: userRepository,
	}
}

func (c *UserService) Query(params *models.UserQueryParams) (*models.UserPaginationResult, error) {
	return c.userRepository.Query(params)
}

func (c *UserService) GetByUsername(username string) (*models.User, error) {
	return c.userRepository.GetByUsername(username)
}
