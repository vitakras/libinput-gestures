package main

import "github.com/vitakras/libinput-gestures/pkg/libinput"

// SupportedActions list of supported libinput events
var SupportedActions = []string{
	"GESTURE_SWIPE_UPDATE",
	"GESTURE_PINCH_UPDATE",
}

// IsValidDebugEvent return true of false whether the debug event is valid
func IsValidDebugEvent(debugEvent *libinput.DebugEvent) bool {
	for _, actionType := range SupportedActions {
		if actionType == debugEvent.Action {
			return true
		}
	}

	return false
}
