package command

import (
	"github.com/spf13/cobra"
	"log"
)

var Init = &cobra.Command{
	Use:   "init",
	Short: "Initialize generoo project.",
	Long:  `Initialize generoo project.`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("We're still implementing this feature")
	},
}
