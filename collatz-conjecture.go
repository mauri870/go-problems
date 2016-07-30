package main

import (
	"flag"
	"fmt"
	"os"
)	

func main() {
	var n = flag.Int("n", 1, "The number")
	var r = flag.Bool("r", false, "Recalculate recursivelly with the result")
	flag.Parse()
	
	getResult(*n, *r)
}

func getResult(n int, r bool) {

	result := 0
	exit := false
	switch {
		case (n == 1):
			result = 1
			exit = true
		case (n % 2 == 0):
			result = n/2
		default:
			result = n * 3 + 1
	}

	fmt.Printf("%d\n", result)
	
	if exit {
		os.Exit(1)
	} else if r == true {
		getResult(result, r)
	}
}
