package main

import (
	"log"
	"os"
	"rabbitmq/pkg"
)

func main() {
	var url string
	for {
		url = pkg.ReceiveFromQueue("tasks_queue")

		log.Printf("Output: %v", url)
		os.WriteFile("output.txt", []byte(url), 0666)
	}
}
