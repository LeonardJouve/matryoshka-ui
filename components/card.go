package components

import (
	. "github.com/LeonardJouve/matryoshka-ui/dsl"
	"github.com/LeonardJouve/matryoshka-ui/utils"
)

// Card is a padded surface container with a border, growing to fill width.
func Card(children ...Element) Element {
	return Div(
		Style(
			LayoutAxis(LAYOUT_VERTICAL),
			Width(Grow(1)),
			Height(Grow(1)),
			Color(T.Surface),
			BorderRadius(0.1),
			BorderColor(T.Border),
			BorderWidth(1),
			Padding(PaddingAll(16)),
			Gap(GapVertical(12)),
		),
		Children(children...),
	)
}

// TitledCard is a card with a heading row above its body.
func TitledCard(title string, children ...Element) Element {
	body := append([]Element{
		Text(title, TextStyle(FontSize(16), Color(T.Text))),
		HDivider(),
	}, children...)

	return Card(body...)
}

// Panel is a borderless raised surface — lighter weight than Card.
func Panel(children ...Element) Element {
	return Div(
		Style(
			LayoutAxis(LAYOUT_VERTICAL),
			Width(Grow(1)),
			Height(Grow(1)),
			Color(T.Surface),
			BorderRadius(0.1),
			BorderColor(T.Border),
			BorderWidth(1),
			Padding(PaddingAll(14)),
			Gap(GapVertical(10)),
		),
		Children(children...),
	)
}

// StatCard shows a small label above a large accent-colored value.
func StatCard(label string, value string, accent utils.Color) Element {
	return Div(
		Style(
			LayoutAxis(LAYOUT_VERTICAL),
			Width(Grow(1)),
			Color(T.Surface),
			BorderRadius(0.1),
			BorderColor(T.Border),
			BorderWidth(1),
			Padding(PaddingAll(14)),
			Gap(GapVertical(6)),
		),
		Children(
			Text(label, TextStyle(FontSize(12), Color(T.TextMute))),
			Text(value, TextStyle(FontSize(22), Color(accent))),
		),
	)
}
