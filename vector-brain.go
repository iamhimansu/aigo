package main

import "fmt"

func main() {

	words := []string{"php", "is", "fast", "php", "is", "flexible", "golang", "is", "efficient"}

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

	word := "is"
	id := vocab[word]
	vector := make([]int, len(vocab))

	vector[id] = 1

	fmt.Println("Vocabulary:", vocab)
	fmt.Println("Sentence as IDs:", wordIDs)
	fmt.Printf("Word: %s | ID: %d | Vector: %v\n", word, id, vector)

	weights :=
		[][]float64{
			{0.1, 0.5},  //php
			{0.9, -0.2}, //is
			{0.3, 0.8},  //fast
		}

	embedding := []float64{0, 0}

	for i := 0; i < len(vector); i++ {
		if vector[i] == 1 {
			embedding = weights[i]
		}
	}

	fmt.Printf("Input Vector: %v\n", vector)
	fmt.Printf("Hidden 'Thought' (Embedding): %v\n", embedding)

}
