package controllers

import (
	"github.com/BetterToPractice/go-echo-setup/api/services"
	"github.com/BetterToPractice/go-echo-setup/models/dto"
	"github.com/BetterToPractice/go-echo-setup/pkg/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) AuthController {
	return AuthController{
		authService: authService,
	}
}

// Register godoc
//
//	@Summary		Register a new User
//	@Description	register a new user
//	@Tags			auth
//	@Accept			application/json
//	@Produce		application/json
//	@Router			/register [post]
//	@Success		200  {object}  response.Response{data=dto.RegisterResponse}  "ok"
func (c AuthController) Register(ctx echo.Context) error {
	register := new(dto.Register)
	if err := ctx.Bind(register); err != nil {
		return response.Response{
			Code:    http.StatusBadRequest,
			Message: err,
		}.JSON(ctx)
	}

	_, err := c.authService.Register(register)
	if err != nil {
		return response.Response{
			Code:    http.StatusBadRequest,
			Message: err,
		}.JSON(ctx)
	}

	return response.Response{
		Message: dto.RegisterResponse{
			Username: register.Username,
			Email:    register.Email,
		},
	}.JSON(ctx)
}

// Login godoc
//
//	@Summary		Login a User
//	@Description	Login as user to application
//	@Tags			auth
//	@Accept			application/json
//	@Produce		application/json
//	@Router			/login [post]
//	@Success		200  {object}  response.Response{data=dto.LoginResponse}  "ok"
func (c AuthController) Login(ctx echo.Context) error {
	login := new(dto.Login)
	if err := ctx.Bind(login); err != nil {
		return response.Response{
			Code:    http.StatusBadRequest,
			Message: err,
		}.JSON(ctx)
	}

	token, err := c.authService.Login(login)
	if err != nil {
		return response.Response{
			Code:    http.StatusUnauthorized,
			Message: err,
		}.JSON(ctx)
	}

	return response.Response{
		Code:    http.StatusOK,
		Message: token,
	}.JSON(ctx)
}
