package controllers

import (
	"github.com/BetterToPractice/go-echo-setup/api/services"
	"github.com/BetterToPractice/go-echo-setup/models"
	"github.com/BetterToPractice/go-echo-setup/pkg/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return UserController{
		userService: userService,
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
func (c UserController) List(ctx echo.Context) error {
	params := new(models.UserQueryParams)
	if err := ctx.Bind(params); err != nil {
		return response.Response{Code: http.StatusBadRequest, Message: err}.JSON(ctx)
	}

	qr, err := c.userService.Query(params)
	if err != nil {
		return response.Response{Code: http.StatusBadRequest, Message: err}.JSON(ctx)
	}

	return response.Response{Code: http.StatusOK, Data: qr}.JSON(ctx)
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
func (c UserController) Detail(ctx echo.Context) error {
	user, err := c.userService.GetByUsername(ctx.Param("username"))
	if err != nil {
		return response.Response{Code: http.StatusNotFound, Message: err}.JSON(ctx)
	}
	return response.Response{Code: http.StatusOK, Data: user}.JSON(ctx)
}

// Destroy godoc
//
//	@Summary		Delete a User
//	@Description	perform delete a user by username
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			username  path  string  true  "Username"
//	@Router			/users/{username} [delete]
func (c UserController) Destroy(ctx echo.Context) error {
	if err := c.userService.Delete(ctx.Param("username")); err != nil {
		return response.Response{Code: http.StatusBadRequest, Message: err}.JSON(ctx)
	}
	return response.Response{Code: http.StatusNoContent}.JSON(ctx)
}
