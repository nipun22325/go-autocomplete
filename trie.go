package main

import (
	"sort"
)

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
	suggestions := collectWords(node, prefix)
	return sortWordsByLength(suggestions);
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

func sortWordsByLength(words []string) []string{
	// Make a copy to avoid modifying the original slice
	sortedResults := make([]string, len(words))
	copy(sortedResults, words)
	// Sort by length (shorter words come first)
	sort.Slice(sortedResults, func(i, j int) bool {
		return len(sortedResults[i]) < len(sortedResults[j])
	})
	return sortedResults
}
