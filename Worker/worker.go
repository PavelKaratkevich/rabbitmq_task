package main

import (
	"log"
	"rabbitmq/pkg"

	"github.com/gocolly/colly"
)

func main() {
	for {
		url := pkg.ReceiveFromQueue("url")
		output := Collect(url)

		pkg.SendToQueue("urls", output.url + "\n" + output.title + "\n")
	}
}

type page struct {
	url   string
	title string
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
