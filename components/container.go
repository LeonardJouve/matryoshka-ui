package components

import (
	. "github.com/LeonardJouve/matryoshka-ui/dsl"
	"github.com/LeonardJouve/matryoshka-ui/utils"
)

// Tabs renders a horizontal tab strip. active is the index highlighted.
func Tabs(active int, labels ...string) Element {
	items := make([]Element, len(labels))
	for i, l := range labels {
		fg := T.TextMute
		bg := T.Surface
		if i == active {
			fg = T.Text
			bg = T.Raised
		}
		items[i] = Div(
			Style(
				LayoutAxis(LAYOUT_HORIZONTAL),
				Color(bg),
				BorderRadius(0.2),
				Padding(Padding2(8, 14)),
				Justify(JustifyCenter),
			),
			Children(Text(l, TextStyle(FontSize(13), Color(fg)))),
		)
	}
	return Div(
		Style(LayoutAxis(LAYOUT_HORIZONTAL), Gap(GapHorizontal(6))),
		Children(items...),
	)
}

// List stacks rows in a surface with a divider between each.
func List(rows ...Element) Element {
	withDividers := make([]Element, 0, len(rows)*2)
	for i, r := range rows {
		if i > 0 {
			withDividers = append(withDividers, HDivider())
		}
		withDividers = append(withDividers, r)
	}
	return Div(
		Style(
			LayoutAxis(LAYOUT_VERTICAL),
			Width(Grow(1)),
			Color(T.Surface),
			BorderRadius(0.1),
			Padding(PaddingAll(4)),
		),
		Children(withDividers...),
	)
}

// TableRow renders cells across the main axis, each growing equally.
// Pass header=true to style it as a header row.
func TableRow(header bool, cells ...string) Element {
	fg := T.Text
	size := uint16(13)
	if header {
		fg = T.TextMute
		size = 11
	}
	items := make([]Element, len(cells))
	for i, c := range cells {
		items[i] = Div(
			Style(Width(Grow(1)), Padding(Padding2(8, 10))),
			Children(Text(c, TextStyle(FontSize(size), Color(fg)))),
		)
	}
	return Div(
		Style(LayoutAxis(LAYOUT_HORIZONTAL), Width(Grow(1)), Align(AlignCenter)),
		Children(items...),
	)
}

// Table assembles a header row plus body rows with dividers.
func Table(header Element, rows ...Element) Element {
	body := []Element{header, HDivider()}
	for i, r := range rows {
		if i > 0 {
			body = append(body, HDivider())
		}
		body = append(body, r)
	}
	return Div(
		Style(
			LayoutAxis(LAYOUT_VERTICAL),
			Width(Grow(1)),
			Color(T.Surface),
			BorderRadius(0.1),
			Padding(PaddingAll(4)),
		),
		Children(body...),
	)
}

// Bubble is a chat message. mine=true right-aligns and uses the primary color.
func Bubble(text string, mine bool) Element {
	fill := T.Raised
	just := JustifyStart
	if mine {
		fill = T.Primary
		just = JustifyEnd
	}
	msg := Div(
		Style(
			LayoutAxis(LAYOUT_HORIZONTAL),
			Width(Fixed(220)),
			Color(fill),
			BorderRadius(0.3),
			Padding(Padding2(8, 12)),
		),
		Children(Text(text, TextStyle(FontSize(13), Color(T.Text)))),
	)
	return Div(
		Style(LayoutAxis(LAYOUT_HORIZONTAL), Width(Grow(1)), Justify(just)),
		Children(msg),
	)
}

// KanbanColumn is a titled vertical stack of cards with a count badge.
func KanbanColumn(title string, count string, cards ...Element) Element {
	head := Div(
		Style(LayoutAxis(LAYOUT_HORIZONTAL), Width(Grow(1)), Align(AlignCenter), Justify(JustifyBetween)),
		Children(
			Text(title, TextStyle(FontSize(13), Color(T.Text))),
			Badge(count, T.Raised),
		),
	)
	body := append([]Element{head, HDivider()}, cards...)
	return Div(
		Style(
			LayoutAxis(LAYOUT_VERTICAL),
			Width(Grow(1)),
			Color(T.Surface),
			BorderRadius(0.1),
			Padding(PaddingAll(12)),
			Gap(GapVertical(10)),
		),
		Children(body...),
	)
}

// TaskCard is a small card for kanban boards: title plus a colored tag.
func TaskCard(title string, tag string, tagColor utils.Color) Element {
	return Div(
		Style(
			LayoutAxis(LAYOUT_VERTICAL),
			Width(Grow(1)),
			Color(T.Raised),
			BorderRadius(0.15),
			Padding(PaddingAll(10)),
			Gap(GapVertical(8)),
		),
		Children(
			Body(title),
			Row(0, Badge(tag, tagColor), Spacer()),
		),
	)
}

// PricingTier is a single pricing card: name, price, feature lines, and a CTA.
func PricingTier(name string, price string, accent utils.Color, featured bool, features ...string) Element {
	border := T.Border
	bw := uint16(1)
	if featured {
		border = accent
		bw = 2
	}
	body := []Element{
		Text(name, TextStyle(FontSize(16), Color(T.Text))),
		Text(price, TextStyle(FontSize(28), Color(accent))),
		HDivider(),
	}
	for _, f := range features {
		body = append(body, Muted("- "+f))
	}
	body = append(body, Spacer(), Button("Choose", accent))
	return Div(
		Style(
			LayoutAxis(LAYOUT_VERTICAL),
			Width(Grow(1)),
			Color(T.Surface),
			BorderRadius(0.1),
			BorderColor(border),
			BorderWidth(bw),
			Padding(PaddingAll(18)),
			Gap(GapVertical(12)),
		),
		Children(body...),
	)
}
