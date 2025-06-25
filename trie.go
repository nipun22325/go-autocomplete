package main

import "fmt"

type TrieNode struct {
	links map[rune]*TrieNode
	isEnd bool
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{root: &TrieNode{links: make(map[rune]*TrieNode)}}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, ch := range word {
		if _, ok := node.links[ch]; !ok {
			node.links[ch] = &TrieNode{links: make(map[rune]*TrieNode)}
		}
		node = node.links[ch]
	}
	node.isEnd = true
}

func (t *Trie) Search(prefix string) []string {
	node := t.root
	for _, ch := range prefix {
		if _, ok := node.links[ch]; !ok {
			return []string{}
		}
		node = node.links[ch]
	}
	return collectWords(node, prefix)
}

func collectWords(node *TrieNode, prefix string) []string {
	results := []string{}
	if node.isEnd {
		results = append(results, prefix)
	}
	for ch, link := range node.links {
		results = append(results, collectWords(link, prefix+string(ch))...)
	}
	return results
}

func main() {
	root := NewTrie()
	root.Insert("app")
	root.Insert("apple")
	root.Insert("apt")
	root.Insert("bat")
	suggestedWords := root.Search("app")
	for _, word := range suggestedWords {
		fmt.Println(word);
	}
}