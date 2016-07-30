package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)	
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!" )
}
