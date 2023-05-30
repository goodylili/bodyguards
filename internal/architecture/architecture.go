package architecture

import (
	"os"
	"os/exec"
)

// ExecuteBashScript executes a bash script based on the specifications from the YAML file to set up an architecture
func ExecuteBashScript(scriptName string) error {
	if scriptName == "clean" {
		return nil
	} else {
		cmd := exec.Command("bash", scriptName)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			return err
		}

	}

	return nil

}
