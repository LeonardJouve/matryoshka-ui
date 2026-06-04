package components

import (
	. "github.com/LeonardJouve/matryoshka-ui/dsl"
)

// NavItem is a single sidebar/nav entry. Pass active=true to highlight it.
func NavItem(label string, active bool) Element {
	textColor := T.TextMute
	if active {
		textColor = T.Text
	}
	bg := T.Surface
	if active {
		bg = T.Raised
	}
	return Div(
		Style(
			LayoutAxis(LAYOUT_HORIZONTAL),
			Width(Grow(1)),
			Color(bg),
			BorderRadius(0.15),
			Padding(Padding2(8, 10)),
			Align(AlignCenter),
		),
		Children(
			Text(label, TextStyle(FontSize(13), Color(textColor))),
		),
	)
}

// Sidebar is a fixed-width vertical nav column with a title and items.
func Sidebar(width uint16, title string, items ...Element) Element {
	body := append([]Element{
		Text(title, TextStyle(FontSize(14), Color(T.Text))),
		HDivider(),
	}, items...)

	return Div(
		Style(
			LayoutAxis(LAYOUT_VERTICAL),
			Width(Fixed(width)),
			Height(Grow(1)),
			Color(T.Surface),
			BorderRadius(0.1),
			Padding(PaddingAll(12)),
			Gap(GapVertical(8)),
		),
		Children(body...),
	)
}

// Navbar is a horizontal top bar with a brand on the left and items on the right.
func Navbar(brand string, right ...Element) Element {
	left := Text(brand, TextStyle(FontSize(16), Color(T.Text)))

	children := []Element{left, Spacer()}
	children = append(children, right...)

	return Div(
		Style(
			LayoutAxis(LAYOUT_HORIZONTAL),
			Width(Grow(1)),
			Color(T.Surface),
			BorderRadius(0.1),
			Padding(Padding2(10, 16)),
			Gap(GapHorizontal(12)),
			Align(AlignCenter),
		),
		Children(children...),
	)
}

// AppShell composes a sidebar and a main content column side by side,
// filling the whole window with the background color. It returns the
// root node ready to hand to a renderer.
func AppShell(sidebar Element, content Element) *Node {
	return Root(Div(
		Style(
			LayoutAxis(LAYOUT_HORIZONTAL),
			Width(Grow(1)),
			Height(Grow(1)),
			Color(T.Bg),
			Padding(PaddingAll(16)),
			Gap(GapHorizontal(16)),
		),
		Children(
			sidebar,
			Div(
				Style(
					LayoutAxis(LAYOUT_VERTICAL),
					Width(Grow(1)),
					Height(Grow(1)),
					Gap(GapVertical(14)),
				),
				Children(content),
			),
		),
	))
}
