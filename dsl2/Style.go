package dsl2

type Color struct{ R, G, B, A uint8 }

type Direction int

const (
	Column Direction = iota // main axis = vertical (default block flow)
	Row                     // main axis = horizontal
)

type Style struct {
	W, H    int
	Padding PaddingS
	Gap     GapS
	Dir     Direction // only meaningful on DivS
	//Grow, Shrink int       // flex factors
	//Wrap         Wrap      // only meaningful on DivS
	//Justify      Justify   // main-axis distribution (DivS)
	//Align        Align     // cross-axis alignment of children (DivS)
	Background Color // background (DivS) / fill
	//FG           Color     // text color
	HasBG bool // whether BG should be drawn
}

// Shared style setters, written once against *Style.
func setW(s *Style, w int)                   { s.W = w }
func setH(s *Style, h int)                   { s.H = h }
func setPadding(s *Style, m PaddingModifier) { m(&s.Padding) }
func setGap(s *Style, g GapModifier)         { g(&s.Gap) }
