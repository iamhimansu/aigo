package main // HOW: Standalone executable program.

import "fmt" // WHY: For printing the final error to the console.

// HOW: Calculates the Mean Squared Error between the AI's guesses and the true answers.
// WHY: This is the ONLY way the AI knows if it is getting smarter or dumber.
// PHP:
//
//	function calculateMSE($predictions, $targets) {
//	    if (count($predictions) !== count($targets)) { throw new Exception("Length mismatch"); }
//	    $sumError = 0.0;
//	    for ($i = 0; $i < count($predictions); $i++) {
//	        $diff = $predictions[$i] - $targets[$i];
//	        $sumError += ($diff * $diff); // Squaring the difference
//	    }
//	    return $sumError / count($predictions); // Mean (Average)
//	}
func calculateMSE(predictions []float64, targets []float64) float64 {
	if len(predictions) != len(targets) {
		panic("Length mismatch")
	}

	sum := 0.0
	for i := 0; i < len(predictions); i++ {
		diff := predictions[i] - targets[i]
		sum += diff * diff
	}
	return sum / float64(len(targets))
}

func main() {
	// The AI predicted these 3 house prices (in millions)
	predictions := []float64{3.0, 2.5, 4.1}

	// The ACTUAL real-world prices of those houses
	targets := []float64{3.0, 2.0, 5.0}

	// 1st prediction is spot on (Loss: 0)
	// 2nd prediction missed by 0.5
	// 3rd prediction missed by 0.9

	errorScore := calculateMSE(predictions, targets)
	fmt.Printf("The AI's total error (MSE) is: %f\n", errorScore)
}
