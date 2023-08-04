package middlewares

import (
	"github.com/BetterToPractice/go-echo-setup/lib"
	"github.com/labstack/echo/v4/middleware"
)

type CorsMiddleware struct {
	handler lib.HttpHandler
	logger  lib.Logger
}

func NewCorsMiddleware(handler lib.HttpHandler, logger lib.Logger) CorsMiddleware {
	return CorsMiddleware{
		handler: handler,
		logger:  logger,
	}
}

func (a CorsMiddleware) Setup() {
	a.logger.Zap.Info("Setup cors middleware")
	a.handler.Engine.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
		AllowHeaders:     []string{"*"},
		AllowMethods:     []string{"*"},
	}))
}
