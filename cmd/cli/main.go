package main

import (
	"fmt"
	"io"
	"os"
	"wordle/dictionary"
	"wordle/scan"
	"wordle/usrcmd"
	"wordle/wordle"
)

func main() {
	if err := Run(os.Getenv, os.Stdin, os.Stdout, os.Stderr); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}

func Run(
	getenv func(string) string,
	stdin io.Reader,
	stdout, stderr io.Writer,
) error {
	dict := getenv("WORDLE_DICTIONARY")
	if dict == "" {
		return fmt.Errorf("missing WORDLE_DICTIONARY environment variable")
	}

	remove := getenv("WORDLE_REMOVE")
	if remove == "" {
		_, _ = fmt.Fprintf(stderr, "No WORDLE_REMOVE environment variable. Will not remove any words.\n")
	}

	words, err := dictionary.Create(stderr, dict, remove)
	if err != nil {
		return err
	}
	_, _ = fmt.Fprintf(stderr, "Loaded %d words\n", len(words))

	return readUserInput(stdout, stderr, stdin, words)
}

func readUserInput(stdout, stderr io.Writer, r io.Reader, words []string) error {
	return scan.Scan(r, createLineHandler(stdout, stderr, words))
}

func createLineHandler(stdout, stderr io.Writer, words []string) func(s string) error {
	return func(s string) error {
		missed, lettersAt, lettersNotAt, err := usrcmd.ReadUserCommand(s)
		if err != nil {
			_, _ = fmt.Fprintf(stderr, "error: %s\n", err)
			return nil
		}
		possibles := wordle.MakePossibles(words, missed, lettersAt, lettersNotAt)
		printPossibles(stdout, possibles)
		return nil
	}
}

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

func printPossibles(stdout io.Writer, possibles []string) {
	const cols = 15

	// Separate words into two groups
	noRepeats, withRepeats := separateWordsByRepetition(possibles)

	// Print words without repeated letters
	if len(noRepeats) > 0 {
		_, _ = fmt.Fprintf(stdout, "No Repeated Letters (%d):\n", len(noRepeats))
		c := 0
		for _, word := range noRepeats {
			_, _ = fmt.Fprintf(stdout, "%s ", word)
			c++
			if c == cols {
				_, _ = fmt.Fprintf(stdout, "\n")
				c = 0
			}
		}
		if c > 0 {
			_, _ = fmt.Fprintf(stdout, "\n")
		}
		_, _ = fmt.Fprintf(stdout, "\n")
	}

	// Print words with repeated letters
	if len(withRepeats) > 0 {
		_, _ = fmt.Fprintf(stdout, "Repeated Letters (%d):\n", len(withRepeats))
		c := 0
		for _, word := range withRepeats {
			_, _ = fmt.Fprintf(stdout, "%s ", word)
			c++
			if c == cols {
				_, _ = fmt.Fprintf(stdout, "\n")
				c = 0
			}
		}
		if c > 0 {
			_, _ = fmt.Fprintf(stdout, "\n")
		}
		_, _ = fmt.Fprintf(stdout, "\n")
	}
}
