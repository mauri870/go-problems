package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
)	

type WebPage struct {
	name string
	url string
}

func main() {
	pages := []WebPage{
		WebPage{"google", "http://google.com"},
		WebPage{"facebook", "http://facebook.com"},
		WebPage{"gmail", "http://gmail.com"},
		WebPage{"microsoft", "http://microsoft.com"},
		WebPage{"you tube", "https://youtube.com"},
		WebPage{"linkedin", "http://linkedin.com"},
	}

	for _, page := range pages {
		res, err := http.Get(page.url)
		if err != nil {
			log.Fatalf("Error: %s", err)	
		}
		defer res.Body.Close()

		bytes, _ := ioutil.ReadAll(res.Body)

		lenght := float32(len(bytes)) / float32(1024)

		fmt.Printf("The %s html page size is ~%v kb\n", page.name, lenght)
	}
}
