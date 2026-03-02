package main

import "fmt"

func relu(x float64) float64 {
	if x < 0 {
		return 0.0
	}
	return x
}

func main() {

	input := 0.23

	fmt.Println(relu(input))

}
