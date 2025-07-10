package pages

import (
	"fiber-dz/views"
	"log/slog"

	"github.com/gofiber/fiber/v2"

	templeadapter "fiber-dz/pkg/templ_adapter"
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
	// tags := []string{
	// 	"Еда",
	// 	"Животные",
	// 	"Машины",
	// 	"Спорт",
	// 	"Музыка",
	// 	"Технологии",
	// 	"Прочее",
	// }
	//
	// _ := struct {
	// 	Tags []string
	// }{Tags: tags}
	component := views.Main()

	return templeadapter.Render(c, component)
}
