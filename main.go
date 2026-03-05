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
	blue := utils.Color{0, 0, 255}

	element := elements.NewElement().
		SetLayout(elements.LAYOUT_VERTICAL).
		SetWidth(width).SetHeight(height).
		SetColor(pink).
		SetGap(elements.NewGap(10, 10)).
		SetPadding(elements.NewPadding(10, 10, 10, 10)).
		ForEach(4, func(index uint) *elements.Element {
			return elements.NewElement().
				SetWidth(200).
				SetHeight(200).
				When(index%2 == 0, func(el *elements.Element) elements.IElement {
					return el.SetColor(blue)
				}).
				End()
		}).
		End()
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
