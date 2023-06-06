/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/

package cmd

import (
	"fmt"
	config2 "github.com/Goodnessuc/bodyguards/config"
	"github.com/Goodnessuc/bodyguards/internal/linter"
	"github.com/spf13/cobra"
	"log"
)

// linterCmd represents the linter command
var linterCmd = &cobra.Command{
	Use:   "linter",
	Short: "The linter command runs a linter that fixes your code based on Go standards",
	Long: `
	The linter command utilizes Golangci-lint, a popular linter for Go programs, to analyze and improve the quality of your code. 
	It applies a set of predefined rules and standards to identify potential issues, enforce best practices, and suggest improvements. 
	By running the linter, you can automatically detect and fix various code problems such as unused variables, incorrect formatting, incorrect imports, improper error handling, code complexity, and more. 
	It helps ensure that your code follows the idiomatic Go style and adheres to recommended practices.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := config2.ReadBodyguardsYAML()
		if err != nil {
			log.Fatalf("Error reading Bodyguards YAML: %v", err)
		}
		enabledSlice := config.Run.Enable
		for _, value := range enabledSlice {
			if value == "linter" {
				fmt.Println(linter.LintPrograms())
				return
			}
		}
		fmt.Println("You didn't specify linting as an option is your YAML file")

	},
}

func init() {
	rootCmd.AddCommand(linterCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// linterCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// linterCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
