package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/lorgioedtech/go-rest-dev/internal/adding"
	"github.com/lorgioedtech/go-rest-dev/internal/config"
	"github.com/lorgioedtech/go-rest-dev/internal/http/rest"
	"github.com/lorgioedtech/go-rest-dev/internal/listing"
	"github.com/lorgioedtech/go-rest-dev/internal/storage/json"
	"github.com/lorgioedtech/go-rest-dev/pkg/log"
)

// StorageType defines available storage types
type Type int

var Version = "development"
var ConfigEnvironment = "local"
var flagConfig = flag.String("config", "./config/local.yml", "path to the config file")

const (
	// JSON will store data in JSON files saved on disk
	JSON Type = iota
	// Memory will store data in memory
	Memory
)

func main() {
	flag.Parse()
	// create root logger tagged with server version
	logger := log.New().With(nil, "version", Version)

	// load application configurations
	cfg, err := config.Load(*flagConfig, logger)
	if err != nil {
		logger.Errorf("failed to load application configuration: %s", err)
		os.Exit(-1)
	}

	logger.Infof("server %v database configuration uri %v", Version, cfg.DatabaseURI)

	// set up storage
	storageType := JSON // this could be a flag; hardcoded here for simplicity

	var adder adding.Service
	var lister listing.Service

	switch storageType {
	case Memory:
		logger.Errorf("Not implemented yet")

	case JSON:
		// error handling omitted for simplicity
		s, _ := json.NewStorage()

		adder = adding.NewService(s)
		lister = listing.NewService(s)
	}

	// set up the HTTP server, add more hanlders like rest.Handler(lister, getting, updating)
	router := rest.Handler(logger, adder, lister)

	logger.Infof("The microservices version %v ir runing now: http://localhost:8080", Version)
	if err := http.ListenAndServe(":8080", router); err != nil && err != http.ErrServerClosed {
		logger.Error(err)
		os.Exit(-1)
	}
}
