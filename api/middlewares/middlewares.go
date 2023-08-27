package middlewares

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewCorsMiddleware),
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
) Middlewares {
	return Middlewares{
		corsMiddleware,
		gzipMiddleware,
		secureMiddleware,
	}
}

func (a Middlewares) Setup() {
	for _, middleware := range a {
		middleware.Setup()
	}
}
