package architecture

import (
	"os"
	"os/exec"
)

func ExecuteBashScript(scriptName string) error {
	cmd := exec.Command("bash", scriptName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
