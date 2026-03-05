package elements

type Gap struct {
	vertical   uint16
	horizontal uint16
}

func NewGap(vertical uint16, horizontal uint16) Gap {
	return Gap{
		vertical,
		horizontal,
	}
}

func NewVerticalGap(gap uint16) Gap {
	return Gap{
		vertical: gap,
	}
}

func NewHorizontalGap(gap uint16) Gap {
	return Gap{
		horizontal: gap,
	}
}
