# Trie

[![GoDoc](https://godoc.org/github.com/richard-lyman/trie?status.svg)](https://godoc.org/github.com/richard-lyman/trie)

A [Trie](https://en.wikipedia.org/wiki/Trie) is used for re*trie*val of values based on keys.
This library provides an exported struct that uses the same string for key and value.
Each rune in a string that is added to the trie.Tree is a node in the trie and a string is identified by beginning and ending empty sentinels.

# Use this library
Assuming you are using go modules, here is a partial example of using this library:

```
import (
	...
	"github.com/richard-lyman/trie"
	...
)

...

var wordTree *trie.Tree

func createWords(words []string){
	for _, word := range words {
		wordTree.Add(word)
	}
}

...

func provideWords(w http.ResponseWriter, r *http.Request){
	jsonWordTree, err := json.Marshal(wordTree)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, string(jsonWordTree))
}
```
