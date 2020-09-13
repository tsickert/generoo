package command

import (
	"fmt"
	"github.com/army-of-one/generoo/pkg/schema"
	"github.com/army-of-one/generoo/pkg/utils"
	"github.com/spf13/cobra"
)

var Validate = &cobra.Command{
	Use:   "validate",
	Short: "Validate a generoo configuration file.",
	Long:  `validate a generoo configuration file against the json schema for generoo configuration files.`,
	Args:  cobra.MinimumNArgs(1),
	Run:   validate,
}

func validate(cmd *cobra.Command, args []string) {
	validator := schema.NewValidator()

	rawData, err := utils.ReadAmbiguousAsJson(args[0])

	if err != nil {
		fmt.Printf("Error: could not read %s", args[1])
		return
	}

	issues, err := validator.Validate(rawData)

	if err != nil {
		fmt.Printf("There were validation failures: %s", err)
	} else {
		if len(issues) < 1 {
			fmt.Println("Provided configuration is valid.")
		} else {
			var failureMsg string
			for _, err := range issues {
				failureMsg = fmt.Sprintf("%s\n%s", failureMsg, err.Message)
			}
			fmt.Println(failureMsg)
		}
	}
}
