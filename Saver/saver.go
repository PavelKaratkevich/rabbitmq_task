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
		out := pkg.ReceiveFromQueue("urls")

		log.Printf("Output: %v", out)

		_, err2 := f.WriteString(out)
		if err2 != nil {
			log.Fatal(err2)
		}
	}
}
