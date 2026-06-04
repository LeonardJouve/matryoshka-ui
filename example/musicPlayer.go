package example

import (
	. "github.com/LeonardJouve/matryoshka-ui/dsl"
	"github.com/LeonardJouve/matryoshka-ui/utils"
)

func MusicPlayer() *Node {
	var (
		bg     = utils.Col(16, 18, 27)
		card   = utils.Col(28, 31, 46)
		track  = utils.Col(45, 50, 75)
		white  = utils.Col(235, 238, 248)
		muted  = utils.Col(140, 146, 175)
		green  = utils.Col(0, 255, 0)
		border = utils.Col(42, 48, 68)
	)

	iconBtn := func(src string, size uint16) Element {
		return Image(src, ImageStyle(Width(Fixed(size)), Height(Fixed(size))))
	}

	return Root(Div(
		Style(
			Width(Grow(1)), Height(Grow(1)),
			LayoutAxis(LAYOUT_VERTICAL),
			Color(bg),
			Justify(JustifyCenter),
			Align(AlignCenter),
		),
		Children(
			Div(
				Style(
					LayoutAxis(LAYOUT_VERTICAL),
					Padding(PaddingAll(24)),
					Gap(GapAll(16)),
					Color(card),
					BorderRadius(0.1),
					BorderColor(border),
					BorderWidth(2),
				),
				Children(
					// album
					Div(
						Style(LayoutAxis(LAYOUT_HORIZONTAL), Width(Grow(1)), Justify(JustifyCenter)),
						Children(
							Image("assets/album.png", ImageStyle(Width(Fixed(240)), Height(Fixed(240)))),
						),
					),

					// title / artist
					Text("Midnight City", TextStyle(FontSize(20), Color(white))),
					Text("M83", TextStyle(Color(muted), FontSize(14))),

					// player
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

					// controls
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
	))
}
