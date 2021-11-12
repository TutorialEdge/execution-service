package http

import (
	"github.com/TutorialEdge/execution-service/internal/search"
	"github.com/gofiber/fiber/v2"
)

// Handler -
type Handler struct {
	SearchService search.Service
}

// ErrorResponse -
type ErrorResponse struct {
	Error string
}

// New - returns a new handler
func New(
	searchSvc search.Service,
) Handler {
	return Handler{
		SearchService: searchSvc,
	}
}

// SetupRoutes sets up all the routes for the app
func (h Handler) SetupRoutes(app *fiber.App) {

	app.Get("/.well-known/acme-challenge/4msn8PtbRxmoqZyG7ZSj-sZnK9nYh9Y2E79CIwZsirc", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("4msn8PtbRxmoqZyG7ZSj-sZnK9nYh9Y2E79CIwZsirc.OPUUBETd9RrcpcPnbjR-M-ZFaKPZrbwI_6aDIpwUXd0")
	})
	app.Get("/.well-known/acme-challenge/Po4LqaiYo__kif1nfgD6Zg1hxiHy_sluuKzasLjnJok", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("Po4LqaiYo__kif1nfgD6Zg1hxiHy_sluuKzasLjnJok.OPUUBETd9RrcpcPnbjR-M-ZFaKPZrbwI_6aDIpwUXd0")
	})
	app.Get("/.well-known/acme-challenge/tvEreqlNeyC9C9Cw0lrIQAQQbXu-k10vD0qfXQtDGlc", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("tvEreqlNeyC9C9Cw0lrIQAQQbXu-k10vD0qfXQtDGlc.OPUUBETd9RrcpcPnbjR-M-ZFaKPZrbwI_6aDIpwUXd0")
	})

	app.Get("/v1/health", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("Healthy")
	})
	app.Post("/v1/execute", h.ExecuteChallenge) // ExecuteChallenge

	searchGroup := app.Group("/v1/search")
	searchGroup.Post("/", h.Search)
}
