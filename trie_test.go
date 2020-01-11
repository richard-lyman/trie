package trie

import (
	"github.com/richard-lyman/trie"
	"testing"
)

func testTree(t *testing.T) {
	tt := &trie.Tree{"", []*trie.Tree{}}
	words := []string{"a", "an", "and", "another", "always"}
	for _, word := range words {
		tt.Add(word)
	}
	//m, err := json.Marshal(tt)
}
