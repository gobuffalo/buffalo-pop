package cmd

import (
	"github.com/gobuffalo/pop/v6/soda/cmd"
)

// popCmd represents the pop command
var popCmd = cmd.RootCmd

func init() {
	popCmd.Use = "pop"
	popCmd.Aliases = append([]string{"db"}, popCmd.Aliases...)

	rootCmd.AddCommand(popCmd)
}
