package components

import "github.com/LeonardJouve/matryoshka-ui/utils"

// Theme holds a palette so components share consistent colors.
// Override any field to reskin the whole component set.
type Theme struct {
	Bg      utils.Color
	Surface utils.Color
	Raised  utils.Color
	Border  utils.Color

	Text     utils.Color
	TextMute utils.Color

	Primary utils.Color
	Success utils.Color
	Warning utils.Color
	Danger  utils.Color
	Info    utils.Color
}

// DefaultTheme is a dark palette matching the dashboard example.
var DefaultTheme = Theme{
	Bg:      utils.Col(16, 18, 27),
	Surface: utils.Col(28, 31, 46),
	Raised:  utils.Col(38, 42, 60),
	Border:  utils.Col(42, 48, 68),

	Text:     utils.Col(235, 238, 248),
	TextMute: utils.Col(140, 146, 175),

	Primary: utils.Col(90, 160, 255),
	Success: utils.Col(0, 255, 120),
	Warning: utils.Col(255, 170, 80),
	Danger:  utils.Col(255, 80, 80),
	Info:    utils.Col(120, 200, 255),
}

// T is the active theme used by all helpers in this package.
// Reassign it before building your tree to reskin everything.
var T = DefaultTheme
