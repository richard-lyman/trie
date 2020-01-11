package trie

import (
	"github.com/richard-lyman/trie"
	"testing"
)

func testTree(t *testing.T) {
	t := &trie.Tree{"", []*trie.Tree{}}
	words := []string{"a", "an", "and", "another", "always"}
	for _, word := range words {
		t.Add(word)
	}
	//m, err := json.Marshal(t)
}
