package gestures

// Pinch represents the Pinch gesture on the trackpad
type Pinch struct {
	diameter float32
}

// In pinch
const In = "In"

// Out Pinch
const Out = "Out"

// NewPinch creates a new Pinch struct
func NewPinch(diameter float32) Pinch {
	return Pinch{
		diameter: diameter,
	}
}

// Direction the direction of the pinch
func (pinch *Pinch) Direction() string {
	if pinch.diameter > 0 {
		return In
	}

	return Out
}
