package example

import (
	. "github.com/LeonardJouve/matryoshka-ui/dsl"
	"github.com/LeonardJouve/matryoshka-ui/utils"
)

func Percentage() *Node {
	var (
		bg   = utils.Col(16, 18, 27)
		half = utils.Col(70, 130, 220) // Percent(50)
		quar = utils.Col(220, 130, 70) // Percent(25)
		grow = utils.Col(70, 200, 130) // Grow(1) — eats the remaining 25%
		fix  = utils.Col(200, 70, 130) // Fixed
	)

	return Root(Div(
		Style(
			Width(Grow(1)), Height(Grow(1)),
			LayoutAxis(LAYOUT_VERTICAL),
			Color(bg),
			Padding(PaddingAll(20)),
			Gap(GapVertical(20)),
		),
		Children(
			// 50% | 25% | Grow(1)
			Div(
				Style(
					LayoutAxis(LAYOUT_HORIZONTAL),
					Width(Grow(1)), Height(Fixed(120)),
					Gap(GapHorizontal(0)),
				),
				Children(
					Div(Style(Width(Percent(50)), Height(Grow(1)), Color(half))),
					Div(Style(Width(Percent(25)), Height(Grow(1)), Color(quar))),
					Div(Style(Width(Grow(1)), Height(Grow(1)), Color(grow))),
				),
			),

			// 100% | 50% | 25% | 40%
			Div(
				Style(
					LayoutAxis(LAYOUT_HORIZONTAL),
					Width(Grow(1)), Height(Fixed(160)),
					Gap(GapHorizontal(20)),
					Align(AlignStart),
				),
				Children(
					Div(Style(Width(Fixed(100)), Height(Percent(100)), Color(half))),
					Div(Style(Width(Fixed(100)), Height(Percent(50)), Color(quar))),
					Div(Style(Width(Fixed(100)), Height(Percent(25)), Color(grow))),
					Div(Style(Width(Fixed(100)), Height(Fixed(40)), Color(fix))),
				),
			),

			Div(
				Style(
					LayoutAxis(LAYOUT_HORIZONTAL),
					Width(Percent(50)), Height(Fixed(120)),
					Color(utils.Col(40, 44, 60)),
					Padding(PaddingAll(10)),
				),
				Children(
					Div(Style(Width(Percent(100)), Height(Grow(1)), Color(half))),
				),
			),
		),
	))
}
