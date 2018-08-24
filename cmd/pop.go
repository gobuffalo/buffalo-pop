package cmd

import (
	"github.com/gobuffalo/buffalo/buffalo/cmd/destroy"
	"github.com/gobuffalo/pop/soda/cmd"
	"github.com/spf13/cobra"
)

// popCmd represents the pop command
var popCmd = cmd.RootCmd

func init() {
	popCmd.Use = "pop"
	popCmd.Aliases = append([]string{"db"}, popCmd.Aliases...)

	var destroyCmd = &cobra.Command{
		Use:     "destroy",
		Short:   "Allows to destroy generated code.",
		Aliases: []string{"d"},
	}

	destroyCmd.AddCommand(destroy.ModelCmd)
	popCmd.AddCommand(destroyCmd)
	rootCmd.AddCommand(popCmd)
}
