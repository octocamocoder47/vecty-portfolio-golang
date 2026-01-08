package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

func Footer() vecty.ComponentOrHTML {
	return elem.Footer(
		vecty.Markup(vecty.Class("footer")),
		elem.Paragraph(
			vecty.Text("© 2026 Pransh Gupta · Built with Go + Vecty"),
		),
	)
}
