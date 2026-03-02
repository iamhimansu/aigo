package main

import "fmt"

// HOW: Calculates dot product of two vectors.
// WHY: Core math engine.
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

// HOW: Linear prediction.
// WHY: Step 1 of a neuron.
// PHP: function predict($inputs, $weights, $bias) { ... }
func predict(inputs []float64, weights []float64, bias float64) float64 {
	return dotProduct(inputs, weights) + bias
}

// HOW: Non-linear activation.
// WHY: Allows the network to learn complex patterns.
// PHP: function relu($x) { return $x < 0 ? 0.0 : $x; }
func relu(x float64) float64 {
	if x < 0 {
		return 0.0
	}
	return x
}

// HOW: Full artificial brain cell.
// WHY: Combines math engine and activation.
// PHP: function neuron($inputs, $weights, $bias) { return relu(predict($inputs, $weights, $bias)); }
func neuron(inputs []float64, weights []float64, bias float64) float64 {
	return relu(predict(inputs, weights, bias))
}

func denseLayer(inputs []float64, weightMatrix [][]float64, biases []float64) []float64 {
	outputs := make([]float64, len(weightMatrix))
	for i := 0; i < len(weightMatrix); i++ {
		outputs[i] = neuron(inputs, weightMatrix[i], biases[i])
	}
	return outputs
}

func forwardPass(inputs []float64, weightsL1 [][]float64, biasesL1 []float64, weightsL2 [][]float64, biasesL2 []float64) []float64 {
	l1 := denseLayer(inputs, weightsL1, biasesL1)
	l2 := denseLayer(l1, weightsL2, biasesL2)
	return l2
}

func main() {
	// RAW DATA
	inputs := []float64{1.0, 2.0, 3.0}

	// LAYER 1: 3 Neurons
	weightsL1 := [][]float64{
		{0.2, 0.8, -0.5},
		{0.5, -0.91, 0.26},
		{-0.26, -0.27, 0.17},
	}
	biasesL1 := []float64{2.0, 3.0, 0.5}

	// LAYER 2: 2 Neurons (Notice the weights per neuron is 3, because Layer 1 outputs 3 numbers!)
	weightsL2 := [][]float64{
		{0.1, -0.14, 0.5},
		{-0.5, 0.12, -0.33},
	}
	biasesL2 := []float64{-1.0, 0.2}

	// FIRE THE NETWORK
	finalOutput := forwardPass(inputs, weightsL1, biasesL1, weightsL2, biasesL2)
	fmt.Printf("Deep Network Final Output: %v\n", finalOutput)
}
