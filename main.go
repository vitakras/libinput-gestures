package main

import (
	"fmt"

	"github.com/vitakras/libinput-gestures/pkg/libinput"
)

func main() {
	stream := libinput.NewDebugEventStream()
	err := stream.Start()

	if err != nil {
		fmt.Print(err)
	}

	for {
		event := stream.Read()
		if event != nil {
			fmt.Printf("Event is %s %f\n ", event.Action, event.Time)
		}
	}
}
