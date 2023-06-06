package linter

import (
	"log"
	"os"
	"os/exec"
)

func LintPrograms() string {
	cmd := exec.Command("golangci-lint", "run", "--config=config/.golangci.yml")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error running golangci-lint: %s", err)
	}

	return "Golangci-lint is running"
}
