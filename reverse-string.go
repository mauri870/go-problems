package main

import (
	"fmt"
)

func main() {
	reverse("Welcome to Go World!")
}

func reverse(s string) {
	for i := len(s) - 1; i >= 0;i-- {
		fmt.Printf("%s", string(s[i]))
	}
}
