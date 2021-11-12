package auth

import (
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

// Protected protect routes
func Protected(next func(c *fiber.Ctx) error) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		log.Info(c.Get("Authorization"))
		if c.Get("Authorization") == "" {
			return c.Status(403).SendString("Not Authorized")
		} else {
			return next(c)
		}
	}
}
