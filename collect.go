package main

import (
	"log"

	"github.com/gocolly/colly"
)

func main() {
	Collect("https://go.dev/learn/")
}

type page struct {
	url   string
	title string
}

func Collect(url string) {

	var output page
	c := colly.NewCollector()

	c.OnHTML("head title", func(e *colly.HTMLElement) {
		output.title = e.Text
	})

	output.url = url

	c.Visit(url)

	log.Printf("Result: %v", output)
}
