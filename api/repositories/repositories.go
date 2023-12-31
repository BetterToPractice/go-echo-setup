package repositories

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewUserRepository),
	fx.Provide(NewPostRepository),
	fx.Provide(NewProfileRepository),
	fx.Provide(NewAuthRepository),
)
