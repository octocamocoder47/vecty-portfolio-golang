package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"vecty-portfolio/cmd/web/data"
)

func Experience(r *data.Resume) vecty.ComponentOrHTML {
	var cards []vecty.MarkupOrChild

	for _, e := range r.Sections.Experience.Items {
		cards = append(cards,
			elem.Div(
				vecty.Markup(vecty.Class("card")),
				elem.Heading3(vecty.Text(e.Position)),
				elem.Small(vecty.Text(e.Company+" · "+e.Date)),
				elem.Div(
					vecty.Markup(vecty.UnsafeHTML(e.Summary)),
				),
			),
		)
	}

	return elem.Section(
		vecty.Markup(vecty.Class("section"), vecty.Attribute("id", "experience")),
		elem.Heading2(vecty.Text("Experience")),
		elem.Div(
			append(
				[]vecty.MarkupOrChild{
					vecty.Markup(vecty.Class("card-grid")),
				},
				cards...,
			)...,
		),
	)
}
