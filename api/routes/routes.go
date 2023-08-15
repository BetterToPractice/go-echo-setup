package routes

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewHomeRoutes),
	fx.Provide(NewRoutes),
)

type IRoute interface {
	Setup()
}

type Routes []IRoute

func NewRoutes(
	homeRoutes HomeRoutes,
) Routes {
	return Routes{
		homeRoutes,
	}
}

func (a Routes) Setup() {
	for _, route := range a {
		route.Setup()
	}
}
