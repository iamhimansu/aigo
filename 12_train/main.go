package main // HOW: Standalone executable.

import "fmt" // WHY: To watch the AI learn in real-time.

// --- PHASE 2 & 3: FORWARD MATH ---
func dotProduct(a []float64, b []float64) float64 {
	sum := 0.00
	for i := 0; i < len(a); i++ {
		sum += a[i] * b[i]
	}
	return sum
}

func predict(inputs []float64, weights []float64, bias float64) float64 {
	return dotProduct(inputs, weights) + bias
}

// --- PHASE 5: BACKWARD MATH ---
func mseDerivative(prediction float64, target float64) float64 {
	return 2.0 * (prediction - target)
}

// PHP: function updateWeight($old, $grad, $lr) { return $old - ($grad * $lr); }
func updateWeight(oldWeight float64, gradient float64, learningRate float64) float64 {
	return oldWeight - (gradient * learningRate)
}

func main() {
	// THE REAL WORLD DATA
	// 1 Bed, 1 Bath. Target Price: $3.0 Million
	inputs := []float64{1.0, 1.0}
	target := 3.0

	// THE AI's GARBAGE STARTING BRAIN
	weights := []float64{0.5, 0.5}
	bias := 0.00 // We will keep bias locked at 0.0 for this pure weight training
	learningRate := 0.01
	epochs := 1_000

	fmt.Println("Beginning AI Training Sequence...")

	// ---------------------------------------------------------
	// YOUR FINAL JOB: Write the Epoch Loop
	// ---------------------------------------------------------
	// HOW: Use a standard Go 'for' loop to iterate from 0 up to 'epochs'.
	// PHP:
	// for ($epoch = 0; $epoch < $epochs; $epoch++) {
	//     $pred = predict($inputs, $weights, $bias);
	//     $errorSlope = mseDerivative($pred, $target);
	//     for ($i = 0; $i < count($weights); $i++) {
	//         $gradient = $errorSlope * 1.0 * $inputs[$i]; // reluSlope is 1.0 here
	//         $weights[$i] = updateWeight($weights[$i], $gradient, $learningRate);
	//     }
	// }

	for epoch := 0; epoch < epochs; epoch++ {
		pred := predict(inputs, weights, bias)
		errorSlope := mseDerivative(pred, target)
		if epoch%1000 == 0 {
			fmt.Printf("Trained epoch...%d\n", epoch)
		}
		for i := 0; i < len(weights); i++ {
			gradient := errorSlope * 1.0 * inputs[i]
			weights[i] = updateWeight(weights[i], gradient, learningRate)
		}
	}

	// Let's see what the AI predicts AFTER training
	finalPrediction := predict(inputs, weights, bias)
	fmt.Printf("\nTraining Complete. Final Weights: %v\n", weights)
	fmt.Printf("AI Final Prediction: $%.2f million\n", finalPrediction)
}
