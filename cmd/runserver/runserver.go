package runserver

import (
	"github.com/BetterToPractice/go-echo-setup/bootstrap"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var StartCmd = &cobra.Command{
	Use: "runserver",
	Run: func(cmd *cobra.Command, args []string) {
		runApplication()
	},
}

func runApplication() {
	fx.New(bootstrap.Module).Run()
}
