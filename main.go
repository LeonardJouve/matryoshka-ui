package main

import (
	. "github.com/LeonardJouve/matryoshka-ui/dsl"
	"github.com/LeonardJouve/matryoshka-ui/renderer"
	"github.com/LeonardJouve/matryoshka-ui/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func col(r, g, b uint8) utils.Color { return utils.Color{Red: r, Green: g, Blue: b} }

func TestUI(w uint16, h uint16, measurer TextMeasurer) *Node {
	var (
		bg    = col(16, 18, 27)
		card  = col(28, 31, 46)
		track = col(45, 50, 75)
		white = col(235, 238, 248)
		muted = col(140, 146, 175)
	)

	iconBtn := func(src string, size uint16) Element {
		return Image(src, ImageStyle(Width(Fixed(size)), Height(Fixed(size))))
	}

	return Root(Div(
		Style(
			Width(Fixed(w)), Height(Fixed(h)),
			LayoutAxis(LAYOUT_VERTICAL),
			Color(bg),
		),
		Children(
			Div(Style(Height(Grow(1)))), // top spacer (center vertically)

			Div(
				Style(LayoutAxis(LAYOUT_HORIZONTAL), Width(Grow(1)), Height(Fixed(420))),
				Children(
					Div(Style(Width(Grow(1)))), // left spacer

					// PLAYER CARD
					Div(
						Style(
							LayoutAxis(LAYOUT_VERTICAL),
							Width(Fixed(300)), Height(Grow(1)),
							Padding(PaddingAll(24)),
							Gap(GapAll(16)),
							Color(card),
						),
						Children(
							// album art — centered, fixed square
							Div(
								Style(LayoutAxis(LAYOUT_HORIZONTAL), Width(Grow(1)), Height(Fixed(240))),
								Children(
									Div(Style(Width(Grow(1)))),
									Image("assets/album.png", ImageStyle(Width(Fixed(240)), Height(Fixed(240)))),
									Div(Style(Width(Grow(1)))),
								),
							),

							// track title + artist
							Text("Midnight City", TextColor(white), FontSize(20)),
							Text("M83", TextColor(muted), FontSize(14)),

							// progress track (full width bar)
							Div(Style(Width(Grow(1)), Height(Fixed(4)), Color(track))),

							// elapsed / duration row
							Div(
								Style(LayoutAxis(LAYOUT_HORIZONTAL), Width(Grow(1)), Height(Fixed(14))),
								Children(
									Text("1:24", TextColor(muted), FontSize(12)),
									Div(Style(Width(Grow(1)))), // spacer pushes duration right
									Text("4:03", TextColor(muted), FontSize(12)),
								),
							),

							// transport controls row — centered group
							Div(
								Style(
									LayoutAxis(LAYOUT_HORIZONTAL),
									Width(Grow(1)), Height(Fixed(48)),
									Gap(GapAll(16)),
								),
								Children(
									Div(Style(Width(Grow(1)))), // left spacer
									iconBtn("assets/shuffle.png", 24),
									iconBtn("assets/prev.png", 36),
									iconBtn("assets/play.png", 48), // bigger play
									iconBtn("assets/next.png", 36),
									iconBtn("assets/repeat.png", 24),
									Div(Style(Width(Grow(1)))), // right spacer
								),
							),
						),
					),

					Div(Style(Width(Grow(1)))), // right spacer
				),
			),

			Div(Style(Height(Grow(1)))), // bottom spacer
		),
	), measurer)
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
		renderer.Render(TestUI(uint16(rl.GetScreenWidth()), uint16(rl.GetScreenHeight()), renderer))

		rl.EndDrawing()
	}
}
