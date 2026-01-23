package components

import (
	g "github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// hasRepeatedLetters checks if a word has any repeated letters
func hasRepeatedLetters(word string) bool {
	seen := make(map[rune]bool)
	for _, char := range word {
		if seen[char] {
			return true
		}
		seen[char] = true
	}
	return false
}

// separateWordsByRepetition divides words into two slices:
// words without repeated letters and words with repeated letters
func separateWordsByRepetition(words []string) (noRepeats []string, withRepeats []string) {
	for _, word := range words {
		if hasRepeatedLetters(word) {
			withRepeats = append(withRepeats, word)
		} else {
			noRepeats = append(noRepeats, word)
		}
	}
	return noRepeats, withRepeats
}

// Results renders the results section with found words
func Results(words []string, count int) g.Node {
	if count == 0 {
		return ResultsCard(
			html.H3(html.Class("mb-3"),
				g.Text("Possible Words "),
				html.Span(html.Class("badge bg-success"), g.Text("0 found")),
			),
			html.Div(html.Class("alert alert-warning"),
				html.H5(html.Class("alert-heading"), g.Text("🤔 No matching words found")),
				html.P(html.Class("mb-0"), g.Text("Please check your constraints - there might be a conflict in your clues.")),
			),
		)
	}

	// Separate words into two groups
	noRepeats, withRepeats := separateWordsByRepetition(words)

	var tip g.Node
	if count > 100 {
		tip = html.Div(html.Class("alert alert-info mt-3 mb-0"),
			html.Small(g.Textf("💡 Tip: With %d possible words, try entering more clues to narrow down the results.", count)),
		)
	}

	// Build sections for each group
	var sections []g.Node

	// Section for words without repeated letters
	if len(noRepeats) > 0 {
		sections = append(sections,
			html.Div(html.Class("word-section mb-4"),
				html.H5(html.Class("mb-2"),
					g.Text("No Repeated Letters "),
					html.Span(html.Class("badge bg-primary"), g.Textf("%d", len(noRepeats))),
				),
				html.Div(html.Class("word-list"),
					g.Group(g.Map(noRepeats, func(word string) g.Node {
						return html.Span(html.Class("word-badge"), g.Text(word))
					})),
				),
			),
		)
	}

	// Section for words with repeated letters
	if len(withRepeats) > 0 {
		sections = append(sections,
			html.Div(html.Class("word-section"),
				html.H5(html.Class("mb-2"),
					g.Text("Repeated Letters "),
					html.Span(html.Class("badge bg-secondary"), g.Textf("%d", len(withRepeats))),
				),
				html.Div(html.Class("word-list"),
					g.Group(g.Map(withRepeats, func(word string) g.Node {
						return html.Span(html.Class("word-badge"), g.Text(word))
					})),
				),
			),
		)
	}

	return ResultsCard(
		html.H3(html.Class("mb-3"),
			g.Text("Possible Words "),
			html.Span(html.Class("badge bg-success"), g.Textf("%d found", count)),
		),
		g.Group(sections),
		tip,
	)
}

// ResultsCard wraps results in a card
func ResultsCard(children ...g.Node) g.Node {
	return html.Div(html.Class("results-card"), g.Group(children))
}

// ReloadSuccessMessage renders a success message after dictionary reload
func ReloadSuccessMessage(wordCount int, timestamp string) g.Node {
	return html.Div(
		html.Class("alert alert-success alert-dismissible fade show mt-2"),
		html.Role("alert"),
		g.Textf("✅ Dictionary reloaded! %d words available (Updated: %s)", wordCount, timestamp),
		html.Button(
			html.Type("button"),
			html.Class("btn-close"),
			g.Attr("data-bs-dismiss", "alert"),
		),
	)
}

// ReloadErrorMessage renders an error message after failed dictionary reload
func ReloadErrorMessage(err error) g.Node {
	return html.Div(
		html.Class("alert alert-danger alert-dismissible fade show mt-2"),
		html.Role("alert"),
		html.Strong(g.Text("Error: ")),
		g.Textf("Failed to reload dictionary: %v", err),
		html.Button(
			html.Type("button"),
			html.Class("btn-close"),
			g.Attr("data-bs-dismiss", "alert"),
		),
	)
}
