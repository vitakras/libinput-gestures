package app

import "github.com/vitakras/libinput-gestures/pkg/libinput"

// DebugStreamer used to read in the DebugEvents
//go:generate counterfeiter . DebugStreamer
type DebugStreamer interface {
	Start() error
	Read() *libinput.DebugEvent
	Closed() bool
}

// Processor used to Process DebugEvents
//go:generate counterfeiter . Processor
type Processor interface {
	Process(debugEvent *libinput.DebugEvent) error
}

// LibinputGestures used to read input from Stream and Process the streams
type LibinputGestures struct {
	debugStreamer DebugStreamer
	processor     Processor
}

// NewLibinputGestures returns a new LibinputGestures struct
func NewLibinputGestures(processor Processor, debugStreamer DebugStreamer) *LibinputGestures {
	return &LibinputGestures{
		debugStreamer: debugStreamer,
		processor:     processor,
	}
}

// ProcessEvents reads and processes the DebugEvent structs
func (lg *LibinputGestures) ProcessEvents() error {
	stream := lg.debugStreamer
	err := stream.Start()
	if err != nil {
		return err
	}

	processor := lg.processor
	for lg.debugStreamer.Closed() == false {
		event := stream.Read()

		err := processor.Process(event)
		if err != nil {
			return err
		}
	}

	return nil
}
