package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	words := []string{"php", "is", "fast", "php", "is", "flexible", "golang", "is", "efficient"}

	brain := make(map[string][]string)

	for i := 0; i < len(words)-2; i++ {
		key := words[i] + " " + words[i+1]

		target := words[i+2]

		brain[key] = append(brain[key], target)
	}

	rand.Seed(time.Now().UnixNano())

	currentKey := "is efficient"

	for i := 0; i < 10; i++ {
		choices := brain[currentKey]

		if len(choices) == 0 {
			fmt.Print("[Dead End! Restarting...]")

			random := rand.Intn(len(words) - 1)

			nextWord := words[random] + " " + words[random+1]

			currentKey = nextWord
			fmt.Println(nextWord)
			continue
		}

		nextWord := choices[rand.Intn(len(choices))]

		fmt.Println(" " + nextWord)

		parts := strings.Split(currentKey, " ")

		newParts := append(parts[1:], nextWord)

		currentKey = strings.Join(newParts, " ")
	}
	fmt.Println()
}
