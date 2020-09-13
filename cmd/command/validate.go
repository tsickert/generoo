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
		fmt.Printf("Could not read provided file or directory")
		return
	}

	issues, err := validator.Validate(rawData)

	if err != nil {
		fmt.Printf("Generoo failed during pre-validation steps. Please add an issue in the Generoo repository " +
			"here: https://github.com/army-of-one/generoo/issues. Thanks, and sorry for the inconvenience.")
	} else {
		if len(issues) < 1 {
			fmt.Println("Provided configuration is valid.")
		} else {
			var failureMsg string
			for _, err := range issues {
				failureMsg = fmt.Sprintf("%s\nError at %s: %s", failureMsg, err.PropertyPath, err.Message)
			}
			fmt.Println(failureMsg)
		}
	}
}
