package main

import (
	"example.com/backend"
)

func main() {

	a := backend.App{}
	a.Port = ":9003"
	a.Initialize()

	a.Run()
}

func calculateDimensions(length, width, height int) (area int, volume int) {
	area = length * width * height
	volume = length * height
	return
}
