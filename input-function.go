package graphinput

import (
	"fmt"
	"math"
)

// storing the function to vizualize.

// requirments: only can render 3d things for now.
// input: x (x-axis), y (y-axis)
// output: z.
func graphFunction(x float64, y float64) float64 {
	var z float64
	z = float64(math.Sin(math.Pow(x, 2)+3*math.Pow(y, 2))/math.Pow(0.1*math.Sqrt(math.Pow(x, 2)+math.Pow(y, 2)), 2) + (math.Pow(x, 2)+5*math.Pow(y, 2))*math.Exp(1-math.Pow(math.Sqrt(math.Pow(x, 2)+math.Pow(y, 2)), 2))/2)
	return z
}

func main() {
	var x, y, test float64
	x = 1
	y = 0
	test = graphFunction(x, y)
	fmt.Println(test)
}
