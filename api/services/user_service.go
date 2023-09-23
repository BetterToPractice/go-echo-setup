package services

import (
	"github.com/BetterToPractice/go-echo-setup/api/repositories"
	"github.com/BetterToPractice/go-echo-setup/lib"
	"github.com/BetterToPractice/go-echo-setup/models"
	"strconv"
)

type UserService struct {
	logger            lib.Logger
	userRepository    repositories.UserRepository
	profileRepository repositories.ProfileRepository
}

func NewUserService(logger lib.Logger, userRepository repositories.UserRepository, profileRepository repositories.ProfileRepository) UserService {
	return UserService{
		logger:            logger,
		userRepository:    userRepository,
		profileRepository: profileRepository,
	}
}

func (c *UserService) Query(params *models.UserQueryParams) (*models.UserPaginationResult, error) {
	return c.userRepository.Query(params)
}

func (c *UserService) GetByUsername(username string) (*models.User, error) {
	return c.userRepository.GetByUsername(username)
}

func (c *UserService) Delete(username string) error {
	user, err := c.userRepository.GetByUsername(username)
	if err != nil {
		return err
	}

	if err := c.profileRepository.DeleteByUserID(strconv.Itoa(int(user.ID))); err != nil {
		return err
	}

	return c.userRepository.Delete(username)
}
