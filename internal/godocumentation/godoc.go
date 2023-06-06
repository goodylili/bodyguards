package godocumentation

import (
	"log"
	"os/exec"
)

func RunDocsCommand() string {

	// Start godoc server
	cmd := exec.Command("godoc", "-http=:6060")
	err := cmd.Start()
	if err != nil {
		log.Fatalf("Error starting godoc server: %v", err)
	}

	return "Head to http://localhost:6060/ to view your project's godocumentation"
}
