package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	h "github.com/maragudk/gomponents/html"
)

func main() {
	http.Handle("/", createHandler(indexPage()))
	http.Handle("/contact", createHandler(contactPage()))
	http.Handle("/about", createHandler(aboutPage()))

	if err := http.ListenAndServe("localhost:8080", nil); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Println("Error:", err)
	}
}

func createHandler(title string, body g.Node) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Rendering a Node is as simple as calling Render and passing an io.Writer
		_ = Page(title, r.URL.Path, body).Render(w)
	}
}

func indexPage() (string, g.Node) {
	return "Welcome!", h.Div(
		h.H1(g.Text("Welcome to this example page")),
		h.P(g.Text("I hope it will make you happy. 😄 It's using TailwindCSS for styling.")),
	)
}

func contactPage() (string, g.Node) {
	return "Contact", h.Div(
		h.H1(g.Text("Contact us")),
		h.P(g.Text("Just do it.")),
	)
}

func aboutPage() (string, g.Node) {
	return "About", h.Div(
		h.H1(g.Text("About this site")),
		h.P(g.Text("This is a site showing off gomponents.")),
	)
}

func Page(title, path string, body g.Node) g.Node {
	// HTML5 boilerplate document
	return c.HTML5(c.HTML5Props{
		Title:    title,
		Language: "en",
		Head: []g.Node{
			h.Script(h.Src("https://cdn.tailwindcss.com?plugins=typography")),
		},
		Body: []g.Node{
			Navbar(path, []PageLink{
				{Path: "/contact", Name: "Contact"},
				{Path: "/about", Name: "About"},
			}),
			Container(
				Prose(body),
				PageFooter(),
			),
		},
	})
}

type PageLink struct {
	Path string
	Name string
}

func Navbar(currentPath string, links []PageLink) g.Node {
	return h.Nav(h.Class("bg-gray-700 mb-4"),
		Container(
			h.Div(h.Class("flex items-center space-x-4 h-16"),
				NavbarLink("/", "Home", currentPath == "/"),

				// We can Map custom slices to Nodes
				g.Group(g.Map(links, func(l PageLink) g.Node {
					return NavbarLink(l.Path, l.Name, currentPath == l.Path)
				})),
			),
		),
	)
}

// NavbarLink is a link in the Navbar.
func NavbarLink(path, text string, active bool) g.Node {
	return h.A(h.Href(path), g.Text(text),
		// Apply CSS classes conditionally
		c.Classes{
			"px-3 py-2 rounded-md text-sm font-medium focus:outline-none focus:text-white focus:bg-gray-700": true,
			"text-white bg-gray-900":                           active,
			"text-gray-300 hover:text-white hover:bg-gray-700": !active,
		},
	)
}

func Container(children ...g.Node) g.Node {
	return h.Div(h.Class("max-w-7xl mx-auto px-2 sm:px-6 lg:px-8"), g.Group(children))
}

func Prose(children ...g.Node) g.Node {
	return h.Div(h.Class("prose"), g.Group(children))
}

func PageFooter() g.Node {
	return h.Footer(h.Class("prose prose-sm prose-indigo"),
		h.P(
			// We can use string interpolation directly, like fmt.Sprintf.
			g.Textf("Rendered %v. ", time.Now().Format(time.RFC3339)),

			// Conditional inclusion
			g.If(time.Now().Second()%2 == 0, g.Text("It's an even second.")),
			g.If(time.Now().Second()%2 == 1, g.Text("It's an odd second.")),
		),

		h.P(h.ID("paragraph1"), h.A(h.Href("https://www.gomponents.com"), g.Text("gomponents"))),
	)
}
