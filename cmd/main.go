package main

import (
	"fmt"
)

func main() {
	config, _ := LoadConfigFromFile("config.json")
	fmt.Print(config.Interval.Swipe)
}
