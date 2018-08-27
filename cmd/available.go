package cmd

import (
	"encoding/json"
	"os"

	"github.com/gobuffalo/buffalo-plugins/plugins"
	"github.com/gobuffalo/buffalo-pop/cmd/destroy"
	"github.com/spf13/cobra"
)

// availableCmd represents the available command
var availableCmd = &cobra.Command{
	Use:   "available",
	Short: "a list of available buffalo plugins",
	RunE: func(cmd *cobra.Command, args []string) error {
		plugs := plugins.Commands{
			{Name: "db", BuffaloCommand: "root", Description: "[DEPRECATED] please use `buffalo pop` instead.", Aliases: popCmd.Aliases},
			{Name: "pop", BuffaloCommand: "root", Description: popCmd.Short, Aliases: popCmd.Aliases},
			{Name: "model", BuffaloCommand: "destroy", Description: destroy.ModelCmd.Short, Aliases: destroy.ModelCmd.Aliases},
		}
		return json.NewEncoder(os.Stdout).Encode(plugs)
	},
}

func init() {
	rootCmd.AddCommand(availableCmd)
}
