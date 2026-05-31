package dsl2

type DivS struct {
	Style
	Kids []Element
}

type DivAttrs struct {
	Children []*Node
}

func Div(opts ...DivOpt) *DivS {
	d := DivS{}
	for _, o := range opts {
		o(&d)
	}
	return &d
}

func (d *DivS) build() *Node {
	n := &Node{Kind: KindDiv, Style: d.Style}
	for _, k := range d.Kids {
		n.Children = append(n.Children, k.build())
	}
	return n
}

type DivOpt func(*DivS)

func Width(w int) DivOpt  { return func(d *DivS) { setW(&d.Style, w) } }
func Height(h int) DivOpt { return func(d *DivS) { setH(&d.Style, h) } }
func Padding(p ...PaddingModifier) DivOpt {
	return func(d *DivS) {
		for _, m := range p {
			setPadding(&d.Style, m)
		}
	}
}
func Gap(g ...GapModifier) DivOpt {
	return func(d *DivS) {
		for _, m := range g {
			setGap(&d.Style, m)
		}
	}
}
func LayoutAxis(dir Direction) DivOpt { return func(d *DivS) { d.Style.Dir = dir } }
func Background(c Color) DivOpt       { return func(d *DivS) { d.Style.Background = c } }
func Children(elems ...Element) DivOpt {
	return func(d *DivS) {
		for _, elem := range elems {
			d.Kids = append(d.Kids, elem)
		}
	}
}
