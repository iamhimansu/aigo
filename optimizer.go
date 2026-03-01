package main

import "fmt"

// ---------------------------------------------------------
// YOUR JOB: Write the updateWeight function
// ---------------------------------------------------------

// HOW: Adjusts a single weight by subtracting a fraction of its gradient.
// WHY: This is the exact moment the AI actually "learns" and gets smarter.
// PHP:
//
//	function updateWeight($oldWeight, $gradient, $learningRate) {
//	    return $oldWeight - ($learningRate * $gradient);
//	}
func updateWeight(oldWeight float64, gradient float64, learningRate float64) float64 {
	return oldWeight - (gradient * learningRate)
}

func main() {
	// The original weight before the AI made a mistake
	oldWeight := 0.5

	// The Chain Rule told us this weight's blame was 7.0
	gradient := 7.0

	// The leash: We only let the AI apply 1% of the blame per step
	learningRate := 0.01

	// Calculate the new, smarter weight
	newWeight := updateWeight(oldWeight, gradient, learningRate)

	fmt.Printf("The original weight was: %f\n", oldWeight)
	fmt.Printf("The newly optimized weight is: %f\n", newWeight)
}
