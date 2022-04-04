package main

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/go-vgo/robotgo"
)

func main() {
	// get position of mouse cursor
	x, y := robotgo.GetMousePos()

	// get color of pixel at that position
	color := robotgo.GetPixelColor(x, y)

	col := fmt.Sprintf("#%s", color)
	clipboard.WriteAll(col)
}
