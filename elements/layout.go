package elements

type Layout struct {
	Width  uint16
	Height uint16
	X      uint16
	Y      uint16
}

func NewLayout() Layout {
	return Layout{}
}
