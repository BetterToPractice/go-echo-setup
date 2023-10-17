package controllers

import (
	"github.com/BetterToPractice/go-echo-setup/api/dto"
	"github.com/BetterToPractice/go-echo-setup/api/policies"
	"github.com/BetterToPractice/go-echo-setup/api/services"
	"github.com/BetterToPractice/go-echo-setup/pkg/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserController struct {
	userService services.UserService
	userPolicy  policies.UserPolicy
	authService services.AuthService
}

func NewUserController(userService services.UserService, authService services.AuthService, userPolicy policies.UserPolicy) UserController {
	return UserController{
		userService: userService,
		userPolicy:  userPolicy,
		authService: authService,
	}
}

// List godoc
//
//	@Summary		List several users
//	@Description	get list several users
//	@Tags			user
//	@Accept			application/json
//	@Produce		application/json
//	@Router			/users [get]
//	@Success		200  {object}  response.Response{data=dto.UserPaginationResponse}  "ok"
//	@Failure		400  {object}  response.Response{data=[]response.ValidationErrors}  "bad request"
func (c UserController) List(ctx echo.Context) error {
	params := new(dto.UserQueryParam)
	if err := ctx.Bind(params); err != nil {
		return response.BadRequest{Message: err}.JSON(ctx)
	}

	qr, err := c.userService.Query(params)
	if err != nil {
		return response.BadRequest{Message: err}.JSON(ctx)
	}

	return response.Response{Data: qr}.JSON(ctx)
}

// Detail godoc
//
//	@Summary		Get a User
//	@Description	get a user by username
//	@Tags			user
//	@Accept			application/json
//	@Produce		application/json
//	@Param			username  path  string  true  "Username"
//	@Router			/users/{username} [get]
//	@Success		200  {object}  response.Response{data=dto.UserResponse}  "ok"
//	@Failure		404  {object}  response.Response  "not found"
func (c UserController) Detail(ctx echo.Context) error {
	user, err := c.userService.GetByUsername(ctx.Param("username"))
	if err != nil {
		return response.NotFound{Message: err}.JSON(ctx)
	}

	resp := dto.UserResponse{}
	resp.Serializer(user)
	return response.Response{Data: resp}.JSON(ctx)
}

// Destroy godoc
//
//	@Summary		Delete a User
//	@Description	perform delete a user by username
//	@Tags			user
//	@Accept			application/json
//	@Produce		application/json
//	@Security 		BearerAuth
//	@Param			username  path  string  true  "Username"
//	@Router			/users/{username} [delete]
//	@Success		204  {object}  nil  "no content"
//	@Failure		404  {object}  response.Response  "not found"
//	@Failure		401  {object}  response.Response  "unauthorized"
//	@Failure		403  {object}  response.Response  "forbidden"
func (c UserController) Destroy(ctx echo.Context) error {
	user, err := c.userService.GetByUsername(ctx.Param("username"))
	if err != nil {
		return response.NotFound{Message: err}.JSON(ctx)
	}

	loggedInUser, _ := c.authService.Authenticate(ctx)
	if can, err := c.userPolicy.CanDelete(loggedInUser, user); !can {
		return response.PolicyResponse{Message: err}.JSON(ctx)
	}

	if err := c.userService.Delete(user); err != nil {
		return response.BadRequest{Message: err}.JSON(ctx)
	}

	return response.Response{Code: http.StatusNoContent}.JSON(ctx)
}
