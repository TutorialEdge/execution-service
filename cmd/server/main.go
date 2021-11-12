package main

import (
	"github.com/TutorialEdge/execution-service/internal/search"
	"github.com/TutorialEdge/execution-service/internal/transport/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	log "github.com/sirupsen/logrus"
)

// Setup instantiates the app
func Setup() (*fiber.App, error) {
	app := fiber.New()

	app.Use(cors.New(
		cors.Config{
			AllowOrigins: "http://localhost:1313, https://tutorialedge.net, http://localhost:8080",
			AllowHeaders: "*",
		},
	))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("TutorialEdge REST API")
	})

	searchService := search.New()

	httpHandler := http.New(
		searchService,
	)

	httpHandler.SetupRoutes(app)

	log.Info("app setup successful")
	return app, nil
}

func main() {
	log.Info("Starting up the TutorialEdge API")
	var app *fiber.App
	var err error
	if app, err = Setup(); err != nil {
		log.Fatal(err)
	}
	if err := app.Listen(":5000"); err != nil {
		log.Fatal(err)
	}
}
