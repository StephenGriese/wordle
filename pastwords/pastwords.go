package pastwords

import (
	"net/http"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func HasAttribe(n *html.Node, key, value string) bool {
	for _, a := range n.Attr {
		if a.Key == key && a.Val == value {
			return true
		}
	}
	return false
}

func MaybeKeep(node *html.Node) (keep bool) {
	if node.Type != html.TextNode {
		return false
	}
	node = node.Parent
	if node.DataAtom != atom.Li {
		return false
	}
	node = node.Parent
	if node.DataAtom != atom.Ul {
		return false
	}
	if !HasAttribe(node, "class", "inline") {
		return false
	}
	return true
}

func GetPastWords(doc *html.Node) []string {
	var used []string
	matcher := func(node *html.Node) (keep bool, exit bool) {
		if MaybeKeep(node) {
			keep = true
		}
		return
	}
	nodes := TraverseNode(doc, matcher)
	for _, node := range nodes {
		used = append(used, node.Data)
	}
	return used
}

func TraverseNode(doc *html.Node, matcher func(node *html.Node) (bool, bool)) (nodes []*html.Node) {
	var keep, exit bool
	var f func(*html.Node)
	f = func(n *html.Node) {
		keep, exit = matcher(n)
		if keep {
			nodes = append(nodes, n)
		}
		if exit {
			return
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return nodes
}

func FetchPastWords() ([]string, error) {
	requestURL := "https://www.rockpapershotgun.com/wordle-past-answers"
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	doc, err := html.Parse(res.Body)
	if err != nil {
		return nil, err
	}
	var used []string
	matcher := func(node *html.Node) (keep bool, exit bool) {
		if MaybeKeep(node) {
			keep = true
		}
		return
	}
	nodes := TraverseNode(doc, matcher)
	for _, node := range nodes {
		// The text is in all-caps, so we convert it to lowercase
		used = append(used, strings.ToLower(node.Data))
	}
	return used, nil
}
