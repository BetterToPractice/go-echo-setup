package controllers

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewMainController),
	fx.Provide(NewUserController),
	fx.Provide(NewAuthController),
	fx.Provide(NewPostController),
)
