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

	accent := utils.Color{99, 102, 241}  // indigo
	surface := utils.Color{24, 24, 27}   // zinc-900
	card := utils.Color{39, 39, 42}      // zinc-800
	cardLight := utils.Color{52, 52, 56} // zinc-700
	green := utils.Color{34, 197, 94}    // green-500
	amber := utils.Color{251, 191, 36}   // amber-400
	red := utils.Color{239, 68, 68}      // red-500
	white := utils.Color{250, 250, 250}

	_ = white

	pill := func(color utils.Color) *Element {
		return Div(Style(
			Width(Fixed(8)),
			Height(Fixed(8)),
			Color(color),
		))
	}

	menuItem := func(active bool) *Element {
		color := cardLight
		if active {
			color = accent
		}
		return Div(Style(
			LayoutAxis(LAYOUT_HORIZONTAL),
			Width(Grow(1)),
			Height(Fixed(36)),
			Color(color),
			Padding(PaddingHorizontal(12), PaddingVertical(4)),
			Gap(GapHorizontal(8)),
		), Children(
			Div(Style(Width(Fixed(16)), Height(Grow(1)), Color(cardLight))),
			Div(Style(Width(Grow(1)), Height(Fixed(10)), Color(cardLight))),
		))
	}

	statCard := func(color utils.Color) *Element {
		return Div(
			Style(
				LayoutAxis(LAYOUT_VERTICAL),
				Width(Grow(1)),
				Height(Grow(1)),
				Color(card),
				Padding(PaddingHorizontal(16), PaddingVertical(12)),
				Gap(GapVertical(8)),
			),
			Children(
				Div(Style(
					LayoutAxis(LAYOUT_HORIZONTAL),
					Width(Grow(1)),
					Height(Fixed(12)),
					Color(card),
					Gap(GapHorizontal(6)),
				), Children(
					pill(color),
					Div(Style(Width(Fixed(60)), Height(Grow(1)), Color(cardLight))),
				)),
				Div(Style(Width(Fixed(80)), Height(Fixed(20)), Color(white))),
				Div(Style(Width(Fixed(50)), Height(Fixed(8)), Color(cardLight))),
			),
		)
	}

	tableRow := func(statusColor utils.Color) *Element {
		return Div(
			Style(
				LayoutAxis(LAYOUT_HORIZONTAL),
				Width(Grow(1)),
				Height(Fixed(36)),
				Color(card),
				Padding(PaddingHorizontal(12), PaddingVertical(8)),
				Gap(GapHorizontal(12)),
			),
			Children(
				Div(Style(Width(Fixed(24)), Height(Grow(1)), Color(cardLight))),
				Div(Style(Width(Grow(2)), Height(Fixed(10)), Color(cardLight))),
				Div(Style(Width(Grow(1)), Height(Fixed(10)), Color(cardLight))),
				pill(statusColor),
				Div(Style(Width(Fixed(48)), Height(Fixed(10)), Color(cardLight))),
			),
		)
	}

	chartBar := func(heightPct uint16, color utils.Color) *Element {
		barH := uint16(120) * heightPct / 100
		return Div(
			Style(
				LayoutAxis(LAYOUT_VERTICAL),
				Width(Grow(1)),
				Height(Fixed(120)),
				Color(card),
				Gap(GapVertical(0)),
			),
			Children(
				Div(Style(Width(Grow(1)), Height(Fixed(120-barH)), Color(card))),
				Div(Style(Width(Grow(1)), Height(Fixed(barH)), Color(color))),
			),
		)
	}

	return Root(Div(
		Style(
			LayoutAxis(LAYOUT_HORIZONTAL),
			Width(Fixed(w)),
			Height(Fixed(h)),
			Color(surface),
		),
		Children(
			// sidebar
			Div(
				Style(
					LayoutAxis(LAYOUT_VERTICAL),
					Width(Fixed(220)),
					Height(Grow(1)),
					Color(card),
					Padding(PaddingHorizontal(12), PaddingVertical(16)),
					Gap(GapVertical(4)),
				),
				Children(
					// logo
					Div(Style(
						LayoutAxis(LAYOUT_HORIZONTAL),
						Width(Grow(1)),
						Height(Fixed(32)),
						Color(card),
						Padding(PaddingHorizontal(12), PaddingVertical(0)),
						Gap(GapHorizontal(8)),
					), Children(
						Div(Style(Width(Fixed(20)), Height(Fixed(20)), Color(accent))),
						Div(Style(Width(Fixed(80)), Height(Fixed(12)), Color(white))),
					)),
					Div(Style(Width(Grow(1)), Height(Fixed(1)), Color(cardLight))),
					menuItem(true),
					menuItem(false),
					menuItem(false),
					menuItem(false),
					menuItem(false),
					// spacer
					Div(Style(Width(Grow(1)), Height(Grow(1)), Color(card))),
					// user
					Div(Style(
						LayoutAxis(LAYOUT_HORIZONTAL),
						Width(Grow(1)),
						Height(Fixed(40)),
						Color(cardLight),
						Padding(PaddingHorizontal(10), PaddingVertical(6)),
						Gap(GapHorizontal(8)),
					), Children(
						Div(Style(Width(Fixed(24)), Height(Fixed(24)), Color(accent))),
						Div(Style(
							LayoutAxis(LAYOUT_VERTICAL),
							Width(Grow(1)),
							Height(Grow(1)),
							Color(cardLight),
							Gap(GapVertical(4)),
						), Children(
							Div(Style(Width(Fixed(60)), Height(Fixed(8)), Color(white))),
							Div(Style(Width(Fixed(80)), Height(Fixed(8)), Color(cardLight))),
						)),
					)),
				),
			),
			// main
			Div(
				Style(
					LayoutAxis(LAYOUT_VERTICAL),
					Width(Grow(1)),
					Height(Grow(1)),
					Color(surface),
					Padding(PaddingHorizontal(24), PaddingVertical(20)),
					Gap(GapVertical(16)),
				),
				Children(
					// topbar
					Div(
						Style(
							LayoutAxis(LAYOUT_HORIZONTAL),
							Width(Grow(1)),
							Height(Fixed(40)),
							Color(surface),
							Gap(GapHorizontal(12)),
						),
						Children(
							Div(Style(Width(Grow(1)), Height(Grow(1)), Color(card))),
							Div(Style(Width(Fixed(36)), Height(Fixed(36)), Color(card))),
							Div(Style(Width(Fixed(36)), Height(Fixed(36)), Color(accent))),
						),
					),
					// stat cards
					Div(
						Style(
							LayoutAxis(LAYOUT_HORIZONTAL),
							Width(Grow(1)),
							Height(Fixed(100)),
							Color(surface),
							Gap(GapHorizontal(12)),
						),
						Children(
							statCard(green),
							statCard(amber),
							statCard(red),
							statCard(accent),
						),
					),
					// chart + side panel
					Div(
						Style(
							LayoutAxis(LAYOUT_HORIZONTAL),
							Width(Grow(1)),
							Height(Fixed(180)),
							Color(surface),
							Gap(GapHorizontal(12)),
						),
						Children(
							// chart
							Div(
								Style(
									LayoutAxis(LAYOUT_VERTICAL),
									Width(Grow(2)),
									Height(Grow(1)),
									Color(card),
									Padding(PaddingHorizontal(16), PaddingVertical(12)),
									Gap(GapVertical(8)),
								),
								Children(
									Div(Style(Width(Fixed(100)), Height(Fixed(12)), Color(cardLight))),
									Div(
										Style(
											LayoutAxis(LAYOUT_HORIZONTAL),
											Width(Grow(1)),
											Height(Grow(1)),
											Color(card),
											Gap(GapHorizontal(4)),
										),
										Children(
											chartBar(40, accent),
											chartBar(70, accent),
											chartBar(55, accent),
											chartBar(90, accent),
											chartBar(60, accent),
											chartBar(75, accent),
											chartBar(85, accent),
											chartBar(50, accent),
											chartBar(95, accent),
											chartBar(65, accent),
											chartBar(80, accent),
											chartBar(45, accent),
										),
									),
								),
							),
							// side panel
							Div(
								Style(
									LayoutAxis(LAYOUT_VERTICAL),
									Width(Grow(1)),
									Height(Grow(1)),
									Color(card),
									Padding(PaddingHorizontal(16), PaddingVertical(12)),
									Gap(GapVertical(10)),
								),
								Children(
									Div(Style(Width(Fixed(80)), Height(Fixed(12)), Color(cardLight))),
									Div(Style(Width(Grow(1)), Height(Fixed(8)), Color(green))),
									Div(Style(Width(Grow(1)), Height(Fixed(8)), Color(accent))),
									Div(Style(Width(Grow(1)), Height(Fixed(8)), Color(amber))),
									Div(Style(Width(Grow(1)), Height(Fixed(8)), Color(red))),
								),
							),
						),
					),
					// table
					Div(
						Style(
							LayoutAxis(LAYOUT_VERTICAL),
							Width(Grow(1)),
							Height(Grow(1)),
							Color(card),
							Padding(PaddingHorizontal(0), PaddingVertical(0)),
							Gap(GapVertical(1)),
						),
						Children(
							tableRow(green),
							tableRow(amber),
							tableRow(green),
							tableRow(red),
							tableRow(green),
							tableRow(amber),
						),
					),
				),
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
