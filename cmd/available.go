package cmd

import (
	"encoding/json"
	"os"

	"github.com/gobuffalo/buffalo-pop/v3/cmd/destroy"
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/pop/v6/logging"
	"github.com/spf13/cobra"
)

// availableCmd represents the available command
var availableCmd = &cobra.Command{
	Use:   "available",
	Short: "a list of available buffalo plugins",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		// we need to mute the pop logger for this cmd because it pollutes
		// the output of the command and plugins can't unmarshal the output
		pop.SetLogger(func(lvl logging.Level, s string, args ...interface{}) {})
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		plugs := Commands{
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

// Command that the plugin supplies
type Command struct {
	// Name "foo"
	Name string `json:"name"`
	// UseCommand "bar"
	UseCommand string `json:"use_command"`
	// BuffaloCommand "generate"
	BuffaloCommand string `json:"buffalo_command"`
	// Description "generates a foo"
	Description string   `json:"description,omitempty"`
	Aliases     []string `json:"aliases,omitempty"`
	Binary      string   `json:"-"`
	Flags       []string `json:"flags,omitempty"`
	// Filters events to listen to ("" or "*") is all events
	ListenFor string `json:"listen_for,omitempty"`
}

// Commands is a slice of Command
type Commands []Command
