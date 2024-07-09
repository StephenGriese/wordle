package pastwords_test

import (
	"os"
	"testing"
	"wordle/pastwords"

	"golang.org/x/net/html"
)

func TestGetPastWords(t *testing.T) {
	reader, err := os.Open("testdata/example.txt")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	doc, err := html.Parse(reader)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	used := pastwords.GetPastWords(doc)
	if used == nil {
		t.Errorf("Error: %v", err)
	}
	if len(used) != 1104 {
		t.Errorf("Expected 1104, got %v", len(used))
	}
	if used[0] != "ABACK" {
		t.Errorf("Expected 'ABACK', got %v", used[0])
	}
	if used[1103] != "ZESTY" {
		t.Errorf("Expected 'ZYMURGY', got %v", used[1103])
	}
}

func TestFetchPastWords(t *testing.T) {
	used, err := pastwords.FetchPastWords()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if used == nil {
		t.Errorf("Error: %v", err)
	}
	if len(used) < 1109 {
		t.Errorf("Expected at least 1109 used words, got %v", len(used))
	}
	if used[0] != "aback" {
		t.Errorf("Expected first word to be 'ABACK', got %v", used[0])
	}
}
