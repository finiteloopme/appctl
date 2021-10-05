package log

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

func Info(msg string) {
	log.Info(msg)

	return
}
