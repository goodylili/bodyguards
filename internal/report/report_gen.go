package report

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func RunCommands() (string, error) {
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

	// Execute goreportcard-cli -v
	cmd = exec.Command("goreportcard-cli", "-v")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to execute 'goreportcard-cli': %w", err)
	}

	return strings.TrimSpace(string(output)), nil
}
