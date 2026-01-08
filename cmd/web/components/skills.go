package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"vecty-portfolio/cmd/web/data"
)

func Skills(r *data.Resume) vecty.ComponentOrHTML {
	items := []vecty.MarkupOrChild{}

	for _, s := range r.Sections.Skills.Items {
		items = append(items,
			elem.Span(
				vecty.Markup(vecty.Class("skill")),
				vecty.Text(s.Name),
			),
		)
	}

	args := []vecty.MarkupOrChild{vecty.Markup(vecty.Class("skills"), vecty.Attribute("id", "skills"),)}
	args = append(args, items...)

	return elem.Section(
		vecty.Markup(vecty.Class("section")),
		elem.Heading2(vecty.Text("Skills")),
		elem.Div(args...),
	)
}
