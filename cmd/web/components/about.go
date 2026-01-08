package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"

	"vecty-portfolio/cmd/web/data"
)

func About(r *data.Resume) vecty.ComponentOrHTML {
	var custom_fields []vecty.MarkupOrChild

	for _, f := range r.Basics.Fields {
		var field vecty.MarkupOrChild
		field = elem.Anchor(
			// vecty.Markup(
			// 	vecty.Attribute("icon", f.ICON),
			// ),
			vecty.Markup(
				vecty.Attribute("href", f.VALUE),
				vecty.Attribute("target", "_blank"),
			),
			vecty.Text(f.NAME),
		)
		custom_fields = append(custom_fields, field)
	}

	return elem.Section(
		vecty.Markup(vecty.Class("about"), vecty.Attribute("id", "about"),),
		elem.Heading1(
			vecty.Text(r.Basics.Name),
		),
		elem.Paragraph(
			vecty.Text(r.Basics.Headline),
		),
		elem.Div(
			append(
				[]vecty.MarkupOrChild{
					vecty.Markup(vecty.Class("about-links")),
				},
				custom_fields...
			)...
		),
	)
}
