package config

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// InstallDependencies rewrite to install dependencies based on yaml
func InstallDependencies() (string, error) {
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

	// Install godoc command
	cmd = exec.Command("go", "install", "golang.org/x/tools/cmd/godoc@latest")
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Error installing godoc: %v", err)
	}

	// Command 1: go install golang-ci-lint
	cmd = exec.Command("go", "install", "github.com/golangci/golangci-lint/cmd/golangci-lint@v1.51.1")
	if err := cmd.Run(); err != nil {
		log.Fatal("Failed to install golangci-lint:", err)
	}

	return "Installations Complete", nil
}
