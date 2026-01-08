package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

func Navbar() vecty.ComponentOrHTML {
	return elem.Navigation(
		vecty.Markup(vecty.Class("navbar")),
		elem.Div(
			vecty.Markup(vecty.Class("nav-inner")),
			elem.Span(
				vecty.Markup(vecty.Class("nav-brand")),
				vecty.Text("Pransh Gupta"),
			),
			elem.Div(
				vecty.Markup(vecty.Class("nav-links")),
				elem.Anchor(vecty.Markup(vecty.Attribute("href", "#experience")), vecty.Text("Experience")),
				elem.Anchor(vecty.Markup(vecty.Attribute("href", "#projects")), vecty.Text("Projects")),
				elem.Anchor(vecty.Markup(vecty.Attribute("href", "#skills")), vecty.Text("Skills")),
			),
		),
	)
}
