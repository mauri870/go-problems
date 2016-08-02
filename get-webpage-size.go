package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)	

type WebPage struct {
	name string
	url string
}

func main() {
	totalReqTime := 0.0
	pages := []WebPage{
		WebPage{"google", "http://google.com"},
		WebPage{"facebook", "http://facebook.com"},
		WebPage{"gmail", "http://gmail.com"},
		WebPage{"microsoft", "http://microsoft.com"},
		WebPage{"you tube", "https://youtube.com"},
		WebPage{"linkedin", "http://linkedin.com"},
	}

	for _, page := range pages {
		startReq := time.Now()

		pageLenght, err := getPageSize(page.url)
		if err != nil {
			log.Fatalf("Error: %s\n", err)
		}

		endReq := time.Now()
		reqDuration := endReq.Sub(startReq).Seconds()
		totalReqTime += reqDuration
		fmt.Printf("The %s html page size is ~%v kb, retrieved in %vs\n", page.name, pageLenght, reqDuration)
	}

	fmt.Printf("Finished %d requests after %vs\n", len(pages), totalReqTime)
}

func getPageSize(url string) (float32, error) {
	res, err := http.Get(url)
	if err != nil {
		return 0.0, err
	}
	defer res.Body.Close()

	bytes, _ := ioutil.ReadAll(res.Body)
	lenght := float32(len(bytes)) / float32(1024)
	return lenght, err
}
