package components

import (
	"github.com/LeonardJouve/matryoshka-ui/components"
	. "github.com/LeonardJouve/matryoshka-ui/dsl"
)

// InboxScreen is an email-style list with a toolbar and rows.
func Inbox() *Node {
	return components.AppShell(
		components.Sidebar(180, "MAIL",
			components.NavItem("Inbox", true),
			components.NavItem("Starred", false),
			components.NavItem("Sent", false),
			components.NavItem("Drafts", false),
			components.NavItem("Trash", false),
		),
		components.Column(14,
			components.Row(0, components.H2("Inbox"), components.Spacer(), components.Badge("12 unread", components.T.Primary), components.Button("Compose", components.T.Primary)),
			components.Tabs(0, "Primary", "Social", "Promotions"),
			components.List(
				mailRow("Grace Hopper", "Compiler results are in", "9:14", true),
				mailRow("Alan Turing", "Re: decidability proof", "8:02", true),
				mailRow("Charles Babbage", "Engine schematics v3", "Yesterday", false),
				mailRow("Howard Aiken", "Mark I status report", "Mon", false),
				mailRow("John von Neumann", "Architecture draft", "Sun", false),
			),
		),
	)
}

func mailRow(sender string, subject string, time string, unread bool) Element {
	senderColor := components.T.TextMute
	if unread {
		senderColor = components.T.Text
	}
	lead := components.Dot(components.T.Bg)
	if unread {
		lead = components.Dot(components.T.Primary)
	}
	return Div(
		Style(LayoutAxis(LAYOUT_HORIZONTAL), Width(Grow(1)), Align(AlignCenter), Gap(GapHorizontal(12)), Padding(components.Padding2(8, 6))),
		Children(
			lead,
			Div(Style(Width(Fixed(140))), Children(Text(sender, TextStyle(FontSize(13), Color(senderColor))))),
			Div(Style(Width(Grow(1))), Children(components.Muted(subject))),
			components.Caption(time),
		),
	)
}

// AnalyticsScreen pairs KPI cards, a bar chart, and a data table.
func AnalyticsScreen() *Node {
	return components.AppShell(
		components.Sidebar(180, "ANALYTICS",
			components.NavItem("Traffic", true),
			components.NavItem("Conversions", false),
			components.NavItem("Revenue", false),
		),
		components.Column(14,
			components.Row(0, components.H2("Traffic Overview"), components.Spacer(), components.Tabs(1, "Day", "Week", "Month")),
			components.Row(12,
				components.StatCard("Visitors", "94.2k", components.T.Primary),
				components.StatCard("Page views", "312k", components.T.Info),
				components.StatCard("Avg. session", "3m 12s", components.T.Success),
				components.StatCard("Bounce", "38%", components.T.Warning),
			),
			components.TitledCard("Visitors by day",
				components.BarChart(180, components.T.Primary, 60, 90, 70, 130, 100, 160, 140),
			),
			components.TitledCard("Top pages",
				components.Table(
					components.TableRow(true, "Page", "Views", "Unique", "Bounce"),
					components.TableRow(false, "/home", "42,108", "31,200", "32%"),
					components.TableRow(false, "/pricing", "18,943", "15,002", "41%"),
					components.TableRow(false, "/docs", "12,330", "9,870", "28%"),
					components.TableRow(false, "/blog", "8,221", "6,540", "55%"),
				),
			),
		),
	)
}
