package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"testing"
)

var mockedExitStatus = 0

func fakeExecCommand(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestExecCommandHelper", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	es := strconv.Itoa(mockedExitStatus)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1",
		"EXIT_STATUS=" + es}
	return cmd
}

func TestExecCommandHelper(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}

	// println("Mocked stdout:", os.Getenv("STDOUT"))
	fmt.Fprintf(os.Stdout, os.Getenv("STDOUT"))
	i, _ := strconv.Atoi(os.Getenv("EXIT_STATUS"))
	os.Exit(i)
}

func TestExecute(t *testing.T) {
	mockedExitStatus = 0
	SetExecCommand(fakeExecCommand)
	defer ResetExecCommand()

	cmd := &Command{
		Command: "echo hello world",
	}

	result := cmd.Execute()
	if result != nil {
		t.Error(result)
	}
}

func TestExecuteFailure(t *testing.T) {
	mockedExitStatus = 1
	SetExecCommand(fakeExecCommand)
	defer ResetExecCommand()

	cmd := &Command{
		Command: "some error hello world",
	}

	result := cmd.Execute()
	if result == nil {
		t.Errorf("Error should have occured")
	}
}
