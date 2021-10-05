package log

import (
	"os"

	"github.com/finiteloopme/appctl/pkg/apis/core/v1alpha1"
	log "github.com/sirupsen/logrus"
)

func init() {
	var LOG_FORMAT = os.Getenv("LOG_FORMAT")

	if LOG_FORMAT == v1alpha1.LogFormatJSON.String() {
		log.SetFormatter(&log.JSONFormatter{})
	} else {
		log.SetFormatter(&log.TextFormatter{})
	}
}

func Info(msg string) {
	log.Info(msg)

	return
}
