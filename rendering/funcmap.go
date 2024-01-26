package rendering

import (
	"go.kyoto.codes/v3/htmx"
	"html/template"
	"strings"

	"go.kyoto.codes/v3/component"
	"go.kyoto.codes/zen/v3/errorsx"
	"go.kyoto.codes/zen/v3/mapx"
)

// FuncMap holds a library predefined template functions.
// You have to include it in your template building to use kyoto properly.
var FuncMap = template.FuncMap{
	// Inline render function.
	// Allows to avoid explicit template syntax
	// and customize render behavior.
	"render": func(f component.Future) template.HTML {
		// Await future
		state := f()
		// Check if state implements render
		if r, ok := state.(Renderer); ok {
			// Render
			var out strings.Builder
			errorsx.Must(0, r.Render(state, &out))
			// Pack and return
			return template.HTML(out.String())
		}
		// Panic if state does not implement render
		panic("state does not implement render")
	},
}

// FuncMapAll holds all funcmap instances of kyoto library.
var FuncMapAll = mapx.Merge(
	FuncMap,
	htmx.FuncMap,
	component.FuncMap,
)
