package routes

import (
	"github.com/BetterToPractice/go-echo-setup/api/controllers"
	"github.com/BetterToPractice/go-echo-setup/lib"
)

type AuthRouter struct {
	authController controllers.AuthController
	handler        lib.HttpHandler
}

func NewAuthRouter(handler lib.HttpHandler, authController controllers.AuthController) AuthRouter {
	return AuthRouter{
		handler:        handler,
		authController: authController,
	}
}

func (r AuthRouter) Setup() {
	r.handler.Engine.POST("/register", r.authController.Register)
	r.handler.Engine.POST("/login", r.authController.Login)
}
