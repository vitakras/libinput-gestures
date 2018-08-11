package libinput

import "testing"

func TestSetNameGetsSet(t *testing.T) {
	const validLine = "Device:           SynPS/2 Synaptics TouchPad"
	device := &Device{}

	isNameSet := device.SetName(validLine)
	if isNameSet != true {
		t.Error("Expected device.SetName to return true but got false")
	}

	const expectedName = "SynPS/2 Synaptics TouchPad"
	if device.Name != expectedName {
		t.Errorf("Expected device.Name to be '%s' but got '%s'", expectedName, device.Name)
	}
}

func TestSetNameInvalidInputNameDoesNotGetSet(t *testing.T) {
	const validLine = "InvalidName:           blah"
	device := &Device{}

	isNameSet := device.SetName(validLine)
	if isNameSet != false {
		t.Error("Expected device.SetName to return false but got true")
	}

	if device.Name != "" {
		t.Errorf("Expected device.Name to be empty but got '%s'", device.Name)
	}
}

func TestSetIDGetsSet(t *testing.T) {
	const validLine = "Kernel:           /dev/input/event6"
	device := &Device{}

	isIDSet := device.SetID(validLine)
	if isIDSet != true {
		t.Error("Expected device.SetID to return true but got false")
	}

	const expectedID = "event6"
	if device.ID != expectedID {
		t.Errorf("Expected device.ID to be '%s' but got '%s'", expectedID, device.ID)
	}
}

func TestSetIDInvalidInputIDDoesNotGetSet(t *testing.T) {
	const validLine = "InvalidID:           blah"
	device := &Device{}

	isIDSet := device.SetID(validLine)
	if isIDSet != false {
		t.Error("Expected device.SetID to return false but got true")
	}

	if device.ID != "" {
		t.Errorf("Expected device.ID to be empty but got '%s'", device.ID)
	}
}

func TestSetAvailableGetsSetTrue(t *testing.T) {
	const validLine = "Nat.scrolling:    disabled"
	device := &Device{}

	isAvailableSet := device.SetAvailable(validLine)
	if isAvailableSet != true {
		t.Error("Expected device.SetAvailable to return true but got false")
	}

	const expectedAvailable = true
	if device.Available != expectedAvailable {
		t.Errorf("Expected device.ID to be '%t' but got '%t'", expectedAvailable, device.Available)
	}
}

func TestSetAvailableGetsSetFalse(t *testing.T) {
	const validLine = "Nat.scrolling:    n/a"
	device := &Device{
		Available: true,
	}

	isAvailableSet := device.SetAvailable(validLine)
	if isAvailableSet != true {
		t.Error("Expected device.SetAvailable to return true but got false")
	}

	const expectedAvailable = false
	if device.Available != expectedAvailable {
		t.Errorf("Expected device.ID to be '%t' but got '%t'", expectedAvailable, device.Available)
	}
}

func TestSetAvailableInvalidInput(t *testing.T) {
	const validLine = "InvalidStuf:           blah"
	device := &Device{}

	isAvailableSet := device.SetAvailable(validLine)
	if isAvailableSet != false {
		t.Error("Expected device.SetAvailable to return false but got true")
	}

	if device.Available != false {
		t.Errorf("Expected device.Available to be false is true")
	}
}
