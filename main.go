package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	var PLATFORM = os.Getenv("AAP_PLATFORM_ENV")
	if PLATFORM == "SERVICE" {
		log.SetFormatter(&log.JSONFormatter{})
	} else {
		log.SetFormatter(&log.TextFormatter{})
	}
}

func Usage() {
	log.Info("Usage:")
	log.Info("appctl deploy --image <URI to OCI container>")
}

func main() {
	Usage()
}
