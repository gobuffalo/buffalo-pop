package cmd

import (
	"fmt"

	"github.com/gobuffalo/buffalo-pop/pop"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "current version of pop",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("pop", pop.Version)
		return nil
	},
}

func init() {
	popCmd.AddCommand(versionCmd)
}
