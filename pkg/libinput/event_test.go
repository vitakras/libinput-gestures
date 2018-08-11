package libinput

import "testing"

func TestPraseLine(t *testing.T) {
	const line = "event11  GESTURE_SWIPE_BEGIN  +1.10s	3"
	expectedEvent := &DebugEvent{
		Device:  "event11",
		Action:  "GESTURE_SWIPE_BEGIN",
		Time:    1.10,
		Fingers: 3,
	}

	actualEvent, err := ParseDebugLine(line)
	if err != nil {
		t.Errorf("Got error %s", err)
	} else {
		compareDebugEvent(t, expectedEvent, actualEvent)
	}
}

func TestParseLineSwipe(t *testing.T) {
	const line = "event11  GESTURE_SWIPE_UPDATE +201.63s	3 -2.03/16.63 (-5.47/44.84 unaccelerated)"
	expectedEvent := &DebugEvent{
		Device:  "event11",
		Action:  "GESTURE_SWIPE_UPDATE",
		Time:    201.63,
		Fingers: 3,
		Direction: &FingerDirection{
			X: -2.03,
			Y: 16.63,
		},
	}

	actualEvent, err := ParseDebugLine(line)
	if err != nil {
		t.Errorf("Got error %s", err)
	} else {
		compareDebugEvent(t, expectedEvent, actualEvent)
	}
}

func TestParseLinePinch(t *testing.T) {
	const line = "event11  GESTURE_PINCH_UPDATE  +1.85s	3  0.00/-0.41 ( 0.00/-1.64 unaccelerated)  0.86 @  0.08"
	expectedEvent := &DebugEvent{
		Device:  "event11",
		Action:  "GESTURE_PINCH_UPDATE",
		Time:    1.85,
		Fingers: 3,
		Direction: &FingerDirection{
			X:    0.0,
			Y:    -0.41,
			Zoom: 0.08,
		},
	}

	actualEvent, err := ParseDebugLine(line)
	if err != nil {
		t.Errorf("Got error %s", err)
	} else {
		compareDebugEvent(t, expectedEvent, actualEvent)
	}
}

func compareDebugEvent(t *testing.T, expected, actual *DebugEvent) {
	if expected.Device != actual.Device {
		t.Errorf("Expected debugEvent.Device to equal '%s' but got '%s'", expected.Device, actual.Device)
	}
	if expected.Action != actual.Action {
		t.Errorf("Expected debugEvent.Action to equal '%s' but got '%s'", expected.Action, actual.Action)
	}
	if expected.Time != actual.Time {
		t.Errorf("Expected debugEvent.Time to equal '%f' but got '%f'", expected.Time, actual.Time)
	}
	if expected.Fingers != actual.Fingers {
		t.Errorf("Expected debugEvent.Fingers to equal '%d' but got '%d'", expected.Fingers, actual.Fingers)
	}

	if expected.Direction != actual.Direction {
		compareFingerDirection(t, expected.Direction, actual.Direction)
	}
}

func compareFingerDirection(t *testing.T, expected, actual *FingerDirection) {
	if expected.X != actual.X {
		t.Errorf("Expected FingerDirection.X to equal '%f' but got '%f'", expected.X, actual.X)
	}
	if expected.Y != actual.Y {
		t.Errorf("Expected FingerDirection.Y to equal '%f' but got '%f'", expected.Y, actual.Y)
	}
	if expected.Zoom != actual.Zoom {
		t.Errorf("Expected FingerDirection.Zoom to equal '%f' but got '%f'", expected.Zoom, actual.Zoom)
	}
}
