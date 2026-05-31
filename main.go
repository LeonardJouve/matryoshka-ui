package main

import (
	. "github.com/LeonardJouve/matryoshka-ui/dsl"
	"github.com/LeonardJouve/matryoshka-ui/renderer"
	"github.com/LeonardJouve/matryoshka-ui/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func col(r, g, b uint8) utils.Color { return utils.Color{Red: r, Green: g, Blue: b} }

func TestUI(w uint16, h uint16) *Node {
	box := func(w, h uint16, c utils.Color) Element {
		return Div(Style(
			Width(Fixed(w)),
			Height(Fixed(h)),
			Color(c),
		))
	}

	return Root(Div(
		Style(
			LayoutAxis(LAYOUT_VERTICAL),
			Padding(PaddingAll(16)),
			Gap(GapAll(12)),
			Color(col(20, 24, 33)),
		),
		Children(
			// ROW 1 — fit container hugging three fixed boxes
			Div(
				Style(
					LayoutAxis(LAYOUT_HORIZONTAL),
					Padding(PaddingAll(8)),
					Gap(GapAll(8)),
					Color(col(40, 50, 70)),
				),
				Children(
					box(60, 40, col(99, 132, 240)),
					box(60, 40, col(46, 196, 132)),
					box(60, 40, col(240, 176, 64)),
				),
			),

			// ROW 2 — grow 1:2:1 across full width
			Div(
				Style(
					LayoutAxis(LAYOUT_HORIZONTAL),
					Width(Grow(1)),
					Height(Fixed(40)),
					Gap(GapAll(8)),
					Color(col(40, 50, 70)),
				),
				Children(
					Div(Style(Width(Grow(1)), Height(Fixed(28)), Color(col(99, 132, 240)))),
					Div(Style(Width(Grow(2)), Height(Fixed(28)), Color(col(46, 196, 132)))),
					Div(Style(Width(Grow(1)), Height(Fixed(28)), Color(col(240, 176, 64)))),
				),
			),

			// ROW 3 — nested mixed axes: fixed column beside a grow filler
			Div(
				Style(
					LayoutAxis(LAYOUT_HORIZONTAL),
					Width(Grow(1)),
					Height(Fixed(120)),
					Gap(GapAll(8)),
					Color(col(40, 50, 70)),
				),
				Children(
					Div(
						Style(
							LayoutAxis(LAYOUT_VERTICAL),
							Width(Fixed(80)),
							Gap(GapAll(6)),
							Color(col(60, 40, 70)),
						),
						Children(
							box(60, 30, col(232, 92, 110)),
							box(60, 30, col(232, 92, 110)),
						),
					),
					Div(Style(
						Width(Grow(1)),
						Color(col(30, 60, 60)),
					)),
				),
			),
		),
	))
}

func main() {
	var width uint16 = 800
	var height uint16 = 800
	renderer := renderer.NewRaylibRenderer()

	renderer.InitWindow(int32(width), int32(height), "Testing101")
	renderer.SetWindowFlag(rl.FlagWindowResizable)
	defer renderer.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Color{R: 24, G: 24, B: 27, A: 255})
		renderer.Render(TestUI(uint16(rl.GetScreenWidth()), uint16(rl.GetScreenHeight())))

		rl.EndDrawing()
	}
}
