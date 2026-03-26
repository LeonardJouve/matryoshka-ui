package main

import (
	"github.com/LeonardJouve/matryoshka-ui/dsl"
	"github.com/LeonardJouve/matryoshka-ui/renderer"
	"github.com/LeonardJouve/matryoshka-ui/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var pink = utils.Color{255, 0, 255}
var red = utils.Color{255, 0, 0}
var blue = utils.Color{0, 0, 255}

func blueDiv() dsl.ElementModifier {
	return dsl.Style(
		dsl.Width(50),
		dsl.Height(50),
		dsl.Color(blue),
	)
}

func main() {
	var width uint16 = 800
	var height uint16 = 450

	renderer := renderer.NewRaylibRenderer()

	element := dsl.Root(dsl.Div(
		dsl.Children(
			dsl.Div(
				dsl.Style(
					dsl.Width(50),
					dsl.Height(50),
					dsl.Color(pink),
				),
			),
			dsl.Div(
				blueDiv(),
			),
		),
		dsl.Style(
			dsl.LayoutAxis(dsl.LAYOUT_HORIZONTAL),
			dsl.Color(red),
			dsl.Gap(
				dsl.GapVertical(10),
				dsl.GapHorizontal(10),
			),
			dsl.Padding(
				dsl.PaddingHorizontal(10),
				dsl.PaddingVertical(10),
			),
		),
	))

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
