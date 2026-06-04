package components

import (
	"github.com/LeonardJouve/matryoshka-ui/components"
	. "github.com/LeonardJouve/matryoshka-ui/dsl"
)

// Pricing is a centered three-tier pricing page.
func Pricing() *Node {
	return Root(Div(
		Style(
			LayoutAxis(LAYOUT_VERTICAL),
			Width(Grow(1)), Height(Grow(1)),
			Color(components.T.Bg),
			Padding(PaddingAll(24)),
			Gap(GapVertical(20)),
			Justify(JustifyCenter),
			Align(AlignCenter),
		),
		Children(
			Div(
				Style(LayoutAxis(LAYOUT_VERTICAL), Align(AlignCenter), Gap(GapVertical(6))),
				Children(components.H1("Pricing"), components.Muted("Choose the plan that fits your team")),
			),
			Div(
				Style(LayoutAxis(LAYOUT_HORIZONTAL), Width(Fixed(720)), Gap(GapHorizontal(16))),
				Children(
					components.PricingTier("Starter", "$0", components.T.Info, false,
						"1 project", "Community support", "1GB storage"),
					components.PricingTier("Pro", "$29", components.T.Primary, true,
						"Unlimited projects", "Priority support", "100GB storage", "Analytics"),
					components.PricingTier("Team", "$99", components.T.Success, false,
						"Everything in Pro", "SSO", "Audit logs", "Dedicated manager"),
				),
			),
		),
	))
}

// KanbanScreen is a three-column task board.
func KanbanScreen() *Node {
	return components.AppShell(
		components.Sidebar(180, "PROJECT",
			components.NavItem("Board", true),
			components.NavItem("Backlog", false),
			components.NavItem("Timeline", false),
		),
		components.Column(14,
			components.Row(0, components.H2("Sprint Board"), components.Spacer(), components.Button("New task", components.T.Primary)),
			components.Row(12,
				components.KanbanColumn("To Do", "3",
					components.TaskCard("Design login flow", "design", components.T.Info),
					components.TaskCard("Set up CI", "devops", components.T.Warning),
					components.TaskCard("Write API spec", "backend", components.T.Success),
				),
				components.KanbanColumn("In Progress", "2",
					components.TaskCard("Layout engine: percent", "core", components.T.Primary),
					components.TaskCard("Component library", "frontend", components.T.Info),
				),
				components.KanbanColumn("Done", "1",
					components.TaskCard("Repo scaffolding", "chore", components.T.TextMute),
				),
			),
		),
	)
}

// PlayerScreen recreates the music-player card from the original example.
func PlayerScreen() *Node {
	card := Div(
		Style(
			LayoutAxis(LAYOUT_VERTICAL),
			Width(Fixed(300)),
			Color(components.T.Surface),
			BorderRadius(0.1),
			BorderColor(components.T.Border),
			BorderWidth(1),
			Padding(PaddingAll(24)),
			Gap(GapVertical(16)),
		),
		Children(
			// album placeholder (no asset dependency)
			Div(
				Style(LayoutAxis(LAYOUT_HORIZONTAL), Width(Grow(1)), Justify(JustifyCenter)),
				Children(Div(Style(Width(Fixed(240)), Height(Fixed(240)), Color(components.T.Raised), BorderRadius(0.1)))),
			),
			components.H3("Midnight City"),
			components.Muted("M83"),
			// progress
			components.Row(10,
				components.Caption("1:24"),
				Div(Style(Width(Grow(1))), Children(components.ProgressBar(0.34, components.T.Success))),
				components.Caption("4:03"),
			),
			// transport controls (dots stand in for icons)
			Div(
				Style(LayoutAxis(LAYOUT_HORIZONTAL), Width(Grow(1)), Justify(JustifyBetween), Align(AlignCenter)),
				Children(
					components.Dot(components.T.TextMute),
					components.Dot(components.T.Text),
					components.Avatar(40, components.T.Success),
					components.Dot(components.T.Text),
					components.Dot(components.T.TextMute),
				),
			),
		),
	)

	return Root(Div(
		Style(
			LayoutAxis(LAYOUT_VERTICAL),
			Width(Grow(1)), Height(Grow(1)),
			Color(components.T.Bg),
			Justify(JustifyCenter), Align(AlignCenter),
		),
		Children(card),
	))
}
