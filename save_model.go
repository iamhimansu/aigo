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
	werr := os.WriteFile("./brain.json", optimizedWeights, 0644)
	if werr != nil {
		panic("Could not save brain")
	}

}

func main() {
	// These are the highly optimized weights your AI learned in the last lesson
	trainedWeights := []float64{1.4831296806411502, 1.4831296806411502}

	// We will save them to a file named 'brain.json'
	saveWeights(trainedWeights, "brain.json")

	fmt.Println("Brain successfully extracted and saved to brain.json")
}
