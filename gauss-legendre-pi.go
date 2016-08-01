package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Calculated pi:     %v \n",  pi(1000))

	fmt.Printf("Pi constant value: %v\n", math.Pi)
}


func pi(iterations int) float64 {
	a := 1.0
	b := 1.0 / math.Sqrt(2)
	t := 1.0 / 4.0
	p := 1.0

	for i := 0; i < iterations; i++ {
		an := (a + b) / 2.0
		bn := math.Sqrt(a * b)
		tn := t - p * math.Pow(a - an, 2)
		pn := 2.0 * p

		a, b, t, p = an, bn, tn, pn
	}

	return math.Pow(a + b, 2) / (4.0 * t)
}
