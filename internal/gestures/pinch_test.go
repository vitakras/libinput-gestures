package gestures

import "testing"

func TestPinchDirection(t *testing.T) {
	var pinch Pinch
	var direction string

	pinch = NewPinch(1.4)
	direction = pinch.Direction()
	checkDirection(t, In, direction)

	pinch = NewPinch(-1.4)
	direction = pinch.Direction()
	checkDirection(t, Out, direction)
}
