package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"

	"vecty-portfolio/cmd/web/data"
)

func Hero(r *data.Resume) vecty.ComponentOrHTML {
	return elem.Section(
		vecty.Markup(vecty.Class("hero")),
		elem.Heading1(
			vecty.Text(r.Basics.Name),
		),
		elem.Paragraph(
			vecty.Text("DevOps Engineer · Golang · AWS · Kubernetes"),
		),
		elem.Div(
			vecty.Markup(vecty.Class("hero-links")),
			elem.Anchor(
				vecty.Markup(
					vecty.Attribute("href", r.Basics.URL.Href),
					vecty.Attribute("target", "_blank"),
				),
				vecty.Text("LinkedIn"),
			),
			elem.Anchor(
				vecty.Markup(
					vecty.Attribute("href", "https://github.com/octocamocoder47"),
					vecty.Attribute("target", "_blank"),
				),
				vecty.Text("GitHub"),
			),
		),
	)
}
