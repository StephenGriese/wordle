package dictionary

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
	"unicode"
	"wordle/pastwords"
	"wordle/scan"
)

func Create(stderr io.Writer, dict, remove string) ([]string, error) {
	loaded, err := loadDictionary(stderr, dict)
	if err != nil {
		return nil, err
	}

	used, err := pastwords.FetchPastWords()
	if err != nil {
		return nil, err
	}
	_, _ = fmt.Fprintf(stderr, "Loaded %d past words\n", len(used))

	for _, word := range used {
		delete(loaded, word)
	}
	_, _ = fmt.Fprintf(stderr, "Loaded has %d words after removing past words\n", len(loaded))

	// Remove words from the remove list
	removeList, err := loadWords(stderr, remove)
	if err != nil {
		return nil, err
	}

	for _, word := range removeList {
		delete(loaded, word)
	}
	_, _ = fmt.Fprintf(stderr, "Loaded has %d words after removing the remove list\n", len(loaded))

	// Convert to a sorted slice
	var words []string
	for k := range loaded {
		words = append(words, k)
	}
	slices.Sort(words)

	return words, nil
}

func loadWords(stderr io.Writer, path string) ([]string, error) {
	if path == "" {
		_, _ = fmt.Fprintf(stderr, "No remove list\n")
		return nil, nil
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = file.Close()
	}()

	var words []string
	scanFunc := func(s string) error {
		words = append(words, strings.TrimSpace(s))
		return nil
	}
	err = scan.Scan(file, scanFunc)
	if err != nil {
		return nil, err
	}

	_, _ = fmt.Fprintf(stderr, "Loaded %d words from %s\n", len(words), path)

	return words, nil
}

func loadDictionary(stderr io.Writer, path string) (map[string]bool, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = file.Close()
	}()

	set := make(map[string]bool)

	scanFunc := func(s string) error {
		if keepWord(s) {
			set[s] = true
		}
		return nil
	}
	err = scan.Scan(file, scanFunc)
	if err != nil {
		return nil, err
	}

	_, _ = fmt.Fprintf(stderr, "Loaded %d words from %s\n", len(set), path)

	return set, nil
}

// KeepWord will use built-in logic to determine if a word should be added to our working dictionary.
func keepWord(word string) bool {
	if !isASCII(word) {
		return false
	}
	if !isLetter(word) {
		return false
	}
	if len(word) != 5 {
		return false
	}
	if startsWithCapitalAscii(word) {
		return false
	}

	return true
}

func isASCII(s string) bool {
	for _, r := range s {
		if r > unicode.MaxASCII {
			return false
		}
	}
	return true
}

func isLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func startsWithCapitalAscii(s string) bool {
	if len(s) == 0 {
		return false
	}
	return s[0] >= 'A' && s[0] <= 'Z'
}
