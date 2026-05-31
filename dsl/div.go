package dsl

type DivS struct {
	StyleS *StyleS
	Kids   []Element
}

type DivAttrs struct {
	Children []*Node
}

type DivOpt func(*DivS)

func Div(opts ...DivOpt) *DivS {
	d := DivS{
		StyleS: NewStyle(),
	}
	for _, opt := range opts {
		opt(&d)
	}
	return &d
}

func (d *DivS) build() *Node {
	n := &Node{
		Kind:     KindDiv,
		Style:    d.StyleS,
		Layout:   NewLayout(),
		lines:    []Line{},
		DivAttrs: DivAttrs{},
	}

	for _, k := range d.Kids {
		n.Children = append(n.Children, k.build())
	}

	return n
}

func Style(modifiers ...StyleModifier) DivOpt {
	return func(s *DivS) {
		for _, modifier := range modifiers {
			modifier(s.StyleS)
		}
	}
}

func Children(els ...Element) DivOpt {
	return func(d *DivS) {
		d.Kids = append(d.Kids, els...)
	}
}
