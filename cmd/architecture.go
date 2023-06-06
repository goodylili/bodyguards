/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
	"fmt"
	config2 "github.com/Goodnessuc/bodyguards/config"
	"github.com/Goodnessuc/bodyguards/internal/architecture"

	"github.com/spf13/cobra"
)

// architectureCmd represents the architecture command
var architectureCmd = &cobra.Command{
	Use:   "architecture",
	Short: "Set up file architecture based on project specifications in YAML file",
	Long:  "This command sets up the file architecture for a project by utilizing the specifications provided in a YAML file. It creates the necessary files and directories to establish the desired project structure.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("architecture called")

		// Retrieve the architecture specified
		YAMLconfig, err := config2.ReadBodyguardsYAML()
		if err != nil {
			fmt.Printf("Error reading YAML file: %v\n", err)
			return
		}

		// Run a function based on the specified architecture
		switch YAMLconfig.Run.Architecture {
		case "hexagonal":
			architecture.CreateHexagonal()
		case "microservice":
			architecture.CreateMicroserviceStructure()
		case "monolithic":
			architecture.CreateMonolithicDirectory()
		case "mvc":
			architecture.CreateMVCDirectories()
		default:
			fmt.Println("Invalid architecture specified in the YAML file.")
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(architectureCmd)

}
