package simulatedannealing

import (
	"fmt"
	"math"
)

// Simulated Annealing algorithm.
// Leverages the Boltzman factor to add simulated entropy and odds of choosing a less optimal position.
func simulatedAnnealing(inputFunction, stepSize, xMin, xMax, yMin, yMax float64) []float64 {
	var rc = make([]float64, 100)

	return rc
}

// Boltzman factor.
// Determines the ability to accept a less than optimal state.
func boltzmanFactor(expectedOld, expectedNew, eAvg, temp float64) float64 {
	var probability = math.Exp(-(expectedOld - expectedNew) / (eAvg * temp))
	return probability
}

func main() {
	var old, new, avg, temp, test float64
	old = 1
	new = 2
	avg = 0.5
	temp = 100
	test = boltzmanFactor(old, new, avg, temp)
	fmt.Println(test)
}
