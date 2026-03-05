package elements

import (
	"github.com/LeonardJouve/matryoshka-ui/utils"
)

type Element struct {
	Children []*Element
	Layout   LayoutType
	Padding  Padding
	Gap      Gap
	Width    uint16
	Height   uint16
	Color    utils.Color
	X        uint16
	Y        uint16
	id       uint
}

type LayoutType uint16

const (
	LAYOUT_HORIZONTAL LayoutType = iota
	LAYOUT_VERTICAL
)

type End interface {
	End()
}

type IElement interface {
	End() *Element
	// TODO Element ?
	AddChild(children ...*Element) IElement
	RemoveChild(children ...*Element) IElement
	GetWidth() uint16
	GetHeight() uint16
	SetLayout(layout LayoutType) IElement
	SetWidth(uint16) IElement
	SetHeight(uint16) IElement
	SetGap(gap Gap) IElement
	SetColor(color utils.Color) IElement
	GetX() uint16
	GetY() uint16
	When(condition bool, fun func(el *Element) IElement) IElement
	ForEach(n uint, fn func(index uint) *Element) IElement
	SetPadding(padding Padding) IElement
}

var idGenerator = utils.NewIDGenerator()

func NewElement() IElement {
	return &Element{
		Children: []*Element{},
		Layout:   LAYOUT_HORIZONTAL,
		Width:    0,
		Height:   0,
		Color: utils.Color{
			Red:   0,
			Green: 0,
			Blue:  0,
		},
		Padding: Padding{},
		X:       0,
		Y:       0,
		id:      <-idGenerator,
	}
}

func (el *Element) End() *Element {
	var layoutOffset uint16
	var crossOffset uint16
	if el.Layout == LAYOUT_HORIZONTAL {
		layoutOffset = el.Padding.left
		crossOffset = el.Padding.top
	} else {
		layoutOffset = el.Padding.top
		crossOffset = el.Padding.left
	}

	var maxCrossSize uint16 = 0

	for _, child := range el.Children {
		if el.Layout == LAYOUT_HORIZONTAL {
			width := child.GetWidth()
			maxCrossSize = max(maxCrossSize, child.GetHeight())

			// if fixed width and content overflows, increment cross offset
			if el.Width != 0 && layoutOffset != 0 && layoutOffset+width > el.Width {
				crossOffset += maxCrossSize + el.Gap.vertical // max height
				layoutOffset = el.Padding.left
				maxCrossSize = 0
			}

			child.X = layoutOffset
			child.Y = crossOffset
			layoutOffset += width + el.Gap.horizontal
		} else {
			height := child.GetHeight()
			maxCrossSize = max(maxCrossSize, child.GetWidth())

			// if fixed width and content overflows, increment cross offset
			if el.Height != 0 && layoutOffset != 0 && layoutOffset+height > el.Height {
				crossOffset += maxCrossSize + el.Gap.horizontal // max height
				layoutOffset = el.Padding.top
				maxCrossSize = 0
			}

			child.Y = layoutOffset
			child.X = crossOffset
			layoutOffset += height + el.Gap.vertical
		}
	}
	return el
}

func (el *Element) When(condition bool, fun func(el *Element) IElement) IElement {
	if condition {
		return fun(el)
	}

	return el
}

func (el *Element) ForEach(n uint, fn func(index uint) *Element) IElement {
	for i := range n {
		el.Children = append(el.Children, fn(i))
	}
	return el
}
func (el *Element) SetLayout(layout LayoutType) IElement {
	el.Layout = layout
	return el
}

func (el *Element) AddChild(elements ...*Element) IElement {
	el.Children = append(el.Children, elements...)
	return el
}

func (el *Element) RemoveChild(elements ...*Element) IElement {
	removeSet := make(map[uint]struct{})

	for _, element := range elements {
		removeSet[element.id] = struct{}{}
	}

	newChildren := []*Element{}

	for _, child := range el.Children {
		if _, shouldRemove := removeSet[child.id]; !shouldRemove {
			newChildren = append(newChildren, child)
		}
	}

	el.Children = newChildren

	return el
}

func (el *Element) SetPadding(padding Padding) IElement {
	el.Padding = padding
	return el
}

func (el *Element) SetGap(gap Gap) IElement {
	el.Gap = gap
	return el
}

func (el *Element) GetX() uint16 {
	return el.X
}

func (el *Element) GetY() uint16 {
	return el.Y
}

func (el *Element) SetColor(color utils.Color) IElement {
	el.Color = color
	return el
}

func (el *Element) SetWidth(u uint16) IElement {
	el.Width = u
	return el
}

func (el *Element) SetHeight(u uint16) IElement {
	el.Height = u
	return el
}

func (el *Element) GetWidth() uint16 {
	if el.Width != 0 {
		return el.Width
	}

	var width uint16 = 0

	for _, child := range el.Children {
		if el.Layout == LAYOUT_HORIZONTAL {
			width += child.GetWidth()
		} else {
			width = max(width, child.GetWidth())
		}
	}

	return width
}

func (el *Element) GetHeight() uint16 {
	if el.Height != 0 {
		return el.Height
	}

	var height uint16 = 0

	for _, child := range el.Children {
		if el.Layout == LAYOUT_HORIZONTAL {
			height += child.GetHeight()
		} else {
			height = max(height, child.GetHeight())
		}
	}

	return height
}
