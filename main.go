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

	renderer := renderer.NewRaylibRenderer()

	pink := utils.Color{255, 0, 255}
	element := elements.NewElement().
		Style(elements.NewStyle().
			LayoutAxis(elements.LAYOUT_HORIZONTAL).
			Width(width).Height(height).
			Color(pink).
			Gap(elements.NewGap(10, 10)).
			Padding(elements.NewPadding(10, 10, 10, 10))).
		Childrens(
			elements.NewElement().Style(elements.NewStyle().Width(200).Height(200)).End(),
			elements.NewElement().Style(elements.NewStyle().Width(200).Height(200)).End(),
			elements.NewElement().Style(elements.NewStyle().Width(200).Height(200)).End(),
			elements.NewElement().Style(elements.NewStyle().Width(200).Height(200)).End(),
		).
		End()

	renderer.InitWindow(int32(width), int32(height), "MaBite")
	renderer.SetWindowFlag(rl.FlagWindowResizable)
	defer renderer.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		renderer.Render(element)

		rl.EndDrawing()
	}
}
