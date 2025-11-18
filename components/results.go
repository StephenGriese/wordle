package components

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

// Results renders the results section with found words
func Results(words []string, count int) g.Node {
	if count == 0 {
		return ResultsCard(
			H3(Class("mb-3"),
				g.Text("Possible Words "),
				Span(Class("badge bg-success"), g.Text("0 found")),
			),
			Div(Class("alert alert-warning"),
				H5(Class("alert-heading"), g.Text("🤔 No matching words found")),
				P(Class("mb-0"), g.Text("Please check your constraints - there might be a conflict in your clues.")),
			),
		)
	}

	var tip g.Node
	if count > 100 {
		tip = Div(Class("alert alert-info mt-3 mb-0"),
			Small(g.Textf("💡 Tip: With %d possible words, try entering more clues to narrow down the results.", count)),
		)
	}

	return ResultsCard(
		H3(Class("mb-3"),
			g.Text("Possible Words "),
			Span(Class("badge bg-success"), g.Textf("%d found", count)),
		),
		Div(Class("word-list"),
			g.Group(g.Map(words, func(word string) g.Node {
				return Span(Class("word-badge"), g.Text(word))
			})),
		),
		tip,
	)
}

// ResultsCard wraps results in a card
func ResultsCard(children ...g.Node) g.Node {
	return Div(Class("results-card"), g.Group(children))
}

// ReloadSuccessMessage renders a success message after dictionary reload
func ReloadSuccessMessage(wordCount int, timestamp string) g.Node {
	return Div(
		Class("alert alert-success alert-dismissible fade show mt-2"),
		Role("alert"),
		g.Textf("✅ Dictionary reloaded! %d words available (Updated: %s)", wordCount, timestamp),
		Button(
			Type("button"),
			Class("btn-close"),
			g.Attr("data-bs-dismiss", "alert"),
		),
	)
}

// ReloadErrorMessage renders an error message after failed dictionary reload
func ReloadErrorMessage(err error) g.Node {
	return Div(
		Class("alert alert-danger alert-dismissible fade show mt-2"),
		Role("alert"),
		Strong(g.Text("Error: ")),
		g.Textf("Failed to reload dictionary: %v", err),
		Button(
			Type("button"),
			Class("btn-close"),
			g.Attr("data-bs-dismiss", "alert"),
		),
	)
}
