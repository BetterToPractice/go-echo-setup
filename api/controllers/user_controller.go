package controllers

import (
	"github.com/BetterToPractice/go-echo-setup/api/services"
	"github.com/BetterToPractice/go-echo-setup/lib"
	"github.com/BetterToPractice/go-echo-setup/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserController struct {
	logger      lib.Logger
	userService services.UserService
}

func NewUserController(logger lib.Logger, userService services.UserService) UserController {
	return UserController{
		logger:      logger,
		userService: userService,
	}
}

// List godoc
//
//	@Summary		List several users
//	@Description	get list several users
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Router			/users [get]
func (c UserController) List(ctx echo.Context) error {
	params := new(models.UserQueryParams)
	if err := ctx.Bind(params); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	qr, err := c.userService.Query(params)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, qr)
}

// Detail godoc
//
//		@Summary		Get a User
//		@Description	get a user by username
//		@Tags			user
//		@Accept			json
//		@Produce		json
//	    @Param			username	path	string	true	"Username"
//		@Router			/users/{username} [get]
func (c UserController) Detail(ctx echo.Context) error {
	user, err := c.userService.GetByUsername(ctx.Param("username"))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, err.Error())
	}
	return ctx.JSON(http.StatusOK, user)
}
