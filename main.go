package main

import (
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	for {
		x, y := robotgo.GetMousePos()
		robotgo.MoveMouse(x+3, y)
		time.Sleep(500 * time.Millisecond)
		robotgo.MoveMouse(x+3, y+3)
		time.Sleep(500 * time.Millisecond)
		robotgo.MoveMouse(x, y+3)
		time.Sleep(500 * time.Millisecond)
		robotgo.MoveMouse(x, y)
		time.Sleep(3 * time.Second)
	}
}
