package dsl2

type RootS struct {
	Element
}

func Root(element Element) *RootS {
	r := &RootS{Element: element}

	element.build()

	return r
}

func (d *RootS) build() *Node {
	return nil
}
