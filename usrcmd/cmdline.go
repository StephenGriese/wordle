package usrcmd

import (
	"fmt"
	"strings"
	"wordle/wordle"
)

func ReadUserCommand(s string) (string, []wordle.LetterAt, []wordle.LettersNotAt, error) {
	s = strings.TrimSpace(s)
	parts := strings.Split(s, " ")
	return readArgs(parts)
}

func readArgs(args []string) (string, []wordle.LetterAt, []wordle.LettersNotAt, error) {
	if len(args) == 0 {
		return "", nil, nil, fmt.Errorf("no arguments")
	}
	cmdlineArgs := args
	if len(cmdlineArgs) < 6 {
		return "", nil, nil, fmt.Errorf("not enough arguments")
	}
	missed := cmdlineArgs[0]
	var lettersAt []wordle.LetterAt
	var lettersNotAt []wordle.LettersNotAt
	for i, v := range cmdlineArgs[1:] {
		if v == "." {
			continue
		}
		if strings.HasPrefix(v, "-") {
			lettersNotAt = append(lettersNotAt, wordle.LettersNotAt{Position: i, Letters: []byte(v[1:])})
		} else {
			lettersAt = append(lettersAt, wordle.LetterAt{Position: i, Letter: v[0]})
		}
	}
	return missed, lettersAt, lettersNotAt, nil
}
