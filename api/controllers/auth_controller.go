package controllers

import (
	"github.com/BetterToPractice/go-echo-setup/api/dto"
	"github.com/BetterToPractice/go-echo-setup/api/services"
	"github.com/BetterToPractice/go-echo-setup/pkg/response"
	"github.com/labstack/echo/v4"
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
//	@Param 			data body dto.RegisterRequest true "Post"
//	@Router			/register [post]
//	@Success		200  {object}  response.Response{data=dto.RegisterResponse}  "ok"
//	@Failure		400  {object}  response.Response{data=[]response.ValidationErrors}  "bad request"
func (c AuthController) Register(ctx echo.Context) error {
	register := new(dto.RegisterRequest)
	if err := ctx.Bind(register); err != nil {
		return response.BadRequest{Req: dto.RegisterRequest{}, Message: err}.JSON(ctx)
	}

	resp, err := c.authService.Register(register)
	if err != nil {
		return response.BadRequest{Message: err}.JSON(ctx)
	}

	return response.Response{Data: resp}.JSON(ctx)
}

// Login godoc
//
//	@Summary		Login a User
//	@Description	Login as user to application
//	@Tags			auth
//	@Accept			application/json
//	@Produce		application/json
//	@Param 			data body dto.LoginRequest true "Post"
//	@Router			/login [post]
//	@Success		200  {object}  response.Response{data=dto.LoginResponse}  "ok"
//	@Failure		400  {object}  response.Response{data=[]response.ValidationErrors}  "bad request"
func (c AuthController) Login(ctx echo.Context) error {
	login := new(dto.LoginRequest)
	if err := ctx.Bind(login); err != nil {
		return response.BadRequest{Req: dto.LoginRequest{}, Message: err}.JSON(ctx)
	}

	token, err := c.authService.Login(login)
	if err != nil {
		return response.PolicyResponse{Message: err}.JSON(ctx)
	}

	return response.Response{Data: token}.JSON(ctx)
}
