package wordle

import (
	"bytes"
)

type LettersNotAt struct {
	Position int // 0-based
	Letters  []byte
}

type LetterAt struct {
	Position int // 0-based
	Letter   byte
}

func MakePossibles(words []string, missed string, lettersAt []LetterAt, lettersNotAt []LettersNotAt) []string {
	var possibles []string
	for _, word := range words {
		if CheckWord(word, missed, lettersAt, lettersNotAt) {
			possibles = append(possibles, word)
		}
	}
	return possibles
}

// CheckWord will take use missed, lettersAt, and lettersNotAt to determine if the word is a possible word.
func CheckWord(word string, missed string, lettersAt []LetterAt, lettersNotAt []LettersNotAt) bool {
	if wordContainsMissed(word, missed) {
		return false
	}

	for _, c := range lettersAt {
		if !positionContainsLetter(word, c.Position, c.Letter) {
			return false
		}
	}

	for _, c := range lettersNotAt {
		for _, letter := range c.Letters {
			if !wordContainsMissed(word, string(letter)) {
				return false
			}
			if positionContainsLetter(word, c.Position, letter) {
				return false
			}
		}
	}
	return true
}

func wordContainsMissed(word, missed string) bool {
	return bytes.ContainsAny([]byte(word), missed)
}

func positionContainsLetter(word string, position int, letter byte) bool {
	b := []byte(word)
	return b[position] == letter
}
