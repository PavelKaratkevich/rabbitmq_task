package main

import (
	"log"
	"rabbitmq/pkg"

	"github.com/gocolly/colly"
)

func main() {
	for {
		url := pkg.ReceiveFromQueue("tasks")
		output := Collect(url)

		pkg.SendToQueue(output.title)
		pkg.SendToQueue(output.url)
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
