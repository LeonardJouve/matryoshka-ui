package components

import (
	. "github.com/LeonardJouve/matryoshka-ui/dsl"
	"github.com/LeonardJouve/matryoshka-ui/utils"
)

// H1 is the largest heading.
func H1(s string) Element {
	return Text(s, TextStyle(FontSize(28), Color(T.Text)))
}

// H2 is a section heading.
func H2(s string) Element {
	return Text(s, TextStyle(FontSize(22), Color(T.Text)))
}

// H3 is a sub-section heading.
func H3(s string) Element {
	return Text(s, TextStyle(FontSize(18), Color(T.Text)))
}

// Body is default-weight readable text.
func Body(s string) Element {
	return Text(s, TextStyle(FontSize(14), Color(T.Text)))
}

// Muted is secondary, de-emphasized text.
func Muted(s string) Element {
	return Text(s, TextStyle(FontSize(13), Color(T.TextMute)))
}

// Label is a small caption, typically above a value or input.
func Label(s string) Element {
	return Text(s, TextStyle(FontSize(11), Color(T.TextMute)))
}

// Caption is the smallest text size.
func Caption(s string) Element {
	return Text(s, TextStyle(FontSize(10), Color(T.TextMute)))
}

// Accent renders text in an arbitrary color at body size.
func Accent(s string, c utils.Color) Element {
	return Text(s, TextStyle(FontSize(14), Color(c)))
}
