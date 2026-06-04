package example

import (
	. "github.com/LeonardJouve/matryoshka-ui/dsl"
	"github.com/LeonardJouve/matryoshka-ui/utils"
)

func Dashboard() *Node {
	var (
		bg     = utils.Col(16, 18, 27)
		card   = utils.Col(28, 31, 46)
		border = utils.Col(42, 48, 68)

		white = utils.Col(235, 238, 248)
		muted = utils.Col(140, 146, 175)

		green  = utils.Col(0, 255, 120)
		blue   = utils.Col(90, 160, 255)
		orange = utils.Col(255, 170, 80)
	)

	kpiCard := func(title string, value string, accent utils.Color) Element {
		return Div(
			Style(
				LayoutAxis(LAYOUT_VERTICAL),
				Padding(PaddingAll(14)),
				Gap(GapAll(6)),
				Color(card),
				BorderRadius(0.1),
				BorderColor(border),
				BorderWidth(1),
				Width(Grow(1)),
			),
			Children(
				Text(title, TextStyle(FontSize(12), Color(muted))),
				Text(value, TextStyle(FontSize(20), Color(accent))),
			),
		)
	}

	return Root(
		Div(
			Style(
				Width(Grow(1)),
				Height(Grow(1)),
				LayoutAxis(LAYOUT_HORIZONTAL),
				Color(bg),
				Padding(PaddingAll(18)),
				Gap(GapAll(16)),
			),
			Children(

				// LEFT SIDEBAR
				Div(
					Style(
						LayoutAxis(LAYOUT_VERTICAL),
						Width(Fixed(180)),
						Color(card),
						BorderRadius(0.1),
						Padding(PaddingAll(12)),
						Gap(GapAll(10)),
					),
					Children(
						Text("DASHBOARD", TextStyle(FontSize(14), Color(white))),
						Text("Overview", TextStyle(FontSize(12), Color(muted))),
						Text("Analytics", TextStyle(FontSize(12), Color(muted))),
						Text("Revenue", TextStyle(FontSize(12), Color(muted))),
						Text("Users", TextStyle(FontSize(12), Color(muted))),
						Text("Settings", TextStyle(FontSize(12), Color(muted))),
					),
				),

				// MAIN CONTENT
				Div(
					Style(
						LayoutAxis(LAYOUT_VERTICAL),
						Width(Grow(1)),
						Gap(GapAll(14)),
					),
					Children(

						// KPI ROW
						Div(
							Style(
								LayoutAxis(LAYOUT_HORIZONTAL),
								Gap(GapAll(12)),
								Width(Grow(1)),
							),
							Children(
								kpiCard("Revenue", "$42.3k", green),
								kpiCard("Users", "8,431", blue),
								kpiCard("Orders", "1,248", orange),
								kpiCard("Churn", "2.4%", utils.Col(255, 80, 80)),
							),
						),

						// MAIN CHART + SIDE PANEL
						Div(
							Style(
								LayoutAxis(LAYOUT_HORIZONTAL),
								Gap(GapAll(12)),
								Width(Grow(1)),
							),
							Children(

								// MAIN BAR CHART AREA
								Div(
									Style(
										LayoutAxis(LAYOUT_VERTICAL),
										Width(Grow(2)),
										Color(card),
										BorderRadius(0.1),
										Padding(PaddingAll(14)),
										Gap(GapAll(10)),
									),
									Children(
										Text("Monthly Revenue",
											TextStyle(FontSize(16), Color(white)),
										),

										// fake bar chart blocks
										Div(
											Style(
												LayoutAxis(LAYOUT_HORIZONTAL),
												Align(AlignEnd),
												Gap(GapHorizontal(6)),
												Height(Fixed(160)),
											),
											Children(
												bar(50, green),
												bar(80, green),
												bar(60, green),
												bar(120, green),
												bar(90, green),
												bar(150, green),
												bar(110, green),
											),
										),
									),
								),

								// SIDE PANEL (ACTIVITY)
								Div(
									Style(
										LayoutAxis(LAYOUT_VERTICAL),
										Width(Grow(1)),
										Color(card),
										BorderRadius(0.1),
										Padding(PaddingAll(14)),
										Gap(GapAll(10)),
									),
									Children(
										Text("Recent Activity",
											TextStyle(FontSize(14), Color(white)),
										),

										Text("- New user signup", TextStyle(FontSize(12), Color(muted))),
										Text("- Invoice paid", TextStyle(FontSize(12), Color(muted))),
										Text("- Deployment success", TextStyle(FontSize(12), Color(muted))),
										Text("- Stripe payout sent", TextStyle(FontSize(12), Color(muted))),
									),
								),
							),
						),

						// LOWER ROW (2 MINI CHARTS)
						Div(
							Style(
								LayoutAxis(LAYOUT_HORIZONTAL),
								Gap(GapAll(12)),
								Width(Grow(1)),
							),
							Children(

								smallChart("Active Users", blue),
								smallChart("Conversion Rate", orange),
							),
						),
					),
				),
			),
		),
	)
}

// helpers

func bar(h uint16, c utils.Color) Element {
	return Div(
		Style(
			Width(Fixed(18)),
			Height(Fixed(h)),
			Color(c),
			BorderRadius(0.2),
		),
	)
}

func smallChart(title string, c utils.Color) Element {
	return Div(
		Style(
			LayoutAxis(LAYOUT_VERTICAL),
			Width(Grow(1)),
			Color(utils.Col(28, 31, 46)),
			BorderRadius(0.1),
			Padding(PaddingAll(12)),
			Gap(GapAll(8)),
		),
		Children(
			Text(title, TextStyle(FontSize(13), Color(utils.Col(235, 238, 248)))),

			Div(
				Style(
					LayoutAxis(LAYOUT_HORIZONTAL),
					Align(AlignEnd),
					Gap(GapHorizontal(4)),
					Height(Fixed(80)),
				),
				Children(
					bar(30, c),
					bar(55, c),
					bar(40, c),
					bar(70, c),
					bar(60, c),
				),
			),
		),
	)
}
