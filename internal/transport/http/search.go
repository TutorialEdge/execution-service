package http

import "github.com/gofiber/fiber/v2"

func (h *Handler) Search(c *fiber.Ctx) error {
	query := c.Query("query")
	results, err := h.SearchService.Search(query)
	if err != nil {
		return c.SendString("error")
	}
	return c.JSON(results)
}
