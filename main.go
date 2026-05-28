package main

import (
	. "github.com/LeonardJouve/matryoshka-ui/dsl"
	"github.com/LeonardJouve/matryoshka-ui/renderer"
	"github.com/LeonardJouve/matryoshka-ui/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func buildUI(width, height int32) *RootS {
	w := uint16(width)
	h := uint16(height)

	accent := utils.Color{99, 102, 241}
	surface := utils.Color{24, 24, 27}
	card := utils.Color{150, 39, 150}
	green := utils.Color{34, 197, 94}
	amber := utils.Color{251, 191, 36}
	red := utils.Color{239, 68, 68}
	pink := utils.Color{236, 72, 153}
	teal := utils.Color{20, 184, 166}

	colors := []utils.Color{accent, green, amber, red, pink, teal}

	// row horizontale qui wrap
	makeRow := func(n int, tagH uint16) *Element {
		tags := make([]*Element, n)
		for i := range tags {
			tags[i] = Div(Style(
				Width(Fixed(80)),
				Height(Fixed(tagH)),
				Color(colors[i%len(colors)]),
			))
		}
		return Div(
			Style(
				LayoutAxis(LAYOUT_HORIZONTAL),
				Width(Fit()),
				//Width(Grow(1)),
				Height(Fit()),
				Color(card),
				Padding(PaddingHorizontal(12), PaddingVertical(12)),
				Gap(GapHorizontal(8), GapVertical(8)),
			),
			Children(tags...),
		)
	}

	// colonne verticale qui wrap
	_ = func(n int, tagW uint16) *Element {
		tags := make([]*Element, n)
		for i := range tags {
			tags[i] = Div(Style(
				Width(Fixed(tagW)),
				Height(Fixed(60)),
				Color(colors[(i+2)%len(colors)]),
			))
		}
		return Div(
			Style(
				LayoutAxis(LAYOUT_VERTICAL),
				Width(Fit()),
				Height(Fixed(200)),
				Color(card),
				Padding(PaddingHorizontal(12), PaddingVertical(12)),
				Gap(GapHorizontal(8), GapVertical(8)),
			),
			Children(tags...),
		)
	}

	return Root(Div(
		Style(
			LayoutAxis(LAYOUT_VERTICAL),
			Width(Fixed(w)),
			Height(Fixed(h)),
			Color(surface),
			Padding(PaddingHorizontal(32), PaddingVertical(32)),
			Gap(GapVertical(16)),
		),
		Children(
			makeRow(8, 40),
			//makeCol(6, 80),
			makeRow(10, 50),
		),
	))
}

func main() {
	var width uint16 = 800
	var height uint16 = 800
	renderer := renderer.NewRaylibRenderer()

	var element *RootS

	renderer.InitWindow(int32(width), int32(height), "Testing101")
	renderer.SetWindowFlag(rl.FlagWindowResizable)
	defer renderer.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		if element == nil || rl.IsWindowResized() {
			element = buildUI(int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()))
		}
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		renderer.Render(element)

		rl.EndDrawing()
	}
}
