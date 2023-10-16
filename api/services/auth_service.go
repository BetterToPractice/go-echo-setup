package services

import (
	"errors"
	"fmt"
	"github.com/BetterToPractice/go-echo-setup/api/mails"
	"github.com/BetterToPractice/go-echo-setup/api/repositories"
	"github.com/BetterToPractice/go-echo-setup/constants"
	"github.com/BetterToPractice/go-echo-setup/lib"
	"github.com/BetterToPractice/go-echo-setup/models"
	"github.com/BetterToPractice/go-echo-setup/models/dto"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"time"
)

type AuthService struct {
	authRepository repositories.AuthRepository
	authMail       mails.AuthMail
	userService    UserService
	config         lib.Config
	opts           *options
	db             lib.Database
}

type options struct {
	issuer       string
	signedMethod jwt.SigningMethod
	signedKey    interface{}
	keyfunc      jwt.Keyfunc
	expired      int
}

func NewAuthService(authRepository repositories.AuthRepository, userService UserService, config lib.Config, db lib.Database, authMail mails.AuthMail) AuthService {
	signingKey := fmt.Sprintf("jwt:%s", config.Secret)
	opts := &options{
		issuer:       config.Name,
		expired:      config.Auth.TokenExpired,
		signedMethod: jwt.SigningMethodES512,
		signedKey:    []byte(signingKey),
		keyfunc: func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("invalid token")
			}
			return []byte(signingKey), nil
		},
	}
	return AuthService{
		authRepository: authRepository,
		userService:    userService,
		authMail:       authMail,
		opts:           opts,
		config:         config,
		db:             db,
	}
}

func (s AuthService) GenerateToken(user *models.User) (string, error) {
	now := time.Now()
	claims := &dto.JWTClaims{
		ID:       user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(s.config.Auth.TokenExpired) * time.Second)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString(s.opts.signedKey)
}

func (s AuthService) ParseToken(tokenStr string) (*dto.JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &dto.JWTClaims{}, s.opts.keyfunc)
	if err != nil {
		return nil, err
	}
	if token != nil {
		if claims, ok := token.Claims.(*dto.JWTClaims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, errors.New("invalid token")
}

func (s AuthService) Register(register *dto.Register) (*models.User, error) {
	user := &models.User{
		Username: register.Username,
		Password: models.HashPassword(register.Password),
		Email:    register.Email,
	}
	if err := s.db.ORM.Create(&user).Error; err != nil {
		return nil, err
	}

	s.authMail.Register(user)

	return user, nil
}

func (s AuthService) Login(login *dto.Login) (*dto.LoginResponse, error) {
	user, err := s.userService.Verify(login.Username, login.Password)
	if err != nil {
		return nil, err
	}

	access, err := s.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{Access: access}, nil
}

func (s AuthService) Authenticate(ctx echo.Context) (*models.User, error) {
	jwtClaims, ok := ctx.Get(constants.CurrentUser).(*dto.JWTClaims)

	if !ok {
		return nil, errors.New("unauthorized")
	}

	if user, err := s.userService.GetByUsername(jwtClaims.Username); err != nil {
		return nil, err
	} else {
		return user, nil
	}
}
