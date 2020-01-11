package trie

import (
	"encoding/json"
	"fmt"
	"unicode/utf8"
)

type Tree struct {
	Prefix   string
	Children []*Tree
}

func (t Tree) MarshalJSON() ([]byte, error) {
	if len(t.Children) == 0 {
		return []byte(fmt.Sprintf("[%q]", t.Prefix)), nil
	}
	c, err := json.Marshal(t.Children)
	if err != nil {
		return nil, err
	}
	return []byte(fmt.Sprintf("[%q,%s]", t.Prefix, c)), nil
}

func (t *Tree) end() {
	for _, child := range t.Children {
		if child.Prefix == "" {
			return
		}
	}
	t.Children = append(t.Children, &Tree{"", []*Tree{}})
}

func (t *Tree) Add(word string) {
	if len(word) == 0 {
		t.end()
		return
	}
	prefixRune, size := utf8.DecodeRuneInString(word)
	prefix := string(prefixRune)
	suffix := word[size:]
	for _, existingChild := range t.Children {
		if existingChild.Prefix == prefix {
			existingChild.Add(suffix)
			return
		}
	}
	newChild := &Tree{prefix, []*Tree{}}
	newChild.Add(suffix)
	t.Children = append(t.Children, newChild)
}
