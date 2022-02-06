package main

import (
	"fmt"

	"rabbitmq/pkg"

	"github.com/gocolly/colly"
)

func main() {

   urls := Collect("https://go.dev/", 2)
	
   for _, url := range urls {
     pkg.SendToQueue("tasks", url)
   }
}

func Collect(url string, depth int) []string {

	var URLs []string

	c := colly.NewCollector(
		colly.MaxDepth(depth),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {

		fmt.Println("Collecting:", r.URL)
		URLs = append(URLs, r.URL.String())
	})

	c.Visit(url)

	return URLs
}
