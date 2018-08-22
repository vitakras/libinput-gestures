package gestures

import "math"

// Swipe represents the swipe gesture on the trackpad
type Swipe struct {
	x float32
	y float32
}

// Up swipe
const Up = "Up"

// Down swipe
const Down = "Down"

// Left swipe
const Left = "Left"

// Right swipe
const Right = "Right"

// None no direction
const None = "None"

// NewSwipe creates a new Swipe struct
func NewSwipe(x, y float32) Swipe {
	return Swipe{
		x: x,
		y: y,
	}
}

// Direction of the swipe
func (swipe *Swipe) Direction() string {
	var direction string

	if math.Abs(float64(swipe.x)) > math.Abs(float64(swipe.y)) {
		if swipe.x > 0 {
			direction = Right
		} else {
			direction = Left
		}
	} else {
		if swipe.x == swipe.y && swipe.x == 0 {
			direction = None
		} else if swipe.y > 0 {
			direction = Down
		} else {
			direction = Up
		}
	}

	return direction
}
