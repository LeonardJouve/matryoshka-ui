package main

import (
	. "github.com/LeonardJouve/matryoshka-ui/dsl"
	"github.com/LeonardJouve/matryoshka-ui/renderer"
	"github.com/LeonardJouve/matryoshka-ui/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func col(r, g, b uint8) utils.Color { return utils.Color{Red: r, Green: g, Blue: b, Alpha: 255} }

func TestUI(w uint16, h uint16, measurer TextMeasurer) *Node {
	var (
		bg    = col(16, 18, 27)
		card  = col(28, 31, 46)
		track = col(45, 50, 75)
		white = col(235, 238, 248)
		muted = col(140, 146, 175)
		green = col(0, 255, 0)
	)

	iconBtn := func(src string, size uint16) Element {
		return Image(src, ImageStyle(Width(Fixed(size)), Height(Fixed(size))))
	}

	return Root(Div(
		Style(
			Width(Fixed(w)), Height(Fixed(h)),
			LayoutAxis(LAYOUT_VERTICAL),
			Color(bg),
			Justify(JustifyCenter),
		),
		Children(
			Div(
				Style(LayoutAxis(LAYOUT_HORIZONTAL), Width(Grow(1)), Justify(JustifyCenter)),
				Children(
					// Player card
					Div(
						Style(
							LayoutAxis(LAYOUT_VERTICAL),
							Width(Fixed(300)),
							Padding(PaddingAll(24)),
							Gap(GapAll(16)),
							Color(card),
							BorderRadius(0.1),
							BorderColor(utils.Color{255, 0, 0, 255}),
							BorderWidth(1),
						),
						Children(
							// album art — centered, fixed square
							Div(
								Style(LayoutAxis(LAYOUT_HORIZONTAL), Width(Grow(1)), Justify(JustifyCenter)),
								Children(
									Image("assets/album.png", ImageStyle(Width(Fixed(240)), Height(Fixed(240)))),
								),
							),

							// track title + artist
							Text("Midnight City", TextStyle(FontSize(20), Color(white))),
							Text("M83", TextStyle(Color(muted), FontSize(14))),

							// elapsed / duration row
							Div(
								Style(LayoutAxis(LAYOUT_HORIZONTAL), Width(Grow(1)), Gap(GapHorizontal(10)), Justify(JustifyBetween), Align(AlignCenter)),
								Children(
									Text("1:24", TextStyle(Color(muted), FontSize(12))),
									Div(
										Style(Width(Grow(1)), Height(Grow(1)), Color(track), Padding(PaddingAll(4)), BorderRadius(0.8)),
										Children(Div(Style(Width(Fixed(60)), Height(Fixed(8)), Color(green), BorderRadius(0.8)))),
									),
									Text("4:03", TextStyle(Color(muted), FontSize(12))),
								),
							),

							// transport controls row — centered group
							Div(
								Style(
									LayoutAxis(LAYOUT_HORIZONTAL),
									Width(Grow(1)),
									Justify(JustifyBetween),
								),
								Children(
									iconBtn("assets/shuffle.png", 24),
									iconBtn("assets/prev.png", 24),
									iconBtn("assets/play.png", 24),
									iconBtn("assets/next.png", 24),
									iconBtn("assets/repeat.png", 24),
								),
							),
						),
					),
				),
			),
		),
	), measurer)
}

func main() {
	var width uint16 = 800
	var height uint16 = 800
	renderer := renderer.NewRaylibRenderer()
	//renderer := renderer.NewTerminalRenderer()

	renderer.InitWindow(width, height, "Testing101")
	renderer.SetWindowFlag(rl.FlagWindowResizable)
	defer renderer.CloseWindow()

	renderer.Render(TestUI(uint16(rl.GetScreenWidth()), uint16(rl.GetScreenHeight()), renderer))
	//renderer.Render(TestUI(width, height, renderer))
}
