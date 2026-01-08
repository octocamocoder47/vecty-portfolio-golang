package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"vecty-portfolio/cmd/web/data"
)

func Header(r *data.Resume) vecty.ComponentOrHTML {
	return elem.Header(
		vecty.Markup(
			vecty.Class("header"),
		),
		elem.Heading1(vecty.Text(r.Basics.Name)),
		elem.Paragraph(
			vecty.Text("DevOps Engineer · Go · AWS · Kubernetes"),
		),
		elem.Anchor(
			vecty.Markup(
				vecty.Attribute("href", r.Basics.URL.Href),
				vecty.Attribute("target", "_blank"),
			),
			vecty.Text("LinkedIn"),
		),
	)
}
