package main // HOW: Standalone executable.

import "fmt" // WHY: For printing output.

// HOW: Calculates the derivative (slope) of the MSE for a single prediction.
// WHY: This is the very first step of Backpropagation. It tells us the direction and size of our final error.
// PHP:
//
//	function mseDerivative($prediction, $target) {
//	    return 2.0 * ($prediction - $target);
//	}
func mseDerivative(prediction float64, target float64) float64 {
	return 2.0 * (prediction - target)
}

func main() {
	// The AI guessed a house price of $4.0 million
	prediction := 4.0

	// The actual price was $3.0 million (The AI guessed TOO HIGH)
	target := 3.0

	// Let's calculate the slope of this mistake
	slope := mseDerivative(prediction, target)

	// If the slope is POSITIVE, it tells the AI to push its weights DOWN next time.
	fmt.Printf("The Error Slope is: %f\n", slope)
}
