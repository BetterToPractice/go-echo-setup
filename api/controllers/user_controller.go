package controllers

import (
	"github.com/BetterToPractice/go-echo-setup/api/services"
	"github.com/BetterToPractice/go-echo-setup/lib"
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

func (c UserController) Index(ctx echo.Context) error {
	result, _ := c.userService.Query()
	return ctx.JSON(http.StatusOK, result)
}