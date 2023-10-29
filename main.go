package main

import (
	"time"

	"github.com/robfig/cron"
	log "github.com/sirupsen/logrus"
)

func main() {

	c := cron.New()

	// Same as "0 * * * *"
	err := c.AddFunc("@every 1m", func() {
		log.Info("Tick from CRON")
	})
	if err != nil {
		log.Fatalf("Error scheduling the job: %s", err)
	} else {
		c.Start()
		for _, e := range c.Entries() {
			log.Infof("Have a job scheduled at %s", e.Schedule)
		}
	}

	for {
		log.Info("Ping")
		time.Sleep(3 * time.Second)
	}
}
