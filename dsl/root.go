package dsl

type RootS struct {
	*Element
}

func Root(element *Element) *RootS {
	position(element, nil, 0, 0)

	return &RootS{
		Element: element,
	}
}

func position(element *Element, parent *Element, horizontalOffset uint16, verticalOffset uint16) {
	// DFS preordre
	if parent == nil {
		element.layout.X = 0
		element.layout.Y = 0
	} else {
		element.layout.X = parent.layout.X
		element.layout.Y = parent.layout.Y
	}

	for _, child := range element.Children() {
		position(child, element, horizontalOffset, verticalOffset)
	}
}
