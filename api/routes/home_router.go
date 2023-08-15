package routes

import (
	"github.com/BetterToPractice/go-echo-setup/api/controllers"
	"github.com/BetterToPractice/go-echo-setup/lib"
)

type HomeRoutes struct {
	homeController controllers.HomeController
	handler        lib.HttpHandler
}

func NewHomeRoutes(handler lib.HttpHandler, homeController controllers.HomeController) HomeRoutes {
	return HomeRoutes{
		homeController: homeController,
		handler:        handler,
	}
}

func (a HomeRoutes) Setup() {
	a.handler.Engine.GET("", a.homeController.Get)
}
