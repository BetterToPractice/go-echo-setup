package routes

import (
	"github.com/BetterToPractice/go-echo-setup/api/controllers"
	"github.com/BetterToPractice/go-echo-setup/lib"
)

type MainRouter struct {
	handler        lib.HttpHandler
	mainController controllers.MainController
}

func NewMainRouter(handler lib.HttpHandler, mainController controllers.MainController) MainRouter {
	return MainRouter{
		handler:        handler,
		mainController: mainController,
	}
}

func (r MainRouter) Setup() {
	r.handler.Engine.GET("", r.mainController.Index)
}
