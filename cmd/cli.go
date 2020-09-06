package main

import (
	"errors"
	"fmt"
	"github.com/army-of-one/generoo/pkg/registry"
	"github.com/army-of-one/generoo/pkg/schema"
	"github.com/army-of-one/generoo/pkg/utils"
	"github.com/spf13/cobra"
	"log"
	"runtime"
)

func main() {

	var outputDir string
	var inputDir string

	validator := schema.NewValidator()

	var cmdGenerate = &cobra.Command{
		Use:   "generate",
		Short: "Render a template with an input file.",
		Long:  `read through a template and render the information as provided through configuration or user input.`,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("Generatin'...")
		},
	}

	var cmdRegister = &cobra.Command{
		Use:   "register",
		Short: "Validate a generoo configuration file.",
		Long:  `validate a generoo configuration file against the json schema for generoo configuration files.`,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			goos := runtime.GOOS
			if !(goos == "linux" || goos == "darwin") {
				log.Printf("sorry, the 'register' command is not a supported on the %s platform for this feature yet\n", goos)
				return
			}

			var err error

			if len(args) == 1 {
				err = registry.Register(args[0])
			} else if len(args) == 0 {
				err = registry.RegisterLocal()
			} else {
				err = errors.New("usage: Generoo register [directory]")
			}

			if err != nil {
				log.Fatal("failed to registered new Generoo template!")
			}

			log.Println("successfully registered new Generoo template!")

		},
	}

	var cmdList = &cobra.Command{
		Use:   "list",
		Short: "List registered generoo templates.",
		Long:  `validate a generoo configuration file against the json schema for generoo configuration files.`,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			os := runtime.GOOS
			if !(os == "linux" || os == "darwin") {
				log.Printf("sorry, the 'list' command is not a supported on the %s platform for this feature yet\n", os)
				return
			}
		},
	}

	var cmdValidate = &cobra.Command{
		Use:   "validate",
		Short: "Validate a generoo configuration file.",
		Long:  `validate a generoo configuration file against the json schema for generoo configuration files.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			rawData, err := utils.ReadAmbiguousAsJson(args[0])
			if err != nil {
				log.Fatalf("could not read %s", args[1])
			}
			issues, err := validator.Validate(rawData)
			if err != nil {
				log.Printf("Failed during validate: %s", err)
			} else {
				if len(issues) < 1 {
					log.Println("Provided configuration is valid.")
				} else {
					failureMsg := "Failed to validate provided configuration file.\n"
					for _, err := range issues {
						failureMsg = fmt.Sprintf("%s\n%s", failureMsg, err.Message)
					}
					log.Println(failureMsg)
				}
			}
		},
	}

	cmdGenerate.Flags().StringVarP(&outputDir, "output", "o", ".", "directory to output rendered template")
	cmdGenerate.Flags().StringVarP(&inputDir, "input", "i", ".", "directory where template is located")

	var rootCmd = &cobra.Command{Use: "generoo"}
	rootCmd.AddCommand(cmdGenerate)
	rootCmd.AddCommand(cmdValidate)
	rootCmd.AddCommand(cmdList)
	rootCmd.AddCommand(cmdRegister)
	rootCmd.Execute()
}
