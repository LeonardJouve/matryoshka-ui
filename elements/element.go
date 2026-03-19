package elements

import (
	"github.com/LeonardJouve/matryoshka-ui/utils"
)

type Element struct {
	id       uint
	children []*Element
	parent   *Element
	style    Style
	layout   Layout
}

type End interface {
	Children() []*Element
	Color() utils.Color
	Width() uint16
	Height() uint16
	X() uint16
	Y() uint16
	getSelf() *Element
}

type IElement interface {
	End() End // TODO element ?
	Childrens(children ...End) IElement
	//When(condition bool, fun func(el *Element) End) IElement
	//ForEach(n uint, fn func(index uint) End) IElement
	Style(style *Style) IElement
}

var idGenerator = utils.NewIDGenerator()

func NewElement() IElement {
	return &Element{
		id:       <-idGenerator,
		children: []*Element{},
		parent:   nil,
		layout:   NewLayout(),
	}
}

func (el *Element) getSelf() *Element {
	return el
}

func (el *Element) End() End {
	var layoutOffset uint16 = 0
	var crossOffset uint16 = 0
	var maxCrossSize uint16 = 0
	if el.style.LayoutAxis == LAYOUT_HORIZONTAL {
		layoutOffset = el.style.Padding.left
		crossOffset = el.style.Padding.top
	} else {
		layoutOffset = el.style.Padding.top
		crossOffset = el.style.Padding.left
	}

	for _, child := range el.children {
		if el.style.LayoutAxis == LAYOUT_HORIZONTAL {
			width := child.Width()
			maxCrossSize = max(maxCrossSize, child.Height())

			// if fixed width and content overflows, increment cross offset
			if el.style.Width != 0 && layoutOffset != 0 && layoutOffset+width > el.style.Width {
				crossOffset += maxCrossSize + el.style.Gap.vertical // max height
				layoutOffset = el.style.Padding.left
				maxCrossSize = 0
			}

			child.layout.X = layoutOffset
			child.layout.Y = crossOffset
			layoutOffset += width + el.style.Gap.horizontal
		} else {
			height := child.Height()
			maxCrossSize = max(maxCrossSize, child.Width())

			// if fixed width and content overflows, increment cross offset
			if el.style.Height != 0 && layoutOffset != 0 && layoutOffset+height > el.style.Height {
				crossOffset += maxCrossSize + el.style.Gap.horizontal // max height
				layoutOffset = el.style.Padding.top
				maxCrossSize = 0
			}

			child.layout.Y = layoutOffset
			child.layout.X = crossOffset
			layoutOffset += height + el.style.Gap.vertical
		}
	}

	if el.style.Width != 0 {
		el.layout.Width = el.style.Width
	} else {
		if el.style.LayoutAxis == LAYOUT_HORIZONTAL {
			el.layout.Width = layoutOffset
		} else {
			el.layout.Width = crossOffset
		}
	}

	if el.style.Height != 0 {
		el.layout.Height = el.style.Height
	} else {
		if el.style.LayoutAxis == LAYOUT_HORIZONTAL {
			el.layout.Height = crossOffset
		} else {
			el.layout.Height = layoutOffset
		}
	}

	return el
}

/*
	func (el *Element) When(condition bool, fun func(el *Element) End) IElement {
		if condition {
			return fun(el).getSelf()
		}

		return el
	}

	func (el *Element) ForEach(n uint, fn func(index uint) End) IElement {
		for i := range n {
			child := fn(i).getSelf()
			child.parent = el
			el.children = append(el.children, child)
		}

		return el
	}
*/
func (el *Element) Childrens(elements ...End) IElement {
	for _, child := range elements {
		child.getSelf().parent = el
		el.children = append(el.children, child.getSelf())
	}
	return el
}

func (el *Element) Style(style *Style) IElement {
	el.style = *style
	return el
}

// ACCESSOR

func (el *Element) Children() []*Element {
	return el.children
}

func (el *Element) Color() utils.Color {
	return el.style.Color
}

func (el *Element) X() uint16 {
	return el.layout.X
}

func (el *Element) Y() uint16 {
	return el.layout.Y
}

func (el *Element) Width() uint16 {
	return el.layout.Width
}

func (el *Element) Height() uint16 {
	return el.layout.Height
}

func (el *Element) getWidth() uint16 {
	if el.style.Width != 0 {
		return el.style.Width
	}

	var width uint16 = 0

	for _, child := range el.children {
		if el.style.LayoutAxis == LAYOUT_HORIZONTAL {
			width += child.layout.Width
		} else {
			width = max(width, child.layout.Width)
		}
	}

	return width
}

func (el *Element) getHeight() uint16 {
	if el.style.Height != 0 {
		return el.style.Height
	}

	var height uint16 = 0

	for _, child := range el.children {
		if el.style.LayoutAxis == LAYOUT_HORIZONTAL {
			height += child.layout.Height
		} else {
			height = max(height, child.layout.Height)
		}
	}

	return height
}
