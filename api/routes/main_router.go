package routes

import (
	"github.com/BetterToPractice/go-echo-setup/api/controllers"
	"github.com/BetterToPractice/go-echo-setup/lib"
)

type MainRouter struct {
	handler        lib.HttpHandler
	mainController controllers.MainController
	swagger        lib.Swagger
}

func NewMainRouter(handler lib.HttpHandler, mainController controllers.MainController, swagger lib.Swagger) MainRouter {
	return MainRouter{
		handler:        handler,
		mainController: mainController,
		swagger:        swagger,
	}
}

func (r MainRouter) Setup() {
	r.swagger.Setup()
	r.handler.Engine.GET("", r.mainController.Index)
}
