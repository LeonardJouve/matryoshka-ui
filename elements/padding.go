package elements

type Padding struct {
	left   uint16
	top    uint16
	right  uint16
	bottom uint16
}

func NewPadding(left uint16, top uint16, right uint16, bottom uint16) Padding {
	return Padding{
		left:   left,
		top:    top,
		right:  right,
		bottom: bottom,
	}
}

func NewPaddingVertical(padding uint16) Padding {
	return Padding{
		top:    padding,
		bottom: padding,
	}
}

func NewPaddingHorizontal(padding uint16) Padding {
	return Padding{
		left:  padding,
		right: padding,
	}
}
