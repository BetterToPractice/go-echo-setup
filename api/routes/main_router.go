package routes

import (
	"github.com/BetterToPractice/go-echo-setup/api/controllers"
	"github.com/BetterToPractice/go-echo-setup/lib"
)

type MainRoutes struct {
	handler        lib.HttpHandler
	mainController controllers.MainController
	swagger        lib.Swagger
}

func NewMainRoutes(handler lib.HttpHandler, mainController controllers.MainController, swagger lib.Swagger) MainRoutes {
	return MainRoutes{
		handler:        handler,
		mainController: mainController,
		swagger:        swagger,
	}
}

func (r MainRoutes) Setup() {
	r.swagger.Setup()
	r.handler.Engine.GET("", r.mainController.Index)
}
