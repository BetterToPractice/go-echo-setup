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

func (s *UserService) Query(params *models.UserQueryParams) (*models.UserPaginationResult, error) {
	return s.userRepository.Query(params)
}

func (s *UserService) GetByUsername(username string) (*models.User, error) {
	return s.userRepository.GetByUsername(username)
}

func (s *UserService) Delete(username string) error {
	user, err := s.userRepository.GetByUsername(username)
	if err != nil {
		return err
	}

	if err := s.profileRepository.DeleteByUserID(strconv.Itoa(int(user.ID))); err != nil {
		return err
	}

	return s.userRepository.Delete(username)
}
