package main

import (
	"flag"
	"fmt"
	"math"
)

func main() {
	var n = flag.Int("n", 5, "How many fibonacci numbers you want?")
	flag.Parse()	

	for i := 1; i <= *n; i++ {
		fmt.Println(getFib(i))
	}
}

func getFib(n int) int {
	return Round(math.Pow((math.Sqrt(5) + 1)/2, float64(n)) / math.Sqrt(5))
}

func Round(val float64) int {
    if val < 0 {
        return int(val-0.5)
    }
 	return int(val+0.5)
}


