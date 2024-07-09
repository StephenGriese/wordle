package usrcmd_test

import (
	"reflect"
	"testing"
	"wordle/usrcmd"
	"wordle/wordle"
)

func TestParams(t *testing.T) {
	tests := map[string]struct {
		args         string
		wantErr      bool
		missed       string
		lettersAt    []wordle.LetterAt
		lettersNotAt []wordle.LettersNotAt
	}{
		"no args": {
			args:    "",
			wantErr: true,
		},
		"too few args": {
			args:    "",
			wantErr: true,
		},
		"get missed": {
			args:         "abcd . . . . .",
			wantErr:      false,
			missed:       "abcd",
			lettersAt:    nil,
			lettersNotAt: nil,
		},
		"get missed and letters": {
			args:    "abc -ab . . . .",
			wantErr: false,
			missed:  "abc",
			lettersNotAt: []wordle.LettersNotAt{
				{Position: 0, Letters: []byte{'a', 'b'}},
			},
			lettersAt: nil,
		},
		"get missed and both at and non-at": {
			args:    "abc -ab c . . .",
			wantErr: false,
			missed:  "abc",
			lettersNotAt: []wordle.LettersNotAt{
				{Position: 0, Letters: []byte{'a', 'b'}},
			},
			lettersAt: []wordle.LetterAt{
				{Position: 1, Letter: 'c'},
			},
		},
		"an actual test case": {
			args:    "ertios -ag . -a -n .",
			wantErr: false,
			missed:  "ertios",
			lettersNotAt: []wordle.LettersNotAt{
				{Position: 0, Letters: []byte{'a', 'g'}},
				{Position: 2, Letters: []byte{'a'}},
				{Position: 3, Letters: []byte{'n'}},
			},
			lettersAt: nil,
		},
	}
	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			missed, lettersAt, lettersNotAt, err := usrcmd.ReadUserCommand(test.args)
			if (err != nil) != test.wantErr {
				t.Errorf("ReadArgs() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if missed != test.missed {
				t.Errorf("ReadArgs() missed = %v, want %v", missed, test.missed)
			}
			if !reflect.DeepEqual(lettersAt, test.lettersAt) {
				t.Errorf("ReadArgs() lettersAt = %v, want %v", lettersAt, test.lettersAt)
			}
			if !reflect.DeepEqual(lettersNotAt, test.lettersNotAt) {
				t.Errorf("ReadArgs() lettersNotAt = %v, want %v", lettersNotAt, test.lettersNotAt)
			}
		})
	}

}
