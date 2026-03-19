package elements

import "github.com/LeonardJouve/matryoshka-ui/utils"

type LayoutAxis uint16

const (
	LAYOUT_HORIZONTAL LayoutAxis = iota
	LAYOUT_VERTICAL
)

type Style struct {
	layoutAxis LayoutAxis
	color      utils.Color
	padding    Padding
	gap        Gap
	width      uint16
	height     uint16
}

func NewStyle() *Style {
	return &Style{
		layoutAxis: LAYOUT_HORIZONTAL,
		color: utils.Color{
			Red:   0,
			Green: 0,
			Blue:  0,
		},
		padding: Padding{},
		gap:     Gap{},
		width:   0,
		height:  0,
	}
}

func (style *Style) LayoutAxis(layoutAxis LayoutAxis) *Style {
	style.LayoutAxis = layoutAxis
	return style
}

func (style *Style) LayoutAxis() *Style {
	return style.layoutAxis
}

func (style *Style) Padding(padding Padding) *Style {
	style.Padding = padding
	return style
}

func (style *Style) Gap(gap Gap) *Style {
	style.Gap = gap
	return style
}

func (style *Style) Color(color utils.Color) *Style {
	style.Color = color
	return style
}

func (style *Style) Width(u uint16) *Style {
	style.Width = u
	return style
}

func (style *Style) Height(u uint16) *Style {
	style.Height = u
	return style
}
