package main

import (
	"log"
	"rabbitmq/pkg"

	"github.com/gocolly/colly"
)

type page struct {
	url   string
	title string
}

func main() {
	for {
		url := pkg.ReceiveFromQueue("url")
		output := Collect(url)

		pkg.SendToQueue("url", output.url)
		pkg.SendToQueue("title", output.title)
	}
}

func Collect(url string) page {

	var output page
	c := colly.NewCollector()

	c.OnHTML("head title", func(e *colly.HTMLElement) {
		output.title = e.Text
	})

	output.url = url

	c.Visit(url)

	log.Printf("Result: %v", output)
	return output
}
