package bootstrap

import (
	"context"
	"errors"
	"fmt"
	"github.com/BetterToPractice/go-echo-setup/api/middlewares"
	"github.com/BetterToPractice/go-echo-setup/lib"
	"go.uber.org/fx"
	"net/http"
)

var Module = fx.Options(
	lib.Module,
	middlewares.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(lifecycle fx.Lifecycle, handler lib.HttpHandler, config lib.Config, middlewares middlewares.Middlewares) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				middlewares.Setup()

				if err := handler.Engine.Start(config.Http.ListenAddr()); err != nil {
					if errors.Is(err, http.ErrServerClosed) {
						fmt.Println("run error", err)
					} else {
						fmt.Println("other error happens", err)
					}
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			if err := handler.Engine.Close(); err != nil {
			}
			return nil
		},
	})
}
