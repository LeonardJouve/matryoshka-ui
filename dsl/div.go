package dsl

func Div(opts ...ElementModifier) *Element {
	el := Element{
		id:       <-idGenerator,
		children: []*Element{},
		layout:   NewLayout(),
		style:    NewStyle(),
	}
	for _, opt := range opts {
		opt(&el)
	}

	return &el
}
