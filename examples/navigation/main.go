package main

import "raitui/theme"

func main() {
	tb1 := Tab("First").BackgroundColor(theme.White)
	tb1.ChildAt(0).TextColor(theme.Gray800)
	tb2 := Tab("Second")
	tb3 := Tab("Third")

	p1 := TabPanel().Children(Text("First panel content").TextColor(theme.Gray600).FontSize(14))
	p2 := TabPanel().Children(Text("Second panel content").TextColor(theme.Gray600).FontSize(14))
	p3 := TabPanel().Children(Text("Third panel content").TextColor(theme.Gray600).FontSize(14))
	p1.Visible(true)

	resetTabs := func() {
		tb1.BackgroundColor(theme.Transparent); tb1.ChildAt(0).TextColor(theme.Gray500)
		tb2.BackgroundColor(theme.Transparent); tb2.ChildAt(0).TextColor(theme.Gray500)
		tb3.BackgroundColor(theme.Transparent); tb3.ChildAt(0).TextColor(theme.Gray500)
		p1.Visible(false); p2.Visible(false); p3.Visible(false)
	}

	tb1.OnClick(func() { resetTabs(); tb1.BackgroundColor(theme.White); tb1.ChildAt(0).TextColor(theme.Gray800); p1.Visible(true) })
	tb2.OnClick(func() { resetTabs(); tb2.BackgroundColor(theme.White); tb2.ChildAt(0).TextColor(theme.Gray800); p2.Visible(true) })
	tb3.OnClick(func() { resetTabs(); tb3.BackgroundColor(theme.White); tb3.ChildAt(0).TextColor(theme.Gray800); p3.Visible(true) })

	acc1 := AccordionItem()
	acc1h := AccordionHeader("What is RaitUI?")
	acc1p := AccordionPanel().Children(Text("RaitUI is a Go GUI framework using Ebitengine.").TextColor(theme.Gray600).FontSize(14))
	acc1h.OnClick(func() { acc1p.Visible(!acc1p.IsVisible()) })
	acc1.Children(acc1h, acc1p)

	acc2 := AccordionItem()
	acc2h := AccordionHeader("How does it work?")
	acc2p := AccordionPanel().Children(Text("It uses goda for layout and ebitengine for rendering.").TextColor(theme.Gray600).FontSize(14))
	acc2h.OnClick(func() { acc2p.Visible(!acc2p.IsVisible()) })
	acc2.Children(acc2h, acc2p)

	app := App().
		Padding("24").Gap("16").
		BackgroundColor(theme.Gray50).
		Children(
			Text("Navigation Components").TextColor(theme.Gray800).FontSize(18),
			Separator().Width("100%"),

			Text("Breadcrumb").TextColor(theme.Gray400).FontSize(11),
			Breadcrumb().Children(
				BreadcrumbLink("Home"),
				BreadcrumbSeparator(),
				BreadcrumbLink("Products"),
				BreadcrumbSeparator(),
				BreadcrumbItem("Details"),
			),
			Separator().Width("100%"),

			Text("Tabs").TextColor(theme.Gray400).FontSize(11),
			Tabs().Children(
				TabList().Children(tb1, tb2, tb3),
				TabPanels().Children(p1, p2, p3),
			),
			Separator().Width("100%"),

			Text("Accordion").TextColor(theme.Gray400).FontSize(11),
			Accordion().Children(acc1, acc2),
		)

	Window().Title("RaitUI - Navigation").Width(700).Height(560).MinSize(400, 300).Children(app).Run()
}
