package main

import (
	"net/http"

	"github.com/TutorialEdge/execution-service/internal/search"
	transportHTTP "github.com/TutorialEdge/execution-service/internal/transport/http"

	log "github.com/sirupsen/logrus"
)

// Setup instantiates the app
func Setup() error {

	searchService := search.New()
	httpHandler := transportHTTP.New(
		searchService,
	)
	httpHandler.SetupRoutes()

	if err := http.ListenAndServe(":5000", httpHandler.Router); err != nil {
		log.Error("Failed to set up server")
		return err
	}

	log.Info("app setup successful")
	return nil
}

func main() {
	log.Info("Starting up the TutorialEdge API")
	var err error
	if err = Setup(); err != nil {
		log.Fatal(err)
	}
}
