package middlewares

import (
	"github.com/BetterToPractice/go-echo-setup/lib"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"strings"
)

type GZipMiddleware struct {
	handler lib.HttpHandler
}

func NewGZipMiddleware(handler lib.HttpHandler) GZipMiddleware {
	return GZipMiddleware{
		handler: handler,
	}
}

func (m GZipMiddleware) Setup() {
	m.handler.Engine.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Skipper: func(c echo.Context) bool {
			if strings.Contains(c.Request().URL.Path, "swagger") {
				return true
			}
			return false
		},
	}))
}
