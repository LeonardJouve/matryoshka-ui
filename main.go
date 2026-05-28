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
	card := utils.Color{180, 39, 140}
	cardLight := utils.Color{52, 52, 56}
	green := utils.Color{34, 197, 94}
	amber := utils.Color{251, 191, 36}
	red := utils.Color{239, 68, 68}
	pink := utils.Color{236, 72, 153}
	teal := utils.Color{20, 184, 166}

	colors := []utils.Color{accent, green, amber, red, pink, teal}

	makeTag := func(i int, w uint16, h uint16) *Element {
		return Div(Style(
			Width(Fixed(w)),
			Height(Fixed(h)),
			Color(colors[i%len(colors)]),
		))
	}

	makeTags := func(n int, w uint16, h uint16) []*Element {
		tags := make([]*Element, n)
		for i := range tags {
			tags[i] = makeTag(i, w, h)
		}
		return tags
	}

	makeHRow := func(tags []*Element) *Element {
		return Div(
			Style(
				LayoutAxis(LAYOUT_HORIZONTAL),
				Width(Grow(1)),
				Height(Fit()),
				Color(card),
				Padding(PaddingHorizontal(12), PaddingVertical(12)),
				Gap(GapHorizontal(8), GapVertical(8)),
			),
			Children(tags...),
		)
	}

	makeVCol := func(tags []*Element, h uint16) *Element {
		return Div(
			Style(
				LayoutAxis(LAYOUT_VERTICAL),
				Width(Fit()),
				Height(Fixed(h)),
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
			Gap(GapVertical(12)),
		),
		Children(
			// row qui wrappe horizontalement
			makeHRow(makeTags(8, 80, 40)),
			// row avec sidebar fixe + contenu wrappable
			Div(
				Style(
					LayoutAxis(LAYOUT_HORIZONTAL),
					Width(Grow(1)),
					Height(Fit()),
					Color(surface),
					Gap(GapHorizontal(12)),
				),
				Children(
					// sidebar fixe
					makeVCol(makeTags(4, 100, 50), 300),
					// contenu qui wrappe
					makeHRow(makeTags(10, 70, 44)),
				),
			),
			// label
			Div(Style(
				Width(Grow(1)),
				Height(Fixed(4)),
				Color(cardLight),
			)),
			// row qui wrappe avec tags plus grands
			makeHRow(makeTags(6, 120, 60)),
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
			element = buildUI(int32(rl.GetRenderWidth()), int32(rl.GetRenderWidth()))
		}
		rl.BeginDrawing()
		rl.ClearBackground(rl.Color{R: 24, G: 24, B: 27, A: 255})
		renderer.Render(element)

		rl.EndDrawing()
	}
}
