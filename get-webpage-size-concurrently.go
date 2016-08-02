package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)	

type WebPage struct {
	name string
	url string
}

type WebPageResult struct {
	WebPage
	lenght float32
	duration time.Duration
}

func main() {
	start := time.Now()

	pages := []WebPage{
		WebPage{"google", "http://google.com"},
		WebPage{"facebook", "http://facebook.com"},
		WebPage{"gmail", "http://gmail.com"},
		WebPage{"microsoft", "http://microsoft.com"},
		WebPage{"you tube", "https://youtube.com"},
		WebPage{"linkedin", "http://linkedin.com"},
	}

	resultCh := make(chan WebPageResult, len(pages))

	for _, page := range pages {		
		go getPageSize(page, resultCh)
	}

	for i := 1; i <= len(pages); i++ {
		webpage := <-resultCh

		fmt.Printf("The %s html page size is ~%v kb in %vs\n", webpage.name, webpage.lenght, webpage.duration.Seconds())
	}	

	end := time.Now()
	lifetime := end.Sub(start).Seconds()

	fmt.Printf("Finished %d requests after %vs\n", len(pages), lifetime)
}

func getPageSize(page WebPage, resultCh chan<- WebPageResult) {
	init := time.Now()

	res, err := http.Get(page.url)
	if err != nil {
		resultCh <- WebPageResult{page, 0.0, 0}
		return
	}
	defer res.Body.Close()

	bytes, _ := ioutil.ReadAll(res.Body)
	lenght := float32(len(bytes)) / float32(1024)

	duration := time.Now().Sub(init)

	resultCh <- WebPageResult{page, lenght, duration}
}
