package app

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// Config represents the config.yaml file
type Config struct {
	Swipe     map[int]SwipeConfig `yaml:"swipe"`
	Pinch     PinchConfig         `yaml:"pinch"`
	Threshold Threshold           `yaml:"threshold"`
	Interval  Interval            `yaml:"interval"`
}

// Threshold represents the threshold in the config.yaml
type Threshold struct {
	Swipe float32 `yaml:"swipe"`
	Pinch float32 `yaml:"pinch"`
}

// Interval represents the interval in the config.yaml
type Interval struct {
	Swipe float32 `yaml:"swipe"`
	Pinch float32 `yaml:"pinch"`
}

// PinchConfig represents the pinch in the config.yaml
type PinchConfig struct {
	In  Command `yaml:"in"`
	Out Command `yaml:"out"`
}

// SwipeConfig represents the swipe in the config.yaml
type SwipeConfig struct {
	Left  Command `yaml:"left"`
	Right Command `yaml:"right"`
	Up    Command `yaml:"up"`
	Down  Command `yaml:"down"`
}

// Gesture interface for a gesture
//go:generate counterfeiter . Gesture
type Gesture interface {
	IsComplete() bool
	Direction() string
	Fingers() int
	Type() string
}

// LoadConfigFromFile Loads the Configuration from file
func LoadConfigFromFile(filePath string) (*Config, error) {
	yamlFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer yamlFile.Close()

	byteValue, _ := ioutil.ReadAll(yamlFile)

	config := &Config{}
	err = yaml.Unmarshal(byteValue, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// GetCommand returns a command to from the config
func (config *Config) GetCommand(gesture Gesture) *Command {
	if gesture.IsComplete() {
		switch gesture.Type() {
		case "swipe":
			return config.getSwipeCommand(gesture.Fingers(), gesture.Direction())
		case "pinch":
			return config.getPinchCommand(gesture.Direction())
		default:
			return nil
		}
	}

	return nil
}

func (config *Config) getSwipeCommand(fingers int, direction string) *Command {
	var command *Command
	swipe, ok := config.Swipe[fingers]
	if ok {
		switch direction {
		case "left":
			command = &swipe.Left
		case "right":
			command = &swipe.Right
		case "up":
			command = &swipe.Up
		case "down":
			command = &swipe.Down
		}
	}

	return command
}

func (config *Config) getPinchCommand(direction string) *Command {
	var command *Command
	pinch := config.Pinch

	switch direction {
	case "in":
		command = &pinch.In
	case "out":
		command = &pinch.Out
	}

	return command
}
