package libinput

import (
	"regexp"
	"strings"
)

// Device Input Device
type Device struct {
	ID        string
	Name      string
	Available bool
}

// SetName parses input and sets Name if possible, returns true if name gets set
func (device *Device) SetName(line string) bool {
	exp := regexp.MustCompile("Device:\\s+([\\s\\S]*)")
	result := exp.FindStringSubmatch(line)
	if len(result) > 0 {
		device.Name = result[1]
		return true
	}

	return false
}

// SetID parses input and sets ID if possible, returns true if ID gets set
func (device *Device) SetID(line string) bool {
	if strings.HasPrefix(line, "Kernel:") {
		exp := regexp.MustCompile("(event\\d*)")
		device.ID = exp.FindString(line)
		return true
	}

	return false
}

// SetAvailable parses input and sets Availability of the device, returns true if Availability gets set
func (device *Device) SetAvailable(line string) bool {
	if strings.HasPrefix(line, "Nat.scrolling:") {
		exp := regexp.MustCompile("n\\/a")
		device.Available = !exp.MatchString(line)
		return true
	}

	return false
}
