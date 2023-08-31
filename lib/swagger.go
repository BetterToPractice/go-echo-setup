package lib

import (
	"github.com/BetterToPractice/go-echo-setup/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Swagger struct {
	config  Config
	handler HttpHandler
}

func NewSwagger(config Config, handler HttpHandler) Swagger {
	return Swagger{
		config:  config,
		handler: handler,
	}
}

func (l Swagger) SetUrl() {
	l.handler.Engine.GET(l.config.Swagger.DocUrl, echoSwagger.WrapHandler)
}

func (l Swagger) Setup() {
	docs.SwaggerInfo.Title = l.config.Swagger.Title
	docs.SwaggerInfo.Description = l.config.Swagger.Description
	docs.SwaggerInfo.Version = l.config.Swagger.Version
	docs.SwaggerInfo.InfoInstanceName = "Meh"

	l.SetUrl()
}
