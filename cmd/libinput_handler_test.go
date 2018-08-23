package main

import (
	"testing"

	"github.com/vitakras/libinput-gestures/pkg/libinput"
)

func TestIsValidDebugEvent(t *testing.T) {
	testcases := []struct {
		input          *libinput.DebugEvent
		expectedResult bool
	}{
		{
			input: &libinput.DebugEvent{
				Action: "GESTURE_SWIPE_UPDATE",
			},
			expectedResult: true,
		},
		{
			input: &libinput.DebugEvent{
				Action: "GESTURE_PINCH_UPDATE",
			},
			expectedResult: true,
		},
		{
			input: &libinput.DebugEvent{
				Action: "POINTER_MOTION",
			},
			expectedResult: false,
		},
	}

	for _, tc := range testcases {
		actualResult := IsValidDebugEvent(tc.input)
		if actualResult != tc.expectedResult {
			t.Errorf("Expected Action %s to return %t, actual %t", tc.input.Action, actualResult, tc.expectedResult)
		}
	}
}
