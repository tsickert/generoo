package command

import (
	"errors"
	"github.com/army-of-one/generoo/pkg/registry"
	"github.com/spf13/cobra"
	"log"
)

var Link = &cobra.Command{
	Use:   "link",
	Short: "Create link to repository.",
	Long:  `Create link to repository.`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		if len(args) == 1 {
			err = registry.Link(args[0])
		} else if len(args) == 0 {
			err = registry.LinkLocal()
		} else {
			err = errors.New("usage: Generoo link [target]")
		}

		if err != nil {
			log.Fatal("failed to link new Generoo template in registry!")
		}

		log.Println("successfully linked new Generoo template in registry!")
	},
}
