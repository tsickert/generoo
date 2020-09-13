package command

import (
	"github.com/spf13/cobra"
	"log"
)

var Generate = &cobra.Command{
	Use:   "generate",
	Short: "Render a template with an input file.",
	Long:  `read through a template and render the information as provided through configuration or user input.`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Generatin'...")
	},
}
