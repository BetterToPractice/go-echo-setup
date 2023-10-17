package services

import (
	"errors"
	"github.com/BetterToPractice/go-echo-setup/api/dto"
	"github.com/BetterToPractice/go-echo-setup/api/repositories"
	appError "github.com/BetterToPractice/go-echo-setup/errors"
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

func (s *UserService) Query(params *dto.UserQueryParam) (*dto.UserPaginationResponse, error) {
	return s.userRepository.Query(params)
}

func (s *UserService) GetByUsername(username string) (*models.User, error) {
	user, err := s.userRepository.GetByUsername(username)
	if errors.Is(err, appError.DatabaseInternalError) {
		return nil, appError.UserNotFound
	}
	return user, err
}

func (s UserService) Register(params *dto.RegisterRequest) (*models.User, error) {
	user := &models.User{
		Username: params.Username,
		Password: params.Password,
		Email:    params.Email,
	}
	if err := s.userRepository.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) Delete(user *models.User) error {
	if err := s.profileRepository.DeleteByUserID(strconv.Itoa(int(user.ID))); err != nil {
		return errors.Join(appError.DatabaseInternalError, err)
	}
	return s.userRepository.Delete(user)
}

func (s UserService) Verify(username string, password string) (*models.User, error) {
	user, err := s.userRepository.GetByUsername(username)
	if err != nil || user.Password != models.HashPassword(password) {
		return nil, appError.UsernameOrPasswordNotMatch
	}
	return user, nil
}
