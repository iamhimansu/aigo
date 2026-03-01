package main

import "fmt"

// 1. THE MATH ENGINE
// PHP: function dotProduct($a, $b) { ... }
func dotProduct(a []float64, b []float64) float64 {
	if len(a) != len(b) {
		panic("Vectors must be the same length")
	}
	sum := 0.00
	for i := 0; i < len(a); i++ {
		sum += a[i] * b[i]
	}
	return sum
}

// 2. THE PREDICTOR (Linear)
// PHP: function predict($inputs, $weights, $bias) { ... }
func predict(inputs []float64, weights []float64, bias float64) float64 {
	return dotProduct(inputs, weights) + bias
}

// 3. THE ACTIVATION (Non-Linear)
// PHP: function relu($x) { return $x < 0 ? 0.0 : $x; }
func relu(x float64) float64 {
	if x < 0 {
		return 0.0
	}
	return x
}

// 4. THE NEURON
// HOW: This function must call predict() to get the raw linear score,
// and then immediately pass that score through relu() before returning it.
// PHP: function neuron($inputs, $weights, $bias) { ... }
func neuron(inputs []float64, weights []float64, bias float64) float64 {
	return relu(predict(inputs, weights, bias))
}

func main() {
	// A scenario that results in a NEGATIVE raw score
	inputs := []float64{1.0, 2.0}
	weights := []float64{-0.5, -0.8}
	bias := 0.5

	// If we just use predict(), the raw score is: (1.0 * -0.5) + (2.0 * -0.8) + 0.5 = -1.6
	// Let's see what the actual Neuron outputs!

	finalOutput := neuron(inputs, weights, bias)
	fmt.Printf("The Neuron's final activated output is: %f\n", finalOutput)
}
