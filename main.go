package main

import (
	"log"
	"time"
)

func main() {
	for {
		log.Println("Hello")
		time.Sleep(3 * time.Second)
	}
}
