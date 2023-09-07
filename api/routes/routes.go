package routes

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewMainRoutes),
	fx.Provide(NewUserRoutes),
	fx.Provide(NewPostRoutes),
	fx.Provide(NewRoutes),
)

type IRoute interface {
	Setup()
}

type Routes []IRoute

func NewRoutes(
	mainRoutes MainRoutes,
	userRoutes UserRoutes,
	postRoutes PostRoutes,
) Routes {
	return Routes{
		mainRoutes,
		userRoutes,
		postRoutes,
	}
}

func (a Routes) Setup() {
	for _, route := range a {
		route.Setup()
	}
}
