package main

import "fmt"

func dotProduct(a []float64, b []float64) float64 {
	aLength := len(a)
	bLength := len(b)

	if aLength-bLength != 0 {
		panic("Dot product can only be calculated with same size vector")
	}

	sum := 0.00

	for i := 0; i < aLength; i++ {
		multiply := a[i] * b[i]
		sum += multiply
	}

	return sum
}

func predict(inputs []float64, weights []float64, bias float64) float64 {
	return dotProduct(inputs, weights) + bias
}

func main() {
	// Let's predict a house price (in hundreds of thousands)
	// Inputs: [Bedrooms, Bathrooms, Age of House]
	inputs := []float64{3.0, 2.0, 10.0}

	// Weights: [Value per bed, Value per bath, Value lost per year of age]
	weights := []float64{0.5, 0.25, -0.02}

	// Bias: Base land value
	bias := 1.5

	estimatedPrice := predict(inputs, weights, bias)
	fmt.Printf("The estimated house price is: $%.2f million\n", estimatedPrice)
}
