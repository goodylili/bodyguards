/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
	"fmt"
	config2 "github.com/Goodnessuc/bodyguards/config"
	"github.com/Goodnessuc/bodyguards/internal/report"
	"github.com/spf13/cobra"
	"log"
)

// reportCmd represents the report command
var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Display report ratings for the Go project",

	Long: `This command generates and displays report ratings for your Go project. 
	It executes the 'goreportcard-cli' command with the '-v' flag to generate the ratings. 
	The output is then returned as a string, trimming any leading or trailing whitespace.`,
	Run: func(cmd *cobra.Command, args []string) {

		config, err := config2.ReadBodyguardsYAML()
		if err != nil {
			log.Fatalf("Error reading Bodyguards YAML: %v", err)
		}
		enabledSlice := config.Run.Enable
		for _, value := range enabledSlice {
			if value == "report" {
				fmt.Println(report.RunReportCommand())
				return
			}
		}

		fmt.Println("You didn't specify report as an option is your YAML file")

	},
}

func init() {
	rootCmd.AddCommand(reportCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// reportCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// reportCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
