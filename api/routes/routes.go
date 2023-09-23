package routes

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewMainRouter),
	fx.Provide(NewUserRouter),
	fx.Provide(NewAuthRouter),
	fx.Provide(NewPostRouter),
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
) Routes {
	return Routes{
		mainRouter,
		userRouter,
		postRouter,
		authRouter,
	}
}

func (a Routes) Setup() {
	for _, route := range a {
		route.Setup()
	}
}
