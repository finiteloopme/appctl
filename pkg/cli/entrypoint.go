package cli

import (
	log "github.com/finiteloopme/appctl/pkg/common/log"
)

func Usage() {
	log.Info("Usage:")
	log.Info("appctl deploy --image <URI to OCI container>")
}

func Run() {
	Usage()
}
