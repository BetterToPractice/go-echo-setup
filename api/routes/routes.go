package routes

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewHomeRoutes),
	fx.Provide(NewUserRoutes),
	fx.Provide(NewRoutes),
)

type IRoute interface {
	Setup()
}

type Routes []IRoute

func NewRoutes(
	homeRoutes HomeRoutes,
	userRoutes UserRoutes,
) Routes {
	return Routes{
		homeRoutes,
		userRoutes,
	}
}

func (a Routes) Setup() {
	for _, route := range a {
		route.Setup()
	}
}
