package main

import "github.com/vitakras/libinput-gestures/pkg/libinput"

// SupportedActionsMap list of supported libinput events
var SupportedActionsMap = map[string]bool{
	"GESTURE_SWIPE_START":  true,
	"GESTURE_SWIPE_UPDATE": true,
	"GESTURE_PINCH_START":  true,
	"GESTURE_PINCH_UPDATE": true,
}

// IsValidDebugEvent return true of false whether the debug event is valid
func IsValidDebugEvent(debugEvent *libinput.DebugEvent) bool {
	_, supported := SupportedActionsMap[debugEvent.Action]
	return supported
}
