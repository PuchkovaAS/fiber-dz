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
	tags := []views.TagData{
		{
			Name:    "Еда",
			PathImg: "./public/images/food/01.png",
		},
		{
			Name:    "Животные",
			PathImg: "/public/images/animal/10.png",
		},

		{
			Name:    "Машины",
			PathImg: "/public/images/car/04.png",
		},

		{
			Name:    "Спорт",
			PathImg: "/public/images/sport/06.png",
		},

		{
			Name:    "Музыка",
			PathImg: "/public/images/music/06.png",
		},

		{
			Name:    "Технологии",
			PathImg: "/public/images/technology/03.png",
		},

		{
			Name:    "Прочее",
			PathImg: "/public/images/abstract/07.png",
		},
	}
	component := views.Main(tags)

	return templeadapter.Render(c, component)
}
