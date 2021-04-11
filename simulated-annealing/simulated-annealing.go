package simulatedannealing

import "math"
import "container/list"
import "fmt"

// Simulated Annealing algorithm.
// Leverages the Boltzman factor to add simulated entropy and odds of choosing a less optimal position.
func simulatedAnnealing(graphFunction func, stepSize float64, xMin float64, xMax float64, yMin float64, yMax float64) list {

}

// Boltzman factor.
// Determines the ability to accept a less than optimal state.
func boltzmanFactor(expectedOld, expectedNew, eAvg, temp float64) float64 {
	var probability = math.Pow(-(expectedOld - expectedNew)/(eAvg * temperature))
	return probability
}

func main() {
	var old, new, avg, temp float64
	old = 1
	new = 2
	avg = 0.5
	temp = 100
	test = boltzmanFactor(old, new, avg, temp)
	fmt.Println(test)
}