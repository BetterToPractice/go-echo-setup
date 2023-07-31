package runserver

import (
	"github.com/BetterToPractice/go-echo-setup/bootstrap"
	"github.com/BetterToPractice/go-echo-setup/lib"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var configFile string

func init() {
	pf := StartCmd.PersistentFlags()
	pf.StringVarP(&configFile, "config", "c", "config/config.yaml", "this parameter is used to start the service application")
}

var StartCmd = &cobra.Command{
	Use: "runserver",
	PreRun: func(cmd *cobra.Command, args []string) {
		lib.SetConfigPath(configFile)
	},
	Run: func(cmd *cobra.Command, args []string) {
		runApplication()
	},
}

func runApplication() {
	fx.New(bootstrap.Module, fx.NopLogger).Run()
}
