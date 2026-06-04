package main

import (
	"github.com/LeonardJouve/matryoshka-ui/example"
	"github.com/LeonardJouve/matryoshka-ui/renderer"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	var width uint16 = 800
	var height uint16 = 800
	renderer := renderer.NewRaylibRenderer()
	//renderer := renderer.NewTerminalRenderer()

	renderer.InitWindow(width, height, "Test")
	renderer.SetWindowFlag(rl.FlagWindowResizable | rl.FlagWindowHighdpi)
	defer renderer.CloseWindow()

	renderer.Render(example.Percentage())
}
