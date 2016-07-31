package main

import (
	"flag"
	"fmt"
)

func main() {
	var s = flag.String("s", "Welcome to Go World!", "What you want to reverse?")
	flag.Parse()

	reverse(*s)
}

func reverse(s string) {
	for i := len(s) - 1; i >= 0;i-- {
		fmt.Printf("%s", string(s[i]))
	}
}
