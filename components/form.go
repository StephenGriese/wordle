package components

import (
	g "github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// FormData represents the form input from the user
type FormData struct {
	Missed string
	Pos0   string
	Pos1   string
	Pos2   string
	Pos3   string
	Pos4   string
}

// WordleForm renders the main Wordle helper form
func WordleForm(data FormData, errorMsg string) g.Node {
	return html.Div(html.Class("row"),
		html.Div(html.Class("col-lg-8 mx-auto"),
			g.If(errorMsg != "", ErrorAlert(errorMsg)),
			InstructionsCard(),
			FormCard(data),
			html.Div(html.ID("results-section")),
		),
	)
}

// ErrorAlert renders an error message
func ErrorAlert(message string) g.Node {
	return html.Div(
		html.Class("alert alert-danger alert-dismissible fade show"),
		html.Role("alert"),
		html.Strong(g.Text("Error: ")),
		g.Text(message),
		html.Button(
			html.Type("button"),
			html.Class("btn-close"),
			g.Attr("data-bs-dismiss", "alert"),
			g.Attr("aria-label", "Close"),
		),
	)
}

// InstructionsCard renders the how-to-use instructions
func InstructionsCard() g.Node {
	return html.Div(html.Class("alert alert-info mb-4"),
		html.H5(html.Class("alert-heading"), g.Text("📖 How to use:")),
		html.Div(html.Class("instructions"),
			html.P(html.Class("mb-2"),
				html.Strong(g.Text("Missed Letters: ")),
				g.Text("Enter letters that are NOT in the word (gray squares ⬛)"),
			),
			html.P(html.Class("mb-2"), html.Strong(g.Text("Position boxes:"))),
			html.Ul(html.Class("mb-2"),
				html.Li(g.Text("Enter a letter if it's in the correct position (green square 🟩) - e.g., "), html.Code(g.Text("a"))),
				html.Li(g.Text("Enter "), html.Code(g.Text("-")), g.Text(" followed by letters if they're in the word but wrong position (yellow square 🟨) - e.g., "), html.Code(g.Text("-abc"))),
				html.Li(g.Text("Leave empty or use "), html.Code(g.Text(".")), g.Text(" if position is unknown")),
			),
			html.P(html.Class("mb-0"),
				html.Strong(g.Text("Example: ")),
				g.Text("If you tried \"CRANE\" and got: C(gray), R(yellow), A(green), N(gray), E(gray)"),
				html.Br(),
				g.Text("→ Missed: "), html.Code(g.Text("cne")),
				g.Text(", Position 1: "), html.Code(g.Text(".")),
				g.Text(", Position 2: "), html.Code(g.Text("-r")),
				g.Text(", Position 3: "), html.Code(g.Text("a")),
				g.Text(", Position 4: "), html.Code(g.Text(".")),
				g.Text(", Position 5: "), html.Code(g.Text(".")),
			),
		),
	)
}

// FormCard renders the input form
func FormCard(data FormData) g.Node {
	return html.Div(html.Class("form-card"),
		html.Form(
			html.Method("POST"),
			html.Action("/wordle/solve"),
			g.Attr("hx-post", "/wordle/solve"),
			g.Attr("hx-target", "#results-section"),
			g.Attr("hx-indicator", "#loading"),

			// Missed Letters field
			html.Div(html.Class("mb-4"),
				html.Label(html.For("missed"), html.Class("form-label fw-bold"), g.Text("Missed Letters (not in word)")),
				html.Input(
					html.Type("text"),
					html.Class("form-control"),
					html.ID("missed"),
					html.Name("missed"),
					html.Value(data.Missed),
					html.Placeholder("e.g., xyz"),
					g.Attr("autocomplete", "off"),
				),
				html.Div(html.Class("form-text"), g.Text("Enter all letters that appeared gray (not in the word)")),
			),

			// Position inputs
			html.Div(html.Class("mb-4"),
				html.Label(html.Class("form-label fw-bold"), g.Text("Word Positions (1-5)")),
				html.Div(html.Class("d-flex justify-content-center gap-3 flex-wrap"),
					PositionInput("pos0", data.Pos0, "1"),
					PositionInput("pos1", data.Pos1, "2"),
					PositionInput("pos2", data.Pos2, "3"),
					PositionInput("pos3", data.Pos3, "4"),
					PositionInput("pos4", data.Pos4, "5"),
				),
				html.Div(html.Class("form-text text-center mt-2"),
					g.Text("Green (correct): "), html.Code(g.Text("a")),
					g.Text(" | Yellow (wrong position): "), html.Code(g.Text("-abc")),
					g.Text(" | Unknown: "), html.Code(g.Text(".")),
				),
			),

			// Submit button
			html.Div(html.Class("text-center"),
				html.Button(
					html.Type("submit"),
					html.Class("btn btn-solve btn-lg"),
					g.Text("Find Possible Words"),
				),
				html.Div(html.ID("loading"), html.Class("htmx-indicator ms-2"),
					html.Div(html.Class("spinner-border text-success"), html.Role("status"),
						html.Span(html.Class("visually-hidden"), g.Text("Loading...")),
					),
				),
			),
		),
	)
}

// PositionInput renders a single position input box
func PositionInput(name, value, label string) g.Node {
	return html.Div(
		html.Input(
			html.Type("text"),
			html.Class("form-control position-input"),
			html.Name(name),
			html.Value(value),
			g.Attr("maxlength", "10"),
			html.Placeholder("."),
			g.Attr("autocomplete", "off"),
		),
		html.Div(html.Class("position-label"), g.Text(label)),
	)
}
