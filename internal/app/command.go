package main

import "os/exec"

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

// Command to Execute
type Command struct {
	Command string `json:"command"`
}

// Execute the command in the Command struct
func (command *Command) Execute() error {
	cmd := execCommand(command.Command)
	return cmd.Run()
}
