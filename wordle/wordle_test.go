package wordle_test

import (
	"fmt"
	"testing"
	"wordle/wordle"
)

type testCase struct {
	want         bool
	word         string
	missed       string
	lettersAt    []wordle.LetterAt
	lettersNotAt []wordle.LettersNotAt
}

func TestKeepName(t *testing.T) {

	foo := []testCase{
		{want: true, word: "least",
			lettersAt: []wordle.LetterAt{{Position: 2, Letter: 'a'}}},
		{want: true, word: "least",
			lettersNotAt: []wordle.LettersNotAt{
				{Position: 0, Letters: []byte{'a'}},
				{Position: 4, Letters: []byte{'a'}},
			},
		},
		{want: false, word: "least",
			lettersNotAt: []wordle.LettersNotAt{
				{Position: 2, Letters: []byte{'a'}},
			},
		},
		{want: false, word: "least",
			lettersAt: []wordle.LetterAt{
				{Position: 2, Letter: 'a'},
				{Position: 3, Letter: 'v'},
			},
		},
	}

	for _, v := range foo {
		actual := wordle.CheckWord(v.word, v.missed, v.lettersAt, v.lettersNotAt)
		if actual != v.want {
			fmt.Println("This test: ", v)
			t.Fatalf(`keepWord(%s, %s) = %v, want %v`, v.word, v.missed, actual, v.want)
		}
	}
}
