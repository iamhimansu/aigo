package main

import "fmt"

func main() {
	reverseVocab := map[int]string{
		0: "php",
		1: "is",
		2: "fast",
	}

	outputIDs := []int{2, 1, 0}

	outputString := ""

	for _, encodedValue := range outputIDs {

		if _, exists := reverseVocab[encodedValue]; !exists {
			continue
		}

		outputString = outputString + reverseVocab[encodedValue] + " "
	}

	fmt.Println(outputString)
}
