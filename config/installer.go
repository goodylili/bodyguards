package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

// SetupDependencies rewrite to install dependencies based on yaml and sets up permissions for architecture bash scripts
func SetupDependencies() error {
	// Give executable permission to files
	err := GiveExecutablePermissionToFiles()
	if err != nil {
		fmt.Println("Error: failed to give executable permission to files:", err)
		return err
	}

	// Install golangCI-lint
	msg, err := installGolangCILint()
	if err != nil {
		fmt.Println("Error: failed to install Golangci-lint:", err)
		return err
	}
	fmt.Println(msg)

	// Install Godoc
	msg, err = installGodoc()
	if err != nil {
		fmt.Println("Error: failed to install Godoc:", err)
		return err
	}
	fmt.Println(msg)

	// Install Goreportcard-cli
	msg, err = installGoReportCard()
	if err != nil {
		fmt.Println("Error: failed to install Goreportcard-cli:", err)
		return err
	}
	fmt.Println(msg)

	err = createBodyguardsYAML()
	if err != nil {
		log.Fatalf("Failed to create bodyguards.yaml: %v", err)
	}

	fmt.Println("bodyguards.yaml created successfully!")

	// All operations completed successfully
	return nil
}

func GiveExecutablePermissionToFiles() error {
	_, filename, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(filename)

	filePaths := []string{
		filepath.Join(baseDir, "hexagonal.go"),
		filepath.Join(baseDir, "microservice.sh"),
		filepath.Join(baseDir, "monolithic.sh"),
		filepath.Join(baseDir, "architecture", "mvc.sh"),
	}

	for _, filePath := range filePaths {
		err := os.Chmod(filePath, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}

func installGolangCILint() (string, error) {
	// Command 1: go install golang-ci-lint
	cmd := exec.Command("go", "install", "github.com/golangci/golangci-lint/cmd/golangci-lint@v1.51.1")
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to install Golangci-lint: %w", err)
	}
	return "GolangCI-lint installed successfully", nil
}

func installGodoc() (string, error) {
	// Install godoc command
	cmd := exec.Command("go", "install", "golang.org/x/tools/cmd/godoc@latest")
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("failed to install Godoc: %w", err)
	}

	return "Godoc installed successfully", nil
}

func installGoReportCard() (string, error) {
	// Git clone
	cmd := exec.Command("git", "clone", "https://github.com/gojp/goreportcard.git")
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("failed to clone repository: %w", err)
	}

	// Change directory
	err = os.Chdir("goreportcard")
	if err != nil {
		return "", fmt.Errorf("failed to change directory: %w", err)
	}

	// Make install
	cmd = exec.Command("make", "install")
	err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("failed to execute 'make install': %w", err)
	}

	// Go install
	cmd = exec.Command("go", "install", "./cmd/goreportcard-cli")
	err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("failed to execute 'go install': %w", err)
	}

	return "Goreportcard-cli successfully installed", nil
}

type BodyguardsConfig struct {
	Run struct {
		Architecture string   `yaml:"architecture"`
		Enable       []string `yaml:"enable"`
	} `yaml:"run"`
}

func createBodyguardsYAML() error {
	config := BodyguardsConfig{
		Run: struct {
			Architecture string   `yaml:"architecture"`
			Enable       []string `yaml:"enable"`
		}{
			Architecture: "clean",
			Enable: []string{
				"linter",
				"report",
				"documentation",
			},
		},
	}

	data, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("failed to marshal YAML: %v", err)
	}

	err = ioutil.WriteFile("bodyguards.yaml", data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write YAML file: %v", err)
	}

	return nil
}
