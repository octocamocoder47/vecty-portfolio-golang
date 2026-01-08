package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"vecty-portfolio/cmd/web/data"
)

func Summary(r *data.Resume) vecty.ComponentOrHTML {
	return elem.Section(
		vecty.Markup(vecty.Class("section")),
		elem.Heading2(vecty.Text("Summary")),
		elem.Div(
			vecty.Markup(
				vecty.UnsafeHTML(r.Sections.Summary.Content),
			),
		),
	)
}
