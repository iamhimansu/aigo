package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// ---------------------------------------------------------
// YOUR JOB: Write the saveWeights function
// ---------------------------------------------------------

// HOW: Converts our float64 slice into a JSON string and saves it to a file.
// WHY: So the AI doesn't get amnesia when the program stops running.
// PHP:
//
//	function saveWeights($weights, $filename) {
//	    $jsonString = json_encode($weights);
//	    if ($jsonString === false) { die("Failed to encode JSON"); }
//
//	    $result = file_put_contents($filename, $jsonString);
//	    if ($result === false) { die("Failed to write file"); }
//	}
func saveWeights(weights []float64, filename string) {
	// 1. Convert 'weights' to JSON bytes using json.Marshal()
	// Remember to catch the 'err' and panic if it's not nil!

	// 2. Write those bytes to 'filename' using os.WriteFile()
	// The syntax for os.WriteFile is: os.WriteFile(filename, data, 0644)
	// (0644 is standard file permissions: read/write for owner, read for others).
	// Remember to catch the 'err' and panic if it's not nil!
	optimizedWeights, jerr := json.Marshal(weights)
	if jerr != nil {
		panic("INVALID WEIGHTS")
	}

	if werr := os.WriteFile(filename, optimizedWeights, 0644); werr != nil {
		panic("Could not save brain")
	}

}

func loadWeights(filename string) ([]float64, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		// Instead of crashing, we return the error to the caller
		return nil, err
	}

	var weights []float64
	if err := json.Unmarshal(bytes, &weights); err != nil {
		return nil, err
	}

	return weights, nil
}

func main() {
	filename := "brain.json"
	// These are the highly optimized weights your AI learned in the last lesson
	trainedWeights := []float64{0.5, 0.5}

	if weights, err := loadWeights(filename); err == nil {
		trainedWeights = weights
		fmt.Println("Loaded existing brain!")
	} else {
		fmt.Println("No brain found, starting fresh!")
	}
	// We will save them to a file named 'brain.json'
	saveWeights(trainedWeights, filename)
	fmt.Println("Brain successfully extracted and saved to brain.json")
}
