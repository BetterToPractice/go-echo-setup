package routes

import (
	"github.com/BetterToPractice/go-echo-setup/api/controllers"
	"github.com/BetterToPractice/go-echo-setup/lib"
)

type PostRoutes struct {
	postController controllers.PostController
	handler        lib.HttpHandler
}

func NewPostRoutes(handler lib.HttpHandler, postController controllers.PostController) PostRoutes {
	return PostRoutes{
		postController: postController,
		handler:        handler,
	}
}

func (r PostRoutes) Setup() {
	r.handler.Engine.GET("/posts", r.postController.List)
	r.handler.Engine.GET("/posts/:id", r.postController.Detail)
	r.handler.Engine.DELETE("/posts/:id", r.postController.Destroy)
}
