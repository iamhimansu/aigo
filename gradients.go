package main

import "fmt"

// THE CHAIN RULE:
// Final Weight Gradient = (Derivative of Loss) * (Derivative of Activation) * (The raw Input data)

// ---------------------------------------------------------
// YOUR JOB: Write the weightGradient function
// ---------------------------------------------------------

// HOW: Multiplies the backward error signals by the input to find the weight's specific gradient.
// WHY: This tells the AI exactly which direction and how forcefully to push the weight.
// PHP:
//
//	function weightGradient($errorSlope, $reluSlope, $input) {
//	    return $errorSlope * $reluSlope * $input;
//	}
func weightGradient(errorSlope float64, reluSlope float64, input float64) float64 {
	return errorSlope * reluSlope * input
}

func main() {
	// The MSE told us the error slope was 2.0
	errorSlope := 2.0

	// The ReLU was open, so its slope was 1.0
	reluSlope := 1.0

	// The raw data fed into this specific weight was 3.5
	input := 3.5

	// Calculate the exact gradient for this weight
	gradient := weightGradient(errorSlope, reluSlope, input)

	fmt.Printf("The gradient (blame) for this weight is: %f\n", gradient)
}
