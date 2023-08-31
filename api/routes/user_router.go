package routes

import (
	"github.com/BetterToPractice/go-echo-setup/api/controllers"
	"github.com/BetterToPractice/go-echo-setup/lib"
)

type UserRoutes struct {
	handler        lib.HttpHandler
	userController controllers.UserController
}

func NewUserRoutes(handler lib.HttpHandler, userController controllers.UserController) UserRoutes {
	return UserRoutes{
		handler:        handler,
		userController: userController,
	}
}

func (a UserRoutes) Setup() {
	a.handler.Engine.GET("/users", a.userController.List)
	a.handler.Engine.GET("/users/:username", a.userController.Detail)
}
