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

func (m CorsMiddleware) Setup() {
	m.logger.Zap.Info("Setup cors middleware")
	m.handler.Engine.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     m.config.Cors.AllowOrigins,
		AllowMethods:     m.config.Cors.AllowMethods,
		AllowHeaders:     m.config.Cors.AllowHeaders,
		AllowCredentials: m.config.Cors.AllowCredentials,
	}))
}
