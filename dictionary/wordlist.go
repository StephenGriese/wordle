package dictionary

import (
	"fmt"
	"io"
	"sync"
)

// WordList manages the in-memory dictionary words.
type WordList struct {
	words      []string
	dictPath   string
	removePath string
	stderr     io.Writer
	mu         sync.RWMutex
}

// NewWordList creates a new managed word list
func NewWordList(stderr io.Writer, dictPath, removePath string) (*WordList, error) {
	wl := &WordList{
		dictPath:   dictPath,
		removePath: removePath,
		stderr:     stderr,
	}

	if err := wl.Reload(); err != nil {
		return nil, err
	}

	return wl, nil
}

// Reload refreshes the word list from the configured files.
func (wl *WordList) Reload() error {
	wl.mu.Lock()
	defer wl.mu.Unlock()

	_, _ = fmt.Fprintf(wl.stderr, "Reloading dictionary from %s\n", wl.dictPath)

	words, err := Create(wl.stderr, wl.dictPath, wl.removePath)
	if err != nil {
		return fmt.Errorf("failed to reload dictionary: %w", err)
	}

	wl.words = words

	_, _ = fmt.Fprintf(wl.stderr, "Dictionary reloaded: %d words available\n", len(wl.words))

	return nil
}

// Words returns a copy of the current word list (thread-safe)
func (wl *WordList) Words() []string {
	wl.mu.RLock()
	defer wl.mu.RUnlock()

	// Return a copy to prevent external modification
	result := make([]string, len(wl.words))
	copy(result, wl.words)
	return result
}

