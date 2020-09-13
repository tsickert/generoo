package main

import (
	"github.com/army-of-one/generoo/cmd/command"
	"github.com/spf13/cobra"
)

func main() {

	var outputDir string
	var inputDir string

	command.Generate.Flags().StringVarP(&outputDir, "output", "o", ".", "directory to output rendered template")
	command.Generate.Flags().StringVarP(&inputDir, "input", "i", ".", "directory where template is located")

	var rootCmd = &cobra.Command{Use: "generoo"}
	rootCmd.AddCommand(command.Generate)
	rootCmd.AddCommand(command.Validate)
	rootCmd.AddCommand(command.List)
	rootCmd.AddCommand(command.Register)
	rootCmd.AddCommand(command.Init)
	rootCmd.AddCommand(command.Link)
	_ = rootCmd.Execute()
}
