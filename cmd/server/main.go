package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lorgioedtech/go-rest-dev/internal/adding"
	"github.com/lorgioedtech/go-rest-dev/internal/http/rest"
	"github.com/lorgioedtech/go-rest-dev/internal/listing"
	"github.com/lorgioedtech/go-rest-dev/internal/storage/json"
)

// StorageType defines available storage types
type Type int

var Version = "development"

const (
	// JSON will store data in JSON files saved on disk
	JSON Type = iota
	// Memory will store data in memory
	Memory
)

func main() {

	// set up storage
	storageType := JSON // this could be a flag; hardcoded here for simplicity

	var adder adding.Service
	var lister listing.Service

	switch storageType {
	case Memory:
		log.Fatal("Not implemented yet")

	case JSON:
		// error handling omitted for simplicity
		s, _ := json.NewStorage()

		adder = adding.NewService(s)
		lister = listing.NewService(s)
	}

	// set up the HTTP server, add more hanlders like rest.Handler(lister, getting, updating)
	router := rest.Handler(adder, lister)

	fmt.Println("The microservices version %s ir runing now: http://localhost:8080", Version)
	log.Fatal(http.ListenAndServe(":8080", router))
}
