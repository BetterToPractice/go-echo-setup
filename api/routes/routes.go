package routes

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewMainRouter),
	fx.Provide(NewUserRouter),
	fx.Provide(NewAuthRouter),
	fx.Provide(NewPostRouter),
	fx.Provide(NewSwaggerRouter),
	fx.Provide(NewRoutes),
)

type IRoute interface {
	Setup()
}

type Routes []IRoute

func NewRoutes(
	mainRouter MainRouter,
	userRouter UserRouter,
	postRouter PostRouter,
	authRouter AuthRouter,
	swaggerRouter SwaggerRouter,
) Routes {
	return Routes{
		mainRouter,
		swaggerRouter,
		userRouter,
		postRouter,
		authRouter,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
