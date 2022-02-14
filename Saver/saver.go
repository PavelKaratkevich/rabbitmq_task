package main

import (
	"log"
	"os"
	"rabbitmq/pkg"
)

func main() {
	f, err := os.Create("output.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

	for {
		url := pkg.ReceiveFromQueue("url")
		title := pkg.ReceiveFromQueue("title")

		log.Printf("Output: %v, %v", url, title)

		_, err2 := f.WriteString("\n" + url + "\n" + title + "\n")
		if err2 != nil {
			log.Fatal(err2)
		}
	}
}
