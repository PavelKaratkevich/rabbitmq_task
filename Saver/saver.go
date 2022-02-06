package main

import (
	"log"
	"rabbitmq/pkg"
)

func main() {
	for {
		url := pkg.ReceiveFromQueue("tasks_queue")

		log.Printf("Output: %v", url)
	}
}