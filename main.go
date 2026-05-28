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
var white = utils.Color{255, 255, 255}
var black = utils.Color{0, 0, 0}

func blueDiv() ElementModifier {
	return Style(
		Width(Fixed(50)),
		Height(Fixed(50)),
		Color(blue),
		Padding(PaddingVertical(5), PaddingHorizontal(10)),
	)
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

	element := Root(Div(
		Children(rows...),
		Style(
			LayoutAxis(LAYOUT_VERTICAL),
			Color(red),
			Width(Fixed(width)),
			Height(Fixed(height)),
			Padding(
				PaddingHorizontal((width-boardSize)/2),
				PaddingVertical((height-boardSize)/2),
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
