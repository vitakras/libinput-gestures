package gestures

import "testing"

func TestSwipeDirection(t *testing.T) {
	var swipe Swipe
	var direction string

	// Test None
	swipe = NewSwipe(0.0, 0.0)
	direction = swipe.Direction()
	checkDirection(t, None, direction)

	// Test Left
	swipe = NewSwipe(-4.0, 0.0)
	direction = swipe.Direction()
	checkDirection(t, Left, direction)

	// Test Right
	swipe = NewSwipe(4.0, 0.0)
	direction = swipe.Direction()
	checkDirection(t, Right, direction)

	// Test Up
	swipe = NewSwipe(0.0, -3.0)
	direction = swipe.Direction()
	checkDirection(t, Up, direction)

	// Test Up
	swipe = NewSwipe(0.0, 3.0)
	direction = swipe.Direction()
	checkDirection(t, Down, direction)
}
