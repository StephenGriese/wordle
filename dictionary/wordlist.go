package dictionary

import (
	"fmt"
	"io"
	"sync"
	"time"
)

// WordList manages the dictionary with reload capability
type WordList struct {
	words      []string
	lastReload time.Time
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

// Reload refreshes the word list from the dictionary and past words
func (wl *WordList) Reload() error {
	wl.mu.Lock()
	defer wl.mu.Unlock()

	_, _ = fmt.Fprintf(wl.stderr, "Reloading dictionary from %s\n", wl.dictPath)

	words, err := Create(wl.stderr, wl.dictPath, wl.removePath)
	if err != nil {
		return fmt.Errorf("failed to reload dictionary: %w", err)
	}

	wl.words = words
	wl.lastReload = time.Now()

	_, _ = fmt.Fprintf(wl.stderr, "Dictionary reloaded: %d words available (last reload: %s)\n",
		len(wl.words), wl.lastReload.Format(time.RFC3339))

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

// LastReload returns the time of the last reload
func (wl *WordList) LastReload() time.Time {
	wl.mu.RLock()
	defer wl.mu.RUnlock()
	return wl.lastReload
}

// Count returns the number of words in the dictionary
func (wl *WordList) Count() int {
	wl.mu.RLock()
	defer wl.mu.RUnlock()
	return len(wl.words)
}

// ShouldReload checks if reload is needed based on time elapsed
func (wl *WordList) ShouldReload(maxAge time.Duration) bool {
	wl.mu.RLock()
	defer wl.mu.RUnlock()
	return time.Since(wl.lastReload) > maxAge
}
