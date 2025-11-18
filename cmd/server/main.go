package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"wordle/dictionary"
	"wordle/handlers"
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

	// Use PORT from Heroku, fallback to WORDLE_PORT or 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = os.Getenv("WORDLE_PORT")
		if port == "" {
			port = "8080"
		}
	}

	// Use WORDLE_HOST or default to 0.0.0.0 for cloud deployment
	host := os.Getenv("WORDLE_HOST")
	if host == "" {
		host = "0.0.0.0"
	}

	// Initialize dictionary with reload capability
	_, _ = fmt.Fprintf(os.Stderr, "Loading dictionary from %s\n", dict)
	wordList, err := dictionary.NewWordList(os.Stderr, dict, remove)
	if err != nil {
		return fmt.Errorf("failed to load dictionary: %w", err)
	}

	// Set up logging
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	// Set up routes
	mux := http.NewServeMux()

	// Static files (must be registered before more specific routes in Go 1.22)
	fs := http.FileServer(http.Dir("web/static"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fs))

	// Main page
	mux.HandleFunc("GET /", handlers.HandleGetForm(logger))

	// Solve endpoint
	mux.HandleFunc("POST /wordle/solve", handlers.HandlePostSolve(logger, wordList))

	// Reload endpoint (for manual dictionary refresh)
	mux.HandleFunc("POST /reload", handlers.HandleReload(logger, wordList))

	// Start server
	addr := host + ":" + port
	logger.Info("Starting Wordle Helper server", "address", addr)
	_, _ = fmt.Fprintf(os.Stderr, "Server running at http://%s\n", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		return fmt.Errorf("server error: %w", err)
	}

	return nil
}
