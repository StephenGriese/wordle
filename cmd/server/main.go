package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"wordle/dictionary"
	"wordle/handlers"
	"wordle/views"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// Load configuration from environment
	dict := os.Getenv("WORDLE_DICTIONARY")
	if dict == "" {
		return fmt.Errorf("missing WORDLE_DICTIONARY environment variable")
	}

	remove := os.Getenv("WORDLE_REMOVE")
	if remove == "" {
		_, _ = fmt.Fprintf(os.Stderr, "No WORDLE_REMOVE environment variable. Will not remove any words.\n")
	}

	port := os.Getenv("WORDLE_PORT")
	if port == "" {
		port = "8080"
	}

	host := os.Getenv("WORDLE_HOST")
	if host == "" {
		host = "localhost"
	}

	// Initialize dictionary
	_, _ = fmt.Fprintf(os.Stderr, "Loading dictionary from %s\n", dict)
	words, err := dictionary.Create(os.Stderr, dict, remove)
	if err != nil {
		return fmt.Errorf("failed to load dictionary: %w", err)
	}
	_, _ = fmt.Fprintf(os.Stderr, "Loaded %d words\n", len(words))

	// Set up logging
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	// Create views
	formView := views.NewView("bootstrap", "web/views/wordleform.gohtml")
	resultsView := views.NewView("results", "web/views/results.gohtml")

	// Set up routes
	mux := http.NewServeMux()

	// Static files (must be registered before more specific routes in Go 1.22)
	fs := http.FileServer(http.Dir("web/static"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fs))

	// Main page
	mux.HandleFunc("GET /", handlers.HandleGetForm(logger, formView))

	// Solve endpoint
	mux.HandleFunc("POST /wordle/solve", handlers.HandlePostSolve(logger, formView, resultsView, words))

	// Start server
	addr := host + ":" + port
	logger.Info("Starting Wordle Helper server", "address", addr)
	_, _ = fmt.Fprintf(os.Stderr, "Server running at http://%s\n", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		return fmt.Errorf("server error: %w", err)
	}

	return nil
}
