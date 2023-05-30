package report

import (
	"fmt"
	"os/exec"
	"strings"
)

func RunReportCommand() (string, error) {

	// Execute goreportcard-cli -v
	cmd := exec.Command("goreportcard-cli", "-v")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to execute 'goreportcard-cli': %w", err)
	}

	return strings.TrimSpace(string(output)), nil
}
