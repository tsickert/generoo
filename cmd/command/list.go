package command

import (
	"github.com/army-of-one/generoo/pkg/registry"
	"github.com/spf13/cobra"
)

var List = &cobra.Command{
	Use:   "list",
	Short: "List registered generoo templates.",
	Long:  `validate a generoo configuration file against the json schema for generoo configuration files.`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		registry.List()
	},
}
