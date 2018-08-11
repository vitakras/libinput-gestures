package gestures

import "testing"

func checkDirection(t *testing.T, expected, actual string) {
	if expected != actual {
		t.Errorf("Expected the Direction to be %s but got %s", expected, actual)
	}
}
