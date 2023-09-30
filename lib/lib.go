package lib

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewHttpHandler),
	fx.Provide(NewConfig),
	fx.Provide(NewLogger),
	fx.Provide(NewDatabase),
	fx.Provide(NewMigration),
	fx.Provide(NewMail),
)
