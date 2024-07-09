package handlers

import (
	"log/slog"
	"net/http"
	"wordle/views"
)

func HandleGetForm(logger *slog.Logger, view *views.View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Getting Wordle form")
		err := view.Render(w, nil)
		if err != nil {
			logger.Error("Error rendering view", "error", err)
			http.Error(w, "Error", http.StatusInternalServerError)
		}
	}
}
