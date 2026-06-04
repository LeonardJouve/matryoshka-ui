package components

import (
	"github.com/LeonardJouve/matryoshka-ui/components"
	. "github.com/LeonardJouve/matryoshka-ui/dsl"
)

// Login is a centered auth card on a full-bleed background.
func Login() *Node {
	card := Div(
		Style(
			LayoutAxis(LAYOUT_VERTICAL),
			Width(Fixed(340)),
			Color(components.T.Surface),
			BorderRadius(0.1),
			BorderColor(components.T.Border),
			BorderWidth(1),
			Padding(PaddingAll(28)),
			Gap(GapVertical(16)),
		),
		Children(
			components.H2("Welcome back"),
			components.Muted("Sign in to your account"),
			components.FixedSpacer(4),
			components.Field("Email", "", "you@example.com"),
			components.Field("Password", "", "••••••••"),
			Div(
				Style(LayoutAxis(LAYOUT_HORIZONTAL), Width(Grow(1)), Align(AlignCenter), Justify(JustifyBetween)),
				Children(
					components.CheckRow("Remember me", true),
					components.Accent("Forgot?", components.T.Primary),
				),
			),
			components.FixedSpacer(4),
			fullWidth(components.PrimaryButton("Sign in")),
			components.Center(components.Muted("Don't have an account? Sign up")),
		),
	)

	return Root(Div(
		Style(
			LayoutAxis(LAYOUT_VERTICAL),
			Width(Grow(1)), Height(Grow(1)),
			Color(components.T.Bg),
			Justify(JustifyCenter),
			Align(AlignCenter),
		),
		Children(card),
	))
}

// SignupScreen is a two-column-ish onboarding card with more fields.
func SignupScreen() *Node {
	card := Div(
		Style(
			LayoutAxis(LAYOUT_VERTICAL),
			Width(Fixed(380)),
			Color(components.T.Surface),
			BorderRadius(0.1),
			BorderColor(components.T.Border),
			BorderWidth(1),
			Padding(PaddingAll(28)),
			Gap(GapVertical(14)),
		),
		Children(
			components.H2("Create account"),
			components.Muted("Start your 14-day free trial"),
			components.FixedSpacer(4),
			components.Row(12,
				components.Field("First name", "Ada", ""),
				components.Field("Last name", "Lovelace", ""),
			),
			components.Field("Email", "ada@example.com", ""),
			components.Field("Password", "", "8+ characters"),
			components.CheckRow("I agree to the Terms of Service", true),
			components.FixedSpacer(4),
			fullWidth(components.PrimaryButton("Create account")),
		),
	)

	return Root(Div(
		Style(
			LayoutAxis(LAYOUT_VERTICAL),
			Width(Grow(1)), Height(Grow(1)),
			Color(components.T.Bg),
			Justify(JustifyCenter),
			Align(AlignCenter),
		),
		Children(card),
	))
}

// fullWidth wraps a (typically fit-sized) child in a row so it stretches.
func fullWidth(child Element) Element {
	return Div(
		Style(LayoutAxis(LAYOUT_HORIZONTAL), Width(Grow(1))),
		Children(
			Div(Style(Width(Grow(1))), Children(child)),
		),
	)
}
