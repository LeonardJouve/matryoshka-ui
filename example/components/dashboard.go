package components

import (
	"github.com/LeonardJouve/matryoshka-ui/components"
	. "github.com/LeonardJouve/matryoshka-ui/dsl"
)

func Dashboard() *Node {
	return components.AppShell(
		components.Sidebar(180, "DASHBOARD",
			components.NavItem("Overview", true),
			components.NavItem("Analytics", false),
			components.NavItem("Revenue", false),
			components.NavItem("Users", false),
			components.NavItem("Settings", false),
		),
		components.Column(14,
			// top bar
			components.Navbar("Acme Inc",
				components.Badge("LIVE", components.T.Success),
				components.PrimaryButton("New Report"),
			),

			// KPI row
			components.Row(12,
				components.StatCard("Revenue", "$42.3k", components.T.Success),
				components.StatCard("Users", "8,431", components.T.Primary),
				components.StatCard("Orders", "1,248", components.T.Warning),
				components.StatCard("Churn", "2.4%", components.T.Danger),
			),

			// chart + side panel
			components.Row(12,
				components.TitledCard("Monthly Revenue",
					components.BarChart(160, components.T.Success, 50, 80, 60, 120, 90, 150, 110),
				),
				components.Panel(
					components.H3("Recent Activity"),
					components.HDivider(),
					components.Muted("- New user signup"),
					components.Muted("- Invoice paid"),
					components.Muted("- Deployment success"),
				),
			),

			// detail row
			components.Row(12,
				components.Card(
					components.H3("Plan Usage"),
					components.KeyValue("Current Plan", "Pro"),
					components.Label("Storage"),
					components.ProgressBar(0.62, components.T.Warning),
					components.Label("Bandwidth"),
					components.ProgressBar(0.31, components.T.Primary),
				),
				components.Card(
					components.H3("Team"),
					components.Row(8,
						components.Avatar(32, components.T.Primary),
						components.Avatar(32, components.T.Success),
						components.Avatar(32, components.T.Warning),
					),
					components.ListRow("Online", "3", components.T.Success),
					components.ListRow("Away", "1", components.T.Warning),
					components.Row(8,
						components.OutlineButton("Manage", components.T.Info),
						components.Button("Invite", components.T.Primary),
					),
				),
			),
		),
	)
}
