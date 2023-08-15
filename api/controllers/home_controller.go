package controllers

import (
	"github.com/BetterToPractice/go-echo-setup/lib"
	"github.com/labstack/echo/v4"
	"net/http"
)

type HomeController struct {
	logger lib.Logger
}

func NewHomeController(logger lib.Logger) HomeController {
	return HomeController{
		logger: logger,
	}
}

func (a HomeController) Get(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello World!")
}
