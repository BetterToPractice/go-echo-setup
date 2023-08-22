package controllers

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewHomeController),
	fx.Provide(NewUserController),
)
