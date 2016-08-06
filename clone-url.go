package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	var u = flag.String("u", "http://google.com", "The url to clone")
	var p = flag.Int("p", 8080, "The port to listen on")
	flag.Parse()

	client := new(http.Client)

	res, err := client.Get(*u)
	if err != nil {
		log.Print(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	log.Printf("Listen on port %d\n", *p)
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, string(body))
	})
	http.ListenAndServe(fmt.Sprintf(":%d", *p), nil)
}