package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"tcp-over-http/server"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("usage: %s <config>", os.Args[0])
	}

	config, err := server.NewConfigFromFile(os.Args[1])
	if err != nil {
		log.WithField("err", err).Fatal("loading config")
	}

	if config.RedirectorAddr != "" {
		go func() {
			if err := server.RunRedirectorServer(config); err != nil {
				log.WithField("err", err).Fatal("running redirector")
			}
		}()
	}

	if err := server.RunHTTPServer(config); err != nil {
		log.WithField("err", err).Fatal("running server")
	}
}
