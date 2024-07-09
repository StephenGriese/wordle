package wordle

import "testing"

func TestWordContainsMissed(t *testing.T) {
	type TestCase struct {
		Word   string
		Missed string
		Wanted bool
	}
	tests := []TestCase{
		{Word: "hello", Missed: "h", Wanted: true},
		{Word: "hello", Missed: "he", Wanted: true},
		{Word: "hello", Missed: "b", Wanted: false},
	}
	for _, v := range tests {
		if got := wordContainsMissed(v.Word, v.Missed); got != v.Wanted {
			t.Fatalf(`wanted %v, got %v`, v.Wanted, got)
		}
	}
}

func TestPositionContainsLetter(t *testing.T) {
	type testCase struct {
		position int // 0-based
		letter   byte
		word     string
		wanted   bool
	}
	tests := []testCase{
		{position: 0, letter: byte('h'), word: "hello", wanted: true},
		{position: 0, letter: byte('b'), word: "hello", wanted: false},
	}
	for _, v := range tests {
		if got := positionContainsLetter(v.word, v.position, v.letter); got != v.wanted {
			t.Fatalf(`wanted %v, got %v`, v.wanted, got)
		}
	}
}
