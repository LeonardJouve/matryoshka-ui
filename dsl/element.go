package dsl

import "github.com/LeonardJouve/matryoshka-ui/utils"

type Element struct {
	id       uint
	style    *StyleS
	children []*Element
	layout   Layout
}

type ElementModifier func(element *Element)

var idGenerator = utils.NewIDGenerator()

func Children(els ...*Element) ElementModifier {
	return func(el *Element) {
		for _, child := range els {
			el.children = append(el.children, child)
		}
	}
}

func Style(modifiers ...StyleModifier) ElementModifier {
	return func(el *Element) {
		for _, modifier := range modifiers {
			modifier(el.style)
		}
	}
}

func (el *Element) Width() uint16 {
	return el.layout.Width
}

func (el *Element) Height() uint16 {
	return el.layout.Height
}

func (el *Element) layoutAxisSize() uint16 {
	if el.style.layoutAxis == LAYOUT_HORIZONTAL {
		return el.Width()
	}

	return el.Height()
}

func (el *Element) crossAxisSize() uint16 {
	if el.style.layoutAxis == LAYOUT_HORIZONTAL {
		return el.Height()
	}

	return el.Width()
}

func (el *Element) axisSize(axis LayoutAxisT) uint16 {
	if axis == LAYOUT_HORIZONTAL {
		return el.Width()
	}

	return el.Height()
}

func (el *Element) axisSizeSet(axis LayoutAxisT, size uint16) {
	if axis == LAYOUT_HORIZONTAL {
		el.layout.Width = size
		return
	}

	el.layout.Height = size
}

func (el *Element) layoutAxisSet(size uint16) {
	if el.style.layoutAxis == LAYOUT_HORIZONTAL {
		el.layout.Width = size
		return
	}
	el.layout.Height = size
}

func (el *Element) crossAxisSet(size uint16) {
	if el.style.layoutAxis == LAYOUT_HORIZONTAL {
		el.layout.Height = size
		return
	}
	el.layout.Width = size
}

func (el *Element) Children() []*Element {
	return el.children
}

func (el *Element) Color() utils.Color {
	return el.style.color
}

func (el *Element) X() uint16 {
	return el.layout.X
}

func (el *Element) layoutAxisPosition() uint16 {
	if el.style.layoutAxis == LAYOUT_HORIZONTAL {
		return el.X()
	}

	return el.Y()
}

func (el *Element) crossAxisPosition() uint16 {
	if el.style.layoutAxis == LAYOUT_HORIZONTAL {
		return el.Y()
	}

	return el.X()
}

func (el *Element) layoutPositionSet(position uint16) {
	if el.style.layoutAxis == LAYOUT_HORIZONTAL {
		el.layout.X = position
		return
	}

	el.layout.Y = position
}

func (el *Element) crossPositionSet(position uint16) {
	if el.style.layoutAxis == LAYOUT_HORIZONTAL {
		el.layout.Y = position
		return
	}

	el.layout.X = position
}

func (el *Element) Y() uint16 {
	return el.layout.Y
}
