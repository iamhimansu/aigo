package main

import (
	"fmt"
	"strings"
)

func main() {

	text := "php is fast php is flexible golang is efficient"

	words := strings.Fields(text)

	vocab := make(map[string]int)

	var wordIDs []int

	idCounter := 0

	for _, word := range words {
		if _, exists := vocab[word]; !exists {
			vocab[word] = idCounter
			idCounter++
		}
		wordIDs = append(wordIDs, vocab[word])
	}

	var encoded []int
	for _, word := range words {
		id := vocab[word]

		encoded = append(encoded, id)
	}

	fmt.Println("Original Text:", text)
	fmt.Println("Vocabulary Map:", vocab)
	fmt.Println("Encoded IDs:", encoded)
}
