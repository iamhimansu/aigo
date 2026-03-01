package main

import "fmt"

// YOUR JOB: Write the dotProduct function here

func main() {
	vectorA := []float64{0.1, 0.5, -0.2}
	vectorB := []float64{0.4, -0.1, 0.8}

	result := dotProduct(vectorA, vectorB)
	fmt.Printf("The dot product is: %f\n", result)
}

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
