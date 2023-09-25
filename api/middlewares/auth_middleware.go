package middlewares

import (
	"github.com/BetterToPractice/go-echo-setup/api/services"
	"github.com/BetterToPractice/go-echo-setup/constants"
	"github.com/BetterToPractice/go-echo-setup/lib"
	"github.com/labstack/echo/v4"
	"strings"
)

type AuthMiddleware struct {
	config      lib.Config
	handler     lib.HttpHandler
	authService services.AuthService
}

func NewAuthMiddleware(config lib.Config, handler lib.HttpHandler, authService services.AuthService) AuthMiddleware {
	return AuthMiddleware{
		config:      config,
		authService: authService,
		handler:     handler,
	}
}

func (m AuthMiddleware) core() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			request := ctx.Request()

			var (
				auth   = request.Header.Get("Authorization")
				prefix = "Bearer "
				token  string
			)

			if auth != "" && strings.HasPrefix(auth, prefix) {
				token = auth[len(prefix):]
			}

			claims, err := m.authService.ParseToken(token)
			if err != nil {
				return next(ctx)
			}

			ctx.Set(constants.CurrentUser, claims)
			return next(ctx)
		}
	}
}

func (m AuthMiddleware) Setup() {
	m.handler.Engine.Use(m.core())
}
