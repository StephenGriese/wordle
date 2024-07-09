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
	dict := getenv("DICTIONARY")
	if dict == "" {
		return fmt.Errorf("missing DICTIONARY environment variable")
	}

	remove := getenv("REMOVE")
	if remove == "" {
		_, _ = fmt.Fprintf(stderr, "No REMOVE environment variable. Will not remove any words.\n")
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

func printPossibles(stdout io.Writer, possibles []string) {
	const cols = 15
	c := 0
	for _, word := range possibles {
		_, _ = fmt.Fprintf(stdout, "%s ", word)
		c++
		if c == cols {
			_, _ = fmt.Fprintf(stdout, "\n")
			c = 0
		}
	}
	_, _ = fmt.Fprintf(stdout, "\n\n")
}
