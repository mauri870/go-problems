package main

import (
	"flag"
	"fmt"
)

func main() {
	var n = flag.Int("n", 5, "How many fizz buzz you want?")
	flag.Parse()

	for i := 1; i < *n + 1;i++ {
		switch {
			case (i % 3 == 0 && i % 5 == 0):
				fmt.Println("FizzBuzz")		
			case (i % 3 == 0):
				fmt.Println("Fizz")	
			case (i % 5 == 0):
				fmt.Println("Buzz")
			default:
				fmt.Println(i)
		}
	}
}
