package main

import (
	"io/ioutil"
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	var u = flag.String("u", "http://google.com", "The url to clone")
	var p = flag.Int("p", 8080, "The port to listen on")
	flag.Parse()

	client := new(http.Client)

	res, err := client.Get(*u)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	os.Mkdir("static", 0755)

	f, err := os.Create("static/index.html")
	if err != nil {	
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
	defer f.Close()

	_, err = f.Write(body)
	if err != nil {	
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	
	fmt.Printf("Listen on port %d\n", *p)
	
	fs := http.FileServer(http.Dir("static"))

	http.Handle("/", fs)
	http.ListenAndServe(fmt.Sprintf(":%d", *p), nil)
}


