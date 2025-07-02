package pages

import "github.com/gofiber/fiber/v2"

type PagesHandler struct {
	router fiber.Router
}

func NewHandler(router fiber.Router) {
	h := &PagesHandler{
		router: router,
	}
	router.Get("/", h.home)
}

func (h *PagesHandler) home(c *fiber.Ctx) error {
	return c.SendString("Hello")
}
