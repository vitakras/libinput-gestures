package libinput

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"testing"
)

var libinputRunResult = "some output"
var mockedExitStatus = 0

func fakeExecCommand(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestExecCommandHelper", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	es := strconv.Itoa(mockedExitStatus)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1",
		"STDOUT=" + libinputRunResult,
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

func TestVersion(t *testing.T) {
	libinputRunResult = "1.3.4"
	mockedExitStatus = 0
	SetExecCommand(fakeExecCommand)
	defer ResetExecCommand()

	out, err := GetVersion()
	if err != nil {
		t.Errorf("Expected nil error, got %#v", err)
	}
	if string(out) != libinputRunResult {
		t.Errorf("Expected %q, got %q", libinputRunResult, out)
	}
}

func TestVersionError(t *testing.T) {
	libinputRunResult = ""
	mockedExitStatus = 1
	SetExecCommand(fakeExecCommand)
	defer ResetExecCommand()

	_, err := GetVersion()
	if err.Error() != "GetVersion: install libinput-tools" {
		t.Errorf("Unexpected error %s", err)
	}
}
