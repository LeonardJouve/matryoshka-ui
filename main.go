package main

import (
	. "github.com/LeonardJouve/matryoshka-ui/dsl"
	"github.com/LeonardJouve/matryoshka-ui/renderer"
	"github.com/LeonardJouve/matryoshka-ui/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var pink = utils.Color{255, 0, 255}
var red = utils.Color{255, 0, 0}
var blue = utils.Color{0, 0, 255}

func blueDiv() ElementModifier {
	return Style(
		Width(Fixed(400)),
		Height(Fixed(50)),
		Color(blue),
		Padding(PaddingVertical(5), PaddingHorizontal(10)),
	)
}

func main() {
	var width uint16 = 800
	var height uint16 = 450

	renderer := renderer.NewRaylibRenderer()

	element := Root(Div(
		Children(
			Div(
				Style(
					Width(Grow(1)),
					Height(Fixed(50)),
					Color(pink),
				),
			),
			Div(
				blueDiv(),
			),
			Div(
				blueDiv(),
			),
			Div(
				Style(
					Width(Grow(2)),
					Height(Fixed(50)),
					Color(pink),
				),
			),
		),
		Style(
			LayoutAxis(LAYOUT_HORIZONTAL),
			Color(red),
			Width(Fixed(width)),
			Height(Fixed(200)),
			Gap(
				GapVertical(10),
				GapHorizontal(10),
			),
			Padding(
				PaddingHorizontal(10),
				PaddingVertical(10),
			),
		),
	))

	renderer.InitWindow(int32(width), int32(height), "Testing101")
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
