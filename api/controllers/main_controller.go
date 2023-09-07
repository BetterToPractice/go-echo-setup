package controllers

import (
	"github.com/BetterToPractice/go-echo-setup/lib"
	"github.com/labstack/echo/v4"
)

type MainController struct {
	db lib.Database
}

func NewMainController(db lib.Database) MainController {
	return MainController{
		db: db,
	}
}

func (c MainController) Index(ctx echo.Context) error {
	return ctx.String(200, "omedeto")
}
