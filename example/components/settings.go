package components

import (
	"github.com/LeonardJouve/matryoshka-ui/components"
	. "github.com/LeonardJouve/matryoshka-ui/dsl"
)

// Settings is a settings page with tabs and grouped preference rows.
func Settings() *Node {
	return components.AppShell(
		components.Sidebar(180, "SETTINGS",
			components.NavItem("Account", true),
			components.NavItem("Notifications", false),
			components.NavItem("Billing", false),
			components.NavItem("Security", false),
		),
		components.Column(14,
			components.H2("Account Settings"),
			components.Tabs(0, "Profile", "Preferences", "Advanced"),

			components.TitledCard("Profile",
				components.Row(12,
					components.Field("Display name", "Ada Lovelace", ""),
					components.Field("Username", "ada", ""),
				),
				components.Field("Bio", "Mathematician & first programmer.", ""),
			),

			components.TitledCard("Notifications",
				components.ToggleRow("Email notifications", true),
				components.HDivider(),
				components.ToggleRow("Push notifications", false),
				components.HDivider(),
				components.ToggleRow("Weekly digest", true),
			),

			components.TitledCard("Theme",
				components.Row(10,
					components.CheckRow("Dark", true),
					components.CheckRow("Compact mode", false),
				),
			),

			components.Row(10, components.Spacer(), components.OutlineButton("Cancel", components.T.TextMute), components.Button("Save changes", components.T.Primary)),
		),
	)
}

// ProfileScreen shows a user header, stats, and a tabbed content area.
func ProfileScreen() *Node {
	header := Div(
		Style(
			LayoutAxis(LAYOUT_HORIZONTAL),
			Width(Grow(1)),
			Color(components.T.Surface),
			BorderRadius(0.1),
			Padding(PaddingAll(20)),
			Gap(GapHorizontal(16)),
			Align(AlignCenter),
		),
		Children(
			components.Avatar(64, components.T.Primary),
			Div(
				Style(LayoutAxis(LAYOUT_VERTICAL), Gap(GapVertical(4))),
				Children(
					components.H3("Ada Lovelace"),
					components.Muted("@ada · Joined 1843"),
					components.Row(6, components.Badge("Pro", components.T.Success), components.Badge("Verified", components.T.Primary)),
				),
			),
			components.Spacer(),
			components.Button("Follow", components.T.Primary),
		),
	)

	stats := components.Row(12,
		components.StatCard("Posts", "248", components.T.Text),
		components.StatCard("Followers", "12.4k", components.T.Primary),
		components.StatCard("Following", "183", components.T.Info),
	)

	return components.AppShell(
		components.Sidebar(180, "PROFILE",
			components.NavItem("Overview", true),
			components.NavItem("Posts", false),
			components.NavItem("Media", false),
		),
		components.Column(14,
			header,
			stats,
			components.Tabs(0, "Activity", "About", "Repos"),
			components.TitledCard("Recent Activity",
				components.ListRow("Pushed to analytical-engine", "2h", components.T.TextMute),
				components.HDivider(),
				components.ListRow("Opened issue #42", "5h", components.T.TextMute),
				components.HDivider(),
				components.ListRow("Commented on PR #7", "1d", components.T.TextMute),
			),
		),
	)
}
