package lib

import (
	"github.com/labstack/echo/v4"
)

type HttpHandler struct {
	Engine *echo.Echo
}

func NewHttpHandler() HttpHandler {
	engine := echo.New()

	httpHandler := HttpHandler{
		Engine: engine,
	}

	return httpHandler
}
