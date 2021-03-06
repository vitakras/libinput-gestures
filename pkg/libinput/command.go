package libinput

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strings"
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
	cmd := libinput("--version")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("GetVersion: install libinput-tools")
	}

	return string(output), nil
}

// ListDevices returns a list of devices from list-devices command
func ListDevices() ([]*Device, error) {
	cmd := libinput("list-devices")
	outPipe, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	err = cmd.Start()
	if err != nil {
		return nil, err
	}

	devices := readDevices(outPipe)

	err = cmd.Wait()
	if err != nil {
		return nil, err
	}

	return devices, nil
}

// DebugEventStream interface for controlling the output from debug-events command
type DebugEventStream struct {
	cmd     *exec.Cmd
	scanner *bufio.Scanner
}

// NewDebugEventStream creates a new DebugEventStream
func NewDebugEventStream() *DebugEventStream {
	cmd := libinput("debug-events")

	reader, err := cmd.StdoutPipe()
	if err != nil {
		return nil
	}

	scanner := bufio.NewScanner(reader)
	return &DebugEventStream{
		cmd:     cmd,
		scanner: scanner,
	}
}

// Start stars the DebugEventSteam
func (stream *DebugEventStream) Start() error {
	return stream.cmd.Start()
}

// Stop stops the DebugEventStream
func (stream *DebugEventStream) Stop() error {
	return stream.cmd.Process.Kill()
}

// Read reads the DebugEventSteam and returns
func (stream *DebugEventStream) Read() *DebugEvent {
	scanner := stream.scanner
	if scanner.Scan() {
		text := scanner.Text()
		if len(text) != 0 {
			text = strings.Trim(text, " ")
			event, err := ParseDebugLine(text)
			if err != nil {
				return event
			}
			return event
		}
	}

	return nil
}

func libinput(args ...string) *exec.Cmd {
	return execCommand("libinput", args...)
}

func readDevices(reader io.Reader) []*Device {
	devices := make([]*Device, 0)
	var device *Device
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			if device != nil {
				devices = append(devices, device)
				device = nil
			}
		} else {
			if device == nil {
				device = &Device{}
			}
			populateDevice(device, line)
		}
	}

	return devices
}

func populateDevice(device *Device, line string) {
	if device.SetName(line) {
		return
	}

	if device.SetID(line) {
		return
	}

	if device.SetAvailable(line) {
		return
	}
}
