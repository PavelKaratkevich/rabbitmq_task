package main

import (
	"fmt"

	"rabbitmq/pkg"

	"github.com/gocolly/colly"
)

func main() {

	c := make(chan string)

	go Collect("https://go.dev/", 2, c)
	
	for {
		pkg.SendToQueue("url", <-c)
	}
}

func Collect(url string, depth int, m chan string) {

	var URL string

	c := colly.NewCollector(
		colly.MaxDepth(depth),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {

		fmt.Println("Collecting:", r.URL)
		URL = r.URL.String()
		m <- URL
	})

	c.Visit(url)

}
