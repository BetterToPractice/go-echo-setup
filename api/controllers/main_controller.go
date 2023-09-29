package controllers

import (
	"github.com/BetterToPractice/go-echo-setup/lib"
	"github.com/labstack/echo/v4"
	"net/http"
)

type MainController struct {
	config lib.Config
}

func NewMainController(config lib.Config) MainController {
	return MainController{
		config: config,
	}
}

func (c MainController) Index(ctx echo.Context) error {
	return ctx.Redirect(http.StatusFound, c.config.Swagger.DocUrl)
}
