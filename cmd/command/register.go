package command

import (
	"errors"
	"fmt"
	"github.com/army-of-one/generoo/pkg/registry"
	"github.com/spf13/cobra"
)

var Register = &cobra.Command{
	Use:   "register",
	Short: "Validate a generoo configuration file.",
	Long:  `validate a generoo configuration file against the json schema for generoo configuration files.`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var name string

		if len(args) == 1 {
			name, err = registry.Register(args[0])
		} else if len(args) == 0 {
			err = registry.RegisterLocal()
		} else {
			err = errors.New("usage: Generoo register [directory]")
		}

		if err != nil {
			fmt.Println("Failed to register Generoo template!")
			return
		}

		fmt.Printf("Successfully registered %s as a Generoo template!", name)
	},
}
