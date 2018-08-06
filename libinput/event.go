package libinput

import (
	"fmt"
	"regexp"
	"strconv"
)

// DebugEvent represents the libinput debug event
type DebugEvent struct {
	Device    string
	Action    string
	Time      float32
	Fingers   int
	Direction *FingerDirection
}

// FingerDirection represents finger movement in the libinput debug event
type FingerDirection struct {
	X    float32
	Y    float32
	Zoom float32
}

// ParseDebugLine parses debug events from libinput and returns the DebugEvent struct
func ParseDebugLine(line string) (*DebugEvent, error) {
	exp := regexp.MustCompile("\\s+")
	rawEvent := exp.Split(line, 5)

	debugEvent := &DebugEvent{
		Device: rawEvent[0],
		Action: rawEvent[1],
	}

	time, err := getTime(rawEvent[2])
	if err != nil {
		return nil, err
	}

	debugEvent.Time = time
	debugEvent.Fingers, err = strconv.Atoi(rawEvent[3])

	if len(rawEvent) == 5 {
		fingerDrirection, err := parseFingerDirection(rawEvent[4])
		if err != nil {
			return nil, err
		}

		debugEvent.Direction = fingerDrirection
	}

	return debugEvent, nil
}

func parseFingerDirection(line string) (*FingerDirection, error) {
	exp := regexp.MustCompile("(\\-?\\d+\\.\\d+)/(\\-?\\d+\\.\\d+)([\\s\\S]*)")

	result := exp.FindStringSubmatch(line)
	if result == nil {
		return nil, fmt.Errorf("parseFingerDirection: Failed to find Direction in string %s", line)
	}

	x, err := strconv.ParseFloat(result[1], 32)
	if err != nil {
		return nil, err
	}

	y, err := strconv.ParseFloat(result[2], 32)
	if err != nil {
		return nil, err
	}

	fingerDirection := &FingerDirection{
		X: float32(x),
		Y: float32(y),
	}

	if len(result) == 4 {
		zoom, err := parseZoom(result[3])
		if err != nil {
			return nil, err
		}

		fingerDirection.Zoom = zoom
	}

	return fingerDirection, nil
}

func parseZoom(line string) (float32, error) {
	exp := regexp.MustCompile("\\@\\s+(\\-?\\d+\\.\\d+)")

	result := exp.FindStringSubmatch(line)
	if result == nil {
		return 0, nil
	}

	zoom, err := strconv.ParseFloat(result[1], 32)
	if err != nil {
		return 0, err
	}

	return float32(zoom), nil
}

func getTime(rawTime string) (float32, error) {
	exp := regexp.MustCompile("([0-9]+\\.[0-9]+)")
	parsedRawTime := exp.FindString(rawTime)

	time, err := strconv.ParseFloat(parsedRawTime, 32)
	if err != nil {
		return 0.0, err
	}

	return float32(time), nil
}
