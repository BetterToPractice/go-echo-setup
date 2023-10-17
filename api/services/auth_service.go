package services

import (
	"errors"
	"fmt"
	"github.com/BetterToPractice/go-echo-setup/api/dto"
	"github.com/BetterToPractice/go-echo-setup/api/mails"
	"github.com/BetterToPractice/go-echo-setup/api/repositories"
	"github.com/BetterToPractice/go-echo-setup/constants"
	appError "github.com/BetterToPractice/go-echo-setup/errors"
	"github.com/BetterToPractice/go-echo-setup/lib"
	"github.com/BetterToPractice/go-echo-setup/models"
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

func (s AuthService) GenerateToken(user *models.User) (*dto.JWTResponse, error) {
	now := time.Now()
	resp := &dto.JWTResponse{}

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
	access, err := token.SignedString(s.opts.signedKey)
	if err != nil {
		return nil, err
	}

	resp.Serializer(access)
	return resp, nil
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

func (s AuthService) Register(register *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	user, err := s.userService.Register(register)
	if err != nil {
		return nil, err
	}

	s.authMail.Register(user)

	resp := &dto.RegisterResponse{}
	resp.Serializer(user)

	return resp, nil
}

func (s AuthService) Login(login *dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := s.userService.Verify(login.Username, login.Password)
	if err != nil {
		return nil, err
	}

	access, err := s.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	resp := &dto.LoginResponse{}
	resp.Serializer(access.Access)

	return resp, nil
}

func (s AuthService) Authenticate(ctx echo.Context) (*models.User, error) {
	jwtClaims, ok := ctx.Get(constants.CurrentUser).(*dto.JWTClaims)
	if !ok {
		return nil, appError.Unauthorized
	}

	if user, err := s.userService.GetByUsername(jwtClaims.Username); err != nil {
		return nil, appError.Unauthorized
	} else {
		return user, nil
	}
}
