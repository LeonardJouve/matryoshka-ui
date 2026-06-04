package components

import (
	. "github.com/LeonardJouve/matryoshka-ui/dsl"
	"github.com/LeonardJouve/matryoshka-ui/utils"
)

// ProgressBar renders a track with a fill. ratio is 0.0..1.0 and uses
// percentage sizing, so the parent must have a determined width.
func ProgressBar(ratio float64, fill utils.Color) Element {
	if ratio < 0 {
		ratio = 0
	}
	if ratio > 1 {
		ratio = 1
	}
	return Div(
		Style(
			LayoutAxis(LAYOUT_HORIZONTAL),
			Width(Grow(1)),
			Height(Fixed(8)),
			Color(T.Border),
			BorderRadius(0.8),
		),
		Children(
			Div(Style(
				Width(Percent(ratio*100)),
				Height(Grow(1)),
				Color(fill),
				BorderRadius(0.8),
			)),
		),
	)
}

// Bar is a single vertical bar of fixed height — a chart building block.
func Bar(h uint16, c utils.Color) Element {
	return Div(Style(
		Width(Fixed(18)),
		Height(Fixed(h)),
		Color(c),
		BorderRadius(0.2),
	))
}

// BarChart lays a row of bars bottom-aligned. Each value is a pixel height.
func BarChart(height uint16, c utils.Color, values ...uint16) Element {
	bars := make([]Element, len(values))
	for i, v := range values {
		bars[i] = Bar(v, c)
	}
	return Div(
		Style(
			LayoutAxis(LAYOUT_HORIZONTAL),
			Width(Grow(1)),
			Height(Fixed(height)),
			Align(AlignEnd),
			Gap(GapHorizontal(6)),
		),
		Children(bars...),
	)
}

// ListRow is a labeled row with an optional trailing accent value.
func ListRow(label string, value string, accent utils.Color) Element {
	return Div(
		Style(
			LayoutAxis(LAYOUT_HORIZONTAL),
			Width(Grow(1)),
			Justify(JustifyBetween),
			Align(AlignCenter),
			Padding(Padding2(6, 0)),
		),
		Children(
			Text(label, TextStyle(FontSize(13), Color(T.Text))),
			Text(value, TextStyle(FontSize(13), Color(accent))),
		),
	)
}

// KeyValue stacks a small label over a value.
func KeyValue(key string, value string) Element {
	return Div(
		Style(LayoutAxis(LAYOUT_VERTICAL), Gap(GapVertical(2))),
		Children(
			Label(key),
			Body(value),
		),
	)
}

// Avatar is a solid rounded square placeholder of a given size.
func Avatar(size uint16, c utils.Color) Element {
	return Div(Style(
		Width(Fixed(size)),
		Height(Fixed(size)),
		Color(c),
		BorderRadius(1),
	))
}

// Dot is a small colored status indicator.
func Dot(c utils.Color) Element {
	return Div(Style(
		Width(Fixed(8)),
		Height(Fixed(8)),
		Color(c),
		BorderRadius(1),
	))
}
