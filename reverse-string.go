package main

import (
	"flag"
	"fmt"
)

func main() {
	var s = flag.String("s", "Welcome to Go World!", "What you want to reverse?")
	flag.Parse()

	fmt.Println(reverse(*s))
}

func reverse(s string) string {
	reversed := ""
	for i := len(s) - 1; i >= 0;i-- {
		reversed += string(s[i])
	}

	return reversed
}
