package libinput

import (
	"fmt"
	"os/exec"
)

type execCommandType func(string, ...string) *exec.Cmd

var execCommand = exec.Command

// SetExecCommand should only be used for testing when you need to mock exec.Command
func SetExecCommand(command execCommandType) {
	execCommand = command
}

// ResetExecCommand should only be used for testing to reset the command to exec.Command
func ResetExecCommand() {
	execCommand = exec.Command
}

// GetVersion returns the version of libinput on the device or error if unable to get the version
func GetVersion() (string, error) {
	cmd := execCommand("libinput", "--version")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("GetVersion: install libinput-tools")
	}

	return string(output), nil
}
