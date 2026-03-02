package main

import "fmt"

// THE ORIGINAL (Forward Pass)
// PHP: function relu($x) { return $x < 0 ? 0.0 : $x; }
func relu(x float64) float64 {
	if x < 0 {
		return 0.0
	}
	return x
}

// ---------------------------------------------------------
// YOUR JOB: Write the reluDerivative function
// ---------------------------------------------------------

// HOW: Calculates the slope of the ReLU function for a given input.
// WHY: Required for the AI to push errors backward through the network.
// PHP:
//
//	function reluDerivative($x) {
//	    if ($x <= 0) {
//	        return 0.0;
//	    }
//	    return 1.0;
//	}
func reluDerivative(x float64) float64 {
	if x <= 0 {
		return 0.0
	}
	return 1.0
}

func main() {
	// Let's test the slope at different points
	fmt.Println("Slope at -5.0:", reluDerivative(-5.0)) // Should be 0
	fmt.Println("Slope at 10.2:", reluDerivative(10.2)) // Should be 1
}
