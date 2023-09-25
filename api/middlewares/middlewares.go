package middlewares

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewCorsMiddleware),
	fx.Provide(NewAuthMiddleware),
	fx.Provide(NewGZipMiddleware),
	fx.Provide(NewSecureMiddleware),
	fx.Provide(NewMiddlewares),
)

type IMiddleware interface {
	Setup()
}

type Middlewares []IMiddleware

func NewMiddlewares(
	corsMiddleware CorsMiddleware,
	gzipMiddleware GZipMiddleware,
	secureMiddleware SecureMiddleware,
	authMiddleware AuthMiddleware,
) Middlewares {
	return Middlewares{
		corsMiddleware,
		gzipMiddleware,
		secureMiddleware,
		authMiddleware,
	}
}

func (a Middlewares) Setup() {
	for _, middleware := range a {
		middleware.Setup()
	}
}
