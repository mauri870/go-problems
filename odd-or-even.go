package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	input   string
	value   int
	typeRes string
)

func main() {
	for {
		fmt.Print("odd or even? ")
		for {
			fmt.Scanln(&input)
			if input == "odd" || input == "even" {
				break
			}

			fmt.Println("You must respond with odd or even")
		}

		fmt.Print("Say a number: ")
		for {
			fmt.Scanln(&value)

			if value < 10 && value >= 0 {
				break
			}

			fmt.Println("Your number must be an int from 0 to 10")
		}

		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		value2 := r.Intn(10)
		fmt.Println("CPU: ", value2)

		result := value + value2
		isOdd := result%2 == 0
		if isOdd {
			typeRes = "odd"
		} else {
			typeRes = "even"
		}

		fmt.Println("Result: ", result)
		if typeRes == input {
			fmt.Println("You win")
		} else {
			fmt.Println("You lose")
		}

		fmt.Println()
	}
}
