//go:build js && wasm

package main

import (
	"log"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"

	"vecty-portfolio/cmd/web/components"
	"vecty-portfolio/cmd/web/data"
)

type App struct {
	vecty.Core
	Resume *data.Resume
}


func (a *App) Render() vecty.ComponentOrHTML {
	return elem.Body(
		vecty.Markup(vecty.Class("app")),
		components.Navbar(),
		components.About(a.Resume),
		components.Experience(a.Resume),
		components.Projects(a.Resume),
		components.Skills(a.Resume),
		components.Footer(),
	)
}


func main() {
	resume, err := LoadResume()
	if err != nil {
		log.Fatal(err)
	}

	vecty.RenderBody(&App{
		Resume: resume,
	})

	select {}
}
