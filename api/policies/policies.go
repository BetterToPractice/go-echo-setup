package policies

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewPostPolicy),
	fx.Provide(NewUserPolicy),
)
