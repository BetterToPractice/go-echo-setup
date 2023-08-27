package middlewares

import (
	"github.com/BetterToPractice/go-echo-setup/lib"
	"github.com/labstack/echo/v4/middleware"
)

type SecureMiddleware struct {
	handler lib.HttpHandler
}

func NewSecureMiddleware(handler lib.HttpHandler) SecureMiddleware {
	return SecureMiddleware{
		handler: handler,
	}
}

func (m SecureMiddleware) Setup() {
	m.handler.Engine.Use(middleware.Secure())
}
