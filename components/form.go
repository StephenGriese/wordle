package components

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
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
	return Div(Class("row"),
		Div(Class("col-lg-8 mx-auto"),
			g.If(errorMsg != "", ErrorAlert(errorMsg)),
			InstructionsCard(),
			FormCard(data),
			Div(ID("results-section")),
		),
	)
}

// ErrorAlert renders an error message
func ErrorAlert(message string) g.Node {
	return Div(
		Class("alert alert-danger alert-dismissible fade show"),
		Role("alert"),
		Strong(g.Text("Error: ")),
		g.Text(message),
		Button(
			Type("button"),
			Class("btn-close"),
			g.Attr("data-bs-dismiss", "alert"),
			g.Attr("aria-label", "Close"),
		),
	)
}

// InstructionsCard renders the how-to-use instructions
func InstructionsCard() g.Node {
	return Div(Class("alert alert-info mb-4"),
		H5(Class("alert-heading"), g.Text("📖 How to use:")),
		Div(Class("instructions"),
			P(Class("mb-2"),
				Strong(g.Text("Missed Letters: ")),
				g.Text("Enter letters that are NOT in the word (gray squares ⬛)"),
			),
			P(Class("mb-2"), Strong(g.Text("Position boxes:"))),
			Ul(Class("mb-2"),
				Li(g.Text("Enter a letter if it's in the correct position (green square 🟩) - e.g., "), Code(g.Text("a"))),
				Li(g.Text("Enter "), Code(g.Text("-")), g.Text(" followed by letters if they're in the word but wrong position (yellow square 🟨) - e.g., "), Code(g.Text("-abc"))),
				Li(g.Text("Leave empty or use "), Code(g.Text(".")), g.Text(" if position is unknown")),
			),
			P(Class("mb-0"),
				Strong(g.Text("Example: ")),
				g.Text("If you tried \"CRANE\" and got: C(gray), R(yellow), A(green), N(gray), E(gray)"),
				Br(),
				g.Text("→ Missed: "), Code(g.Text("cne")),
				g.Text(", Position 1: "), Code(g.Text(".")),
				g.Text(", Position 2: "), Code(g.Text("-r")),
				g.Text(", Position 3: "), Code(g.Text("a")),
				g.Text(", Position 4: "), Code(g.Text(".")),
				g.Text(", Position 5: "), Code(g.Text(".")),
			),
		),
	)
}

// FormCard renders the input form
func FormCard(data FormData) g.Node {
	return Div(Class("form-card"),
		Form(
			Method("POST"),
			Action("/wordle/solve"),
			g.Attr("hx-post", "/wordle/solve"),
			g.Attr("hx-target", "#results-section"),
			g.Attr("hx-indicator", "#loading"),

			// Missed Letters field
			Div(Class("mb-4"),
				Label(For("missed"), Class("form-label fw-bold"), g.Text("Missed Letters (not in word)")),
				Input(
					Type("text"),
					Class("form-control"),
					ID("missed"),
					Name("missed"),
					Value(data.Missed),
					Placeholder("e.g., xyz"),
					g.Attr("autocomplete", "off"),
				),
				Div(Class("form-text"), g.Text("Enter all letters that appeared gray (not in the word)")),
			),

			// Position inputs
			Div(Class("mb-4"),
				Label(Class("form-label fw-bold"), g.Text("Word Positions (1-5)")),
				Div(Class("d-flex justify-content-center gap-3 flex-wrap"),
					PositionInput("pos0", data.Pos0, "1"),
					PositionInput("pos1", data.Pos1, "2"),
					PositionInput("pos2", data.Pos2, "3"),
					PositionInput("pos3", data.Pos3, "4"),
					PositionInput("pos4", data.Pos4, "5"),
				),
				Div(Class("form-text text-center mt-2"),
					g.Text("Green (correct): "), Code(g.Text("a")),
					g.Text(" | Yellow (wrong position): "), Code(g.Text("-abc")),
					g.Text(" | Unknown: "), Code(g.Text(".")),
				),
			),

			// Submit button
			Div(Class("text-center"),
				Button(
					Type("submit"),
					Class("btn btn-solve btn-lg"),
					g.Text("Find Possible Words"),
				),
				Div(ID("loading"), Class("htmx-indicator ms-2"),
					Div(Class("spinner-border text-success"), Role("status"),
						Span(Class("visually-hidden"), g.Text("Loading...")),
					),
				),
			),
		),
	)
}

// PositionInput renders a single position input box
func PositionInput(name, value, label string) g.Node {
	return Div(
		Input(
			Type("text"),
			Class("form-control position-input"),
			Name(name),
			Value(value),
			g.Attr("maxlength", "10"),
			Placeholder("."),
			g.Attr("autocomplete", "off"),
		),
		Div(Class("position-label"), g.Text(label)),
	)
}
