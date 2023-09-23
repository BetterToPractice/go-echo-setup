package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type MainController struct {
}

func NewMainController() MainController {
	return MainController{}
}

func (c MainController) Index(ctx echo.Context) error {
	return ctx.Redirect(http.StatusFound, "/swagger/index.html")
}
