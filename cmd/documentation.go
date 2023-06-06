/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/

package cmd

import (
	"fmt"
	config2 "github.com/Goodnessuc/bodyguards/config"
	"github.com/Goodnessuc/bodyguards/internal/godocumentation"
	"github.com/spf13/cobra"
	"log"
)

// documentationCmd represents the docs command
var documentationCmd = &cobra.Command{
	Use:   "documentation",
	Short: "View project documentation",
	Long: `View the project's documentation by starting a local server and accessing it in a web browser.
This command starts a godoc server and provides the URL to access the documentation.`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := config2.ReadBodyguardsYAML()
		if err != nil {
			log.Fatalf("Error reading Bodyguards YAML: %v", err)
		}
		enabledSlice := config.Run.Enable
		for _, value := range enabledSlice {
			if value == "documentation" {
				fmt.Println(godocumentation.RunDocsCommand())
				return
			}
		}
		fmt.Println("You didn't specify documentation as an option is your YAML file")

	},
}

func init() {
	rootCmd.AddCommand(documentationCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// documentationCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// documentationCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
