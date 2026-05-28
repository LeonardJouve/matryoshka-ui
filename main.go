package main

import (
	. "github.com/LeonardJouve/matryoshka-ui/dsl"
	"github.com/LeonardJouve/matryoshka-ui/renderer"
	"github.com/LeonardJouve/matryoshka-ui/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var pink = utils.Color{255, 0, 255}
var red = utils.Color{255, 0, 0}
var white = utils.Color{255, 255, 255}
var black = utils.Color{0, 0, 0}
var darkGray = utils.Color{30, 30, 30}
var darkBlue = utils.Color{20, 40, 80}
var blue = utils.Color{50, 100, 200}
var lightBlue = utils.Color{80, 140, 220}
var green = utils.Color{30, 120, 80}

func blueDiv() ElementModifier {
	return Style(
		Width(Fixed(50)),
		Height(Fixed(50)),
		Color(blue),
		Padding(PaddingVertical(5), PaddingHorizontal(10)),
	)
}

func buildUI(width, height int32) *RootS {
	w := uint16(width)
	h := uint16(height)

	accent := utils.Color{99, 102, 241}
	surface := utils.Color{24, 24, 27}
	card := utils.Color{230, 39, 42}
	green := utils.Color{34, 197, 94}
	amber := utils.Color{251, 191, 36}
	red := utils.Color{239, 68, 68}
	pink := utils.Color{236, 72, 153}
	teal := utils.Color{20, 184, 166}

	colors := []utils.Color{accent, green, amber, red, pink, teal}

	tags := make([]*Element, 12)
	for i := range tags {
		tags[i] = Div(Style(
			Width(Fixed(80)),
			Height(Fixed(40)),
			Color(colors[i%len(colors)]),
		))
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
			// container FIT — sa hauteur grandit quand les tags wrappent
			Div(
				Style(
					LayoutAxis(LAYOUT_VERTICAL),
					Width(Fixed(w-100)),
					Height(Fit()),
					Color(card),
					Padding(PaddingHorizontal(12), PaddingVertical(12)),
					Gap(GapHorizontal(8), GapVertical(8)),
				),
				Children(tags...),
			),
		),
	))
}

func main() {
	var width uint16 = 800
	var height uint16 = 800
	renderer := renderer.NewRaylibRenderer()

	var boardSize uint16 = 400
	var cellSize uint16 = boardSize / 8

	makeCell := func(isWhite bool) *Element {
		color := white
		if !isWhite {
			color = black
		}
		return Div(Style(
			Width(Fixed(cellSize)),
			Height(Fixed(cellSize)),
			Color(color),
		))
	}

	makeRow := func(startsWhite bool) *Element {
		cells := make([]*Element, 8)
		for col := 0; col < 8; col++ {
			isWhite := (col%2 == 0) == startsWhite
			cells[col] = makeCell(isWhite)
		}
		return Div(
			Children(cells...),
			Style(
				LayoutAxis(LAYOUT_HORIZONTAL),
				Width(Fixed(boardSize)),
				Height(Fixed(cellSize)),
			),
		)
	}

	rows := make([]*Element, 8)
	for row := 0; row < 8; row++ {
		rows[row] = makeRow(row%2 == 0)
	}

	var element *RootS

	//element := Root(Div(
	//	Children(rows...),
	//	Style(
	//		LayoutAxis(LAYOUT_VERTICAL),
	//		Color(red),
	//		Width(Fixed(width)),
	//		Height(Fixed(height)),
	//		Padding(
	//			PaddingHorizontal((width-boardSize)/2),
	//			PaddingVertical((height-boardSize)/2),
	//		),
	//	),
	//))

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
