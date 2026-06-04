package utils

type Color struct {
	Red   uint8
	Green uint8
	Blue  uint8
	Alpha uint8
}

func Col(r, g, b uint8) Color { return Color{Red: r, Green: g, Blue: b, Alpha: 255} }
