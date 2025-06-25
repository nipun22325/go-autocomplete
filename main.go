package main

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
)

var root = NewTrie()


func loadWordsFromFile(filePath string, t *Trie){
	file, err := os.Open(filePath)
	if err!= nil {
		log.Fatalf("Failed to open word list: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		word:= strings.ToLower(scanner.Text())
		if len(word) > 0{
			t.Insert(word)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
}

func autocompleteHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	suggestions := root.Search(query)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(suggestions)
}

func main() {
	loadWordsFromFile("english-words.txt", root);
	http.HandleFunc("/autocomplete", autocompleteHandler)
	http.Handle("/", http.FileServer(http.Dir("./static"))) // Serve HTML
	println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}