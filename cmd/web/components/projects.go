package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"

	"vecty-portfolio/cmd/web/data"
)

func Projects(r *data.Resume) vecty.ComponentOrHTML {
	var cards []vecty.MarkupOrChild

	for _, p := range r.Sections.Projects.Items {

		var link vecty.MarkupOrChild
		if p.URL != nil && p.URL.Href != "" {
			link = elem.Anchor(
				vecty.Markup(
					vecty.Attribute("href", p.URL.Href),
					vecty.Attribute("target", "_blank"),
					vecty.Class("project-link"),
				),
				vecty.Text("View Project →"),
			)
		}

		cards = append(cards,
			elem.Div(
				vecty.Markup(vecty.Class("card")),
				elem.Heading3(vecty.Text(p.Name)),
				elem.Div(
					vecty.Markup(vecty.UnsafeHTML(p.Summary)),
				),
				link,
			),
		)
	}

	return elem.Section(
		vecty.Markup(
			vecty.Class("section"),
			vecty.Attribute("id", "projects"),
		),
		elem.Heading2(vecty.Text("Projects")),
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
