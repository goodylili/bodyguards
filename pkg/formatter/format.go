package formatter

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func FormatGoCode(projectPath string) error {
	// Traverse through all the Go files in the project and format them using gofmt
	err := filepath.Walk(projectPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			// Run gofmt on the file to format it
			cmd := exec.Command("gofmt", "-w", path)
			if err := cmd.Run(); err != nil {
				return fmt.Errorf("failed to format file %s: %s", path, err)
			}
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to traverse project path: %s", err)
	}
	return nil
}

func LogFormattingIssues(projectPath string) error {
	// Traverse through all the Go files in the project and check their import statements using goimports
	err := filepath.Walk(projectPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			// Run goimports on the file to check its import statements
			cmd := exec.Command("goimports", "-l", path)
			output, err := cmd.CombinedOutput()
			if err != nil {
				return fmt.Errorf("failed to check file %s: %s", path, err)
			}
			if len(output) > 0 {
				fmt.Printf("Formatting issues found in file %s:\n%s\n", path, output)
			}
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to traverse project path: %s", err)
	}
	return nil
}
