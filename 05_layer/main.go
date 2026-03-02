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

func main() {
	// 1 set of inputs (The evidence)
	inputs := []float64{1.0, 2.0, 3.0}

	// 3 sets of weights (3 Judges, each with 3 weights for the 3 inputs)
	weightMatrix := [][]float64{
		{0.2, 0.8, -0.5},     // Weights for Neuron 0
		{0.5, -0.91, 0.26},   // Weights for Neuron 1
		{-0.26, -0.27, 0.17}, // Weights for Neuron 2
	}

	// 3 biases (1 for each Judge)
	biases := []float64{2.0, 3.0, 0.5}

	// Execute the layer
	layerOutputs := denseLayer(inputs, weightMatrix, biases)
	fmt.Printf("Layer Outputs: %v\n", layerOutputs)
}
