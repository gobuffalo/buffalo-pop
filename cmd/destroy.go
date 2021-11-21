package cmd

import (
	"github.com/gobuffalo/buffalo-pop/v3/cmd/destroy"
	"github.com/spf13/cobra"
)

var destroyCmd = &cobra.Command{
	Use:     "destroy",
	Short:   "Allows to destroy generated code.",
	Aliases: []string{"d"},
}

func init() {
	destroyCmd.PersistentFlags().BoolVarP(&destroy.YesToAll, "yes", "y", false, "confirms all beforehand")
	destroyCmd.AddCommand(destroy.ModelCmd)
	popCmd.AddCommand(destroyCmd)
	rootCmd.AddCommand(destroyCmd)
}
