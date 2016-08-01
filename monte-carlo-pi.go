package main

import(
	"fmt"
	"log"
	"math"
)

func main() {
	fmt.Printf("Calculated Pi:            %v\n", pi(10))
	fmt.Printf("The value of Pi constant: %v\n", math.Pi)
}

func pi(n int) float64{
	if n > 10 {
		log.Fatalf("The precision must be an int from 1 to 10, got %d", n)
	
	}

	const FACTOR = 100000

	ch := make(chan float64)
	for k := 0; k <= n * FACTOR; k++ {
		go formula(ch, float64(k))
	}

	f := 0.0
	for k := 0; k <= n * FACTOR; k++ {
		f+= <-ch
	}

	return f
}

func formula(ch chan float64, k float64) {
	ch <- 4 * (math.Pow(-1, k) / (2.0*k + 1.0))
}
