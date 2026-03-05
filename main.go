package main

import (
	"github.com/LeonardJouve/matryoshka-ui/elements"
	"github.com/LeonardJouve/matryoshka-ui/renderer"
	"github.com/LeonardJouve/matryoshka-ui/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	var width uint16 = 800
	var height uint16 = 450
	rl.InitWindow(int32(width), int32(height), "matryoshka-ui")

	pink := utils.Color{255, 0, 255}
	green := utils.Color{0, 255, 0}

	element := elements.NewElement().SetLayout(elements.LAYOUT_VERTICAL).SetWidth(width).SetHeight(height).SetColor(pink).SetGap(elements.NewGap(10, 10)).SetPadding(elements.NewPadding(10, 10, 10, 10)).AddChild(
		elements.NewElement().SetWidth(200).SetHeight(200).End(),
		elements.NewElement().SetWidth(200).SetHeight(200).SetColor(green).End(),
		elements.NewElement().SetWidth(200).SetHeight(200).End(),
		elements.NewElement().SetWidth(200).SetHeight(200).SetColor(green).End(),
	).End()
	renderer := renderer.RaylibRenderer{}
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	rl.SetWindowState(rl.FlagWindowResizable)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		renderer.Render(element)

		rl.EndDrawing()
	}
}
