package dsl

type Kind int

const (
	KindDiv Kind = iota
	KindText
	KindImage
)

type Node struct {
	Kind   Kind
	Style  *StyleS
	Layout LayoutT

	DivAttrs
	ImageAttrs
	TextAttrs
}

func (n *Node) setAxisSize(l LayoutAxisT, size uint16) {
	if l == LAYOUT_HORIZONTAL {
		n.Layout.Width = size
		return
	}

	n.Layout.Height = size
}

func (n *Node) setMainAxisSize(size uint16) {
	n.setAxisSize(n.Style.LayoutAxis, size)
}

func (n *Node) setCrossAxisSize(size uint16) {
	n.setAxisSize(n.Style.oppositeAxis(), size)
}

func (n *Node) axisSize(l LayoutAxisT) uint16 {
	if l == LAYOUT_HORIZONTAL {
		return n.Layout.Width
	}

	return n.Layout.Height
}

func (n *Node) mainAxisSize() uint16 {
	return n.axisSize(n.Style.LayoutAxis)
}

func (n *Node) crossAxisSize() uint16 {
	return n.axisSize(n.Style.oppositeAxis())
}

func (n *Node) setPosition(l LayoutAxisT, position uint16) {
	if l == LAYOUT_HORIZONTAL {
		n.Layout.X = position
	}

	n.Layout.Y = position
}

func (n *Node) setMainPosition(position uint16) {
	n.setPosition(n.Style.LayoutAxis, position)
}

func (n *Node) setCrossPosition(position uint16) {
	n.setPosition(n.Style.oppositeAxis(), position)
}
