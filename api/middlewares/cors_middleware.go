package middlewares

import (
	"github.com/BetterToPractice/go-echo-setup/lib"
	"github.com/labstack/echo/v4/middleware"
)

type CorsMiddleware struct {
	handler lib.HttpHandler
	logger  lib.Logger
	config  lib.Config
}

func NewCorsMiddleware(handler lib.HttpHandler, logger lib.Logger, config lib.Config) CorsMiddleware {
	return CorsMiddleware{
		handler: handler,
		logger:  logger,
		config:  config,
	}
}

func (a CorsMiddleware) Setup() {
	a.logger.Zap.Info("Setup cors middleware")
	a.handler.Engine.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     a.config.Cors.AllowOrigins,
		AllowMethods:     a.config.Cors.AllowMethods,
		AllowHeaders:     a.config.Cors.AllowHeaders,
		AllowCredentials: a.config.Cors.AllowCredentials,
	}))
}
