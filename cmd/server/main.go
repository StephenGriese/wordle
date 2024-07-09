package main

import (
	"errors"
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	h "github.com/maragudk/gomponents/html"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.Handle("GET /", getIndex())
	http.Handle("POST /", postForm())

	if err := http.ListenAndServe("localhost:8080", nil); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Println("Error:", err)
	}
}

func getIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		form := h.Form(
			h.Method("POST"),
			h.Label(h.For("name"), h.Input(h.Name("name"), h.ID("name"))),
			h.Br(),
			h.Button(g.Text("Submit")),
		)

		_ = Page("Welcome!", r.URL.Path, form).Render(w)
	}
}

func postForm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		possibleWords := []string{"hello", "world"}
		form := h.Form(
			h.Method("POST"),
			h.Label(h.For("name"), h.Input(h.Name("name"), h.ID("name"))),
			h.Br(),
			g.Text("Possible words:"),
			h.Br(),
			g.Text(strings.Join(possibleWords, " ")),
			h.Br(),
			h.Button(g.Text("Submit")),
		)

		_ = Page("Welcome!", r.URL.Path, form).Render(w)
	}
}

func Page(title, _ string, body g.Node) g.Node {
	// HTML5 boilerplate document
	return c.HTML5(c.HTML5Props{
		Title:    title,
		Language: "en",
		Head:     []g.Node{},
		Body: []g.Node{
			body,
		},
	})
}
