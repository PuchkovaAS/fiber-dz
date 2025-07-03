package pages

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

type PagesHandler struct {
	router       fiber.Router
	customLogger *slog.Logger
}

func NewHandler(router fiber.Router, customLogger *slog.Logger) {
	h := &PagesHandler{
		router:       router,
		customLogger: customLogger,
	}
	router.Get("/", h.home)
}

func (h *PagesHandler) home(c *fiber.Ctx) error {
	h.customLogger.Info("This is Home Page")
	return c.SendString("Hello")
}
