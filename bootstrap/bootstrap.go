package bootstrap

import (
	"context"
	"errors"
	"fmt"
	"github.com/BetterToPractice/go-echo-setup/lib"
	"go.uber.org/fx"
	"net/http"
)

var Module = fx.Options(
	lib.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(lifecycle fx.Lifecycle, handler lib.HttpHandler) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := handler.Engine.Start(":1323"); err != nil {
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
			handler.Engine.Close()
			return nil
		},
	})
}
