package routes

import (
	"github.com/BetterToPractice/go-echo-setup/api/controllers"
	"github.com/BetterToPractice/go-echo-setup/lib"
)

type UserRouter struct {
	handler        lib.HttpHandler
	userController controllers.UserController
}

func NewUserRouter(handler lib.HttpHandler, userController controllers.UserController) UserRouter {
	return UserRouter{
		handler:        handler,
		userController: userController,
	}
}

func (r UserRouter) Setup() {
	r.handler.Engine.GET("/users", r.userController.List)
	r.handler.Engine.GET("/users/:username", r.userController.Detail)
	r.handler.Engine.DELETE("/users/:username", r.userController.Destroy)
}
