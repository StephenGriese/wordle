package handlers

import (
	"log/slog"
	"net/http"
	"strings"
	"wordle/usrcmd"
	"wordle/views"
	"wordle/wordle"
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

// PageData represents the data passed to templates
type PageData struct {
	FormData  FormData
	Results   []string
	WordCount int
	Error     string
}

// HandleGetForm renders the initial empty form
func HandleGetForm(logger *slog.Logger, view *views.View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Getting Wordle form")

		data := PageData{
			FormData: FormData{
				Missed: "",
				Pos0:   ".",
				Pos1:   ".",
				Pos2:   ".",
				Pos3:   ".",
				Pos4:   ".",
			},
		}

		err := view.Render(w, data)
		if err != nil {
			logger.Error("Error rendering view", "error", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}

// HandlePostSolve processes the form submission and returns filtered words
func HandlePostSolve(logger *slog.Logger, view *views.View, resultsView *views.View, words []string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Solving Wordle")

		// Parse form data
		if err := r.ParseForm(); err != nil {
			logger.Error("Error parsing form", "error", err)
			renderError(w, view, logger, "Invalid form data", FormData{})
			return
		}

		formData := FormData{
			Missed: strings.TrimSpace(r.FormValue("missed")),
			Pos0:   strings.TrimSpace(r.FormValue("pos0")),
			Pos1:   strings.TrimSpace(r.FormValue("pos1")),
			Pos2:   strings.TrimSpace(r.FormValue("pos2")),
			Pos3:   strings.TrimSpace(r.FormValue("pos3")),
			Pos4:   strings.TrimSpace(r.FormValue("pos4")),
		}

		// Normalize empty positions to dots
		normalizePosition(&formData.Pos0)
		normalizePosition(&formData.Pos1)
		normalizePosition(&formData.Pos2)
		normalizePosition(&formData.Pos3)
		normalizePosition(&formData.Pos4)

		// Convert form data to wordle types
		missed, lettersAt, lettersNotAt, err := parseFormToWordleInputs(formData)
		if err != nil {
			logger.Error("Error parsing wordle inputs", "error", err)
			renderError(w, view, logger, "Invalid input format: "+err.Error(), formData)
			return
		}

		// Find possible words
		possibles := wordle.MakePossibles(words, missed, lettersAt, lettersNotAt)

		logger.Info("Found possible words", "count", len(possibles))

		// Render results
		data := PageData{
			FormData:  formData,
			Results:   possibles,
			WordCount: len(possibles),
		}

		// Check if this is an HTMX request - if so, render only the results partial
		isHTMX := r.Header.Get("HX-Request") == "true"

		var renderView *views.View
		if isHTMX {
			renderView = resultsView // Render just the results partial
		} else {
			renderView = view // Render full page (for non-HTMX fallback)
		}

		err = renderView.Render(w, data)
		if err != nil {
			logger.Error("Error rendering view", "error", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}

// normalizePosition converts empty strings to dots
func normalizePosition(pos *string) {
	if *pos == "" {
		*pos = "."
	}
}

// parseFormToWordleInputs converts form data to wordle types using existing usrcmd logic
func parseFormToWordleInputs(formData FormData) (string, []wordle.LetterAt, []wordle.LettersNotAt, error) {
	// Build command line format that usrcmd expects
	positions := []string{
		formData.Pos0,
		formData.Pos1,
		formData.Pos2,
		formData.Pos3,
		formData.Pos4,
	}

	// Create space-separated string like CLI input
	cmdLine := formData.Missed + " " + strings.Join(positions, " ")

	// Use existing command line parser
	return usrcmd.ReadUserCommand(cmdLine)
}

// renderError renders the form with an error message
func renderError(w http.ResponseWriter, view *views.View, logger *slog.Logger, errMsg string, formData FormData) {
	data := PageData{
		FormData: formData,
		Error:    errMsg,
	}

	err := view.Render(w, data)
	if err != nil {
		logger.Error("Error rendering error view", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
