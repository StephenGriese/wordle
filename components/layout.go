package components

import (
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

// Page creates the full HTML page layout with Bootstrap
func Page(title string, content g.Node) g.Node {
	return c.HTML5(c.HTML5Props{
		Title:    title,
		Language: "en",
		Head: []g.Node{
			html.Link(html.Href("https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css"), html.Rel("stylesheet")),
			html.Script(html.Src("/static/js/htmx-1.9.11.js")),
			html.StyleEl(g.Raw(pageStyles)),
		},
		Body: []g.Node{
			PageHeader(),
			html.Div(html.Class("container"),
				content,
			),
			PageFooter(),
			html.Script(html.Src("https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/js/bootstrap.bundle.min.js")),
		},
	})
}

// PageHeader renders the page header with reload button
func PageHeader() g.Node {
	return html.Div(html.Class("wordle-header"),
		html.Div(html.Class("container"),
			html.Div(html.Class("d-flex justify-content-between align-items-center"),
				html.Div(
					html.H1(g.Text("🎯 Wordle Helper")),
					html.P(html.Class("text-muted mb-0"), g.Text("Find possible words based on your Wordle clues")),
				),
				html.Div(html.Class("text-end"),
					html.Div(html.Class("reload-container"),
						html.Button(
							html.Class("btn btn-reload btn-sm"),
							g.Attr("hx-post", "/reload"),
							g.Attr("hx-target", "#reload-message"),
							g.Attr("hx-swap", "innerHTML"),
							g.Attr("hx-disabled-elt", "this"),
							g.Attr("hx-indicator", "#reload-spinner"),
							g.Text("🔄 Reload Word List"),
						),
						html.Div(
							html.ID("reload-spinner"),
							html.Class("spinner-border spinner-border-sm text-info"),
							html.Role("status"),
							html.Span(html.Class("visually-hidden"), g.Text("Reloading...")),
						),
					),
					html.Div(html.ID("reload-message")),
				),
			),
		),
	)
}

// PageFooter renders the page footer
func PageFooter() g.Node {
	return html.Footer(html.Class("container mt-5"),
		html.Div(html.Class("text-center text-muted"),
			html.Small(g.Text("Enter your Wordle clues to find all possible matching words")),
		),
	)
}

const pageStyles = `
:root {
    --wordle-green: #538d4e;
    --wordle-yellow: #b59f3b;
    --wordle-gray: #787c7e;
    --wordle-light-gray: #d3d6da;
}

body {
    background-color: #f8f9fa;
    padding: 20px 0;
}

.wordle-header {
    background-color: white;
    border-bottom: 1px solid var(--wordle-light-gray);
    padding: 20px 0;
    margin-bottom: 30px;
}

.wordle-header h1 {
    color: #212529;
    font-weight: bold;
    margin: 0;
}

.form-card {
    background: white;
    border-radius: 8px;
    box-shadow: 0 2px 4px rgba(0,0,0,0.1);
    padding: 30px;
    margin-bottom: 20px;
}

.position-input {
    width: 60px;
    height: 60px;
    text-align: center;
    font-size: 24px;
    font-weight: bold;
    text-transform: lowercase;
    border: 2px solid var(--wordle-light-gray);
}

.position-input:focus {
    border-color: var(--wordle-gray);
    box-shadow: 0 0 0 0.2rem rgba(120,124,126,0.25);
}

.position-label {
    font-size: 12px;
    color: #6c757d;
    margin-top: 5px;
    text-align: center;
}

.btn-solve {
    background-color: var(--wordle-green);
    border-color: var(--wordle-green);
    color: white;
    font-weight: bold;
    padding: 12px 40px;
    font-size: 18px;
}

.btn-solve:hover {
    background-color: #446f41;
    border-color: #446f41;
    color: white;
}

.results-card {
    background: white;
    border-radius: 8px;
    box-shadow: 0 2px 4px rgba(0,0,0,0.1);
    padding: 30px;
}

.word-badge {
    display: inline-block;
    padding: 8px 12px;
    margin: 5px;
    background-color: #e9ecef;
    border-radius: 4px;
    font-size: 16px;
    font-weight: 500;
}

.alert-info {
    border-left: 4px solid #0dcaf0;
}

.instructions {
    font-size: 14px;
    color: #6c757d;
}

.htmx-indicator {
    display: none;
}

.htmx-request .htmx-indicator {
    display: inline-block;
}

.htmx-request .btn-solve {
    opacity: 0.6;
}

.btn-reload {
    background-color: #0dcaf0;
    border-color: #0dcaf0;
    color: white;
    font-size: 14px;
    padding: 6px 12px;
}

.btn-reload:hover:not(:disabled) {
    background-color: #0aa2c0;
    border-color: #0aa2c0;
    color: white;
}

.btn-reload:disabled {
    opacity: 0.6;
    cursor: not-allowed;
}

.reload-container {
    display: inline-flex;
    align-items: center;
    gap: 8px;
}

#reload-spinner {
    display: none;
}

#reload-spinner.htmx-request {
    display: inline-block;
}

#reload-message {
    margin-top: 10px;
}
`
