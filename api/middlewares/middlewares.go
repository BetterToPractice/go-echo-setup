package middlewares

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewCorsMiddleware),
	fx.Provide(NewMiddlewares),
)

type IMiddleware interface {
	Setup()
}

type Middlewares []IMiddleware

func NewMiddlewares(
	corsMiddleware CorsMiddleware,
) Middlewares {
	return Middlewares{
		corsMiddleware,
	}
}

func (a Middlewares) Setup() {
	for _, middleware := range a {
		middleware.Setup()
	}
}
