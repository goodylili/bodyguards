package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
)

func main() {
	// Retrieve the package name of the main package
	packageName := getPackageName()

	// Execute go fmt command to format the package
	err := formatPackage(packageName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Go files in the package have been formatted successfully.")
}

// getPackageName retrieves the package name of the main package
func getPackageName() string {
	// Retrieve the absolute path of the main executable
	executablePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	// Retrieve the directory of the main executable
	executableDir := filepath.Dir(executablePath)

	// Retrieve the absolute path of the main package
	packageName := filepath.Base(executableDir)

	// Retrieve the import path of the main package
	importPath := fmt.Sprintf("%s/%s", os.Getenv("GOPATH"), packageName)

	// Resolve the import path to the package name
	pkg, err := reflect.Import(importPath)
	if err != nil {
		log.Fatal(err)
	}

	return pkg.Name()
}

// formatPackage executes the "go fmt" command to format all Go files in the specified package
func formatPackage(packageName string) error {
	cmd := exec.Command("go", "fmt", packageName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
