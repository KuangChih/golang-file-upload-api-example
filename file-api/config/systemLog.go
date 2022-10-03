package config

import (
	log "github.com/sirupsen/logrus"
)

func LogInit() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
}
