/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/

package cmd

import (
	"fmt"
	"github.com/Goodnessuc/bodyguards/config"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "The init command initializes go-bodyguards for your project",
	Long: `
	The init command is a fundamental command in the go-bodyguards framework. 
	It is designed to initialize and set up the go-bodyguards framework for your project. 
	This command is an essential step to quickly get started with using the go-bodyguards framework for your Go applications.

    When executing the init command, it performs several tasks to ensure that your project is properly set up and ready to use.
	These tasks include installing dependencies, setting up necessary files, and configuring the framework to match your project's requirements.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
		err := config.SetupDependencies()
		if err != nil {
			fmt.Println("Error: failed to install dependencies:", err)
		} else {
			fmt.Println("Setup files and Dependencies installed successfully")
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
