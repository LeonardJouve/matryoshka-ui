package components

import (
	"github.com/LeonardJouve/matryoshka-ui/components"
	. "github.com/LeonardJouve/matryoshka-ui/dsl"
)

// Chat is a messaging layout: conversation list + message thread.
func Chat() *Node {
	convoList := Div(
		Style(
			LayoutAxis(LAYOUT_VERTICAL),
			Width(Fixed(220)),
			Height(Grow(1)),
			Color(components.T.Surface),
			BorderRadius(0.1),
			Padding(PaddingAll(10)),
			Gap(GapVertical(6)),
		),
		Children(
			Text("Messages", TextStyle(FontSize(14), Color(components.T.Text))),
			components.HDivider(),
			convoRow("Alan T.", "See you at the lab", true),
			convoRow("Grace H.", "Compiler's done!", false),
			convoRow("Charles B.", "Re: the engine", false),
			convoRow("Howard A.", "Mark I update", false),
		),
	)

	thread := Div(
		Style(
			LayoutAxis(LAYOUT_VERTICAL),
			Width(Grow(1)),
			Height(Grow(1)),
			Gap(GapVertical(10)),
		),
		Children(
			Div(
				Style(
					LayoutAxis(LAYOUT_HORIZONTAL),
					Width(Grow(1)),
					Color(components.T.Surface),
					BorderRadius(0.1),
					Padding(PaddingAll(12)),
					Align(AlignCenter),
					Gap(GapHorizontal(10)),
				),
				Children(components.Avatar(32, components.T.Success), components.H3("Alan T."), components.Spacer(), components.Dot(components.T.Success)),
			),
			// message area grows to fill
			Div(
				Style(
					LayoutAxis(LAYOUT_VERTICAL),
					Width(Grow(1)),
					Height(Grow(1)),
					Color(components.T.Surface),
					BorderRadius(0.1),
					Padding(PaddingAll(14)),
					Gap(GapVertical(8)),
				),
				Children(
					components.Bubble("Did the test run finish?", false),
					components.Bubble("Yes — all green.", true),
					components.Bubble("The engine computed Bernoulli numbers.", false),
					components.Bubble("Beautiful. Shipping it.", true),
				),
			),
			// composer
			Div(
				Style(LayoutAxis(LAYOUT_HORIZONTAL), Width(Grow(1)), Gap(GapHorizontal(10)), Align(AlignCenter)),
				Children(
					Div(Style(Width(Grow(1))), Children(components.Field("", "", "Type a message..."))),
					components.Button("Send", components.T.Primary),
				),
			),
		),
	)

	return Root(Div(
		Style(
			LayoutAxis(LAYOUT_HORIZONTAL),
			Width(Grow(1)), Height(Grow(1)),
			Color(components.T.Bg),
			Padding(PaddingAll(16)),
			Gap(GapHorizontal(16)),
		),
		Children(convoList, thread),
	))
}

func convoRow(name string, preview string, active bool) Element {
	bg := components.T.Surface
	if active {
		bg = components.T.Raised
	}
	return Div(
		Style(
			LayoutAxis(LAYOUT_HORIZONTAL),
			Width(Grow(1)),
			Color(bg),
			BorderRadius(0.15),
			Padding(PaddingAll(8)),
			Gap(GapHorizontal(8)),
			Align(AlignCenter),
		),
		Children(
			components.Avatar(32, components.T.Primary),
			Div(
				Style(LayoutAxis(LAYOUT_VERTICAL), Width(Grow(1)), Gap(GapVertical(2))),
				Children(components.Body(name), components.Caption(preview)),
			),
		),
	)
}
