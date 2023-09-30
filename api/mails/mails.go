package mails

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewAuthMail),
	fx.Provide(NewPostMail),
)
