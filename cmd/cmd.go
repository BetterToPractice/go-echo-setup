package cmd

import (
	"errors"
	"github.com/BetterToPractice/go-echo-setup/cmd/runserver"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(runserver.StartCmd)
}

var rootCmd = &cobra.Command{
	Use: "go-echo-setup",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New(
				"requires at least one arg, " +
					"you can view the available parameters through `--help`",
			)
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
