package dsl2

type Kind int

const (
	KindDiv Kind = iota
	KindText
	KindImage
)

type Node struct {
	Kind   Kind
	Style  Style
	Layout Rect
	lines  []Line

	DivAttrs
	TextAttrs
	ImageAttrs
}

type TextAttrs struct {
	Content string
	FG      Color
}

func mainSize(n *Node, dir Direction) int {
	if dir == Row {
		return n.Layout.W
	}
	return n.Layout.H
}

func setMainSize(n *Node, dir Direction, v int) {
	if dir == Row {
		n.Layout.W = v
	} else {
		n.Layout.H = v
	}
}

func crossSize(n *Node, dir Direction) int {
	if dir == Row {
		return n.Layout.H
	}
	return n.Layout.W
}
