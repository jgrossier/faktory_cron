package main

import (
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
)

var logger *Log = new(Log)

func configureLogging() {
	log.SetHandler(text.New(os.Stderr))
	logger.Entry = log.WithFields(log.Fields{})
}