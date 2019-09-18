package main

import (
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	for {
		log.Info("Hello")
		time.Sleep(3 * time.Second)
	}
}
