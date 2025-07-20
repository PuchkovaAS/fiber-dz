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
	router.Get("/register", h.register)
}

func (h *PagesHandler) register(c *fiber.Ctx) error {
	tags := []views.TagData{
		{
			Name:    "Еда",
			PathImg: "./public/images/food/01.png",
			Url:     "/",
		},
		{
			Name:    "Животные",
			PathImg: "/public/images/animal/10.png",
			Url:     "/",
		},

		{
			Name:    "Машины",
			PathImg: "/public/images/car/04.png",

			Url: "/",
		},

		{
			Name:    "Спорт",
			PathImg: "/public/images/sport/06.png",
			Url:     "/",
		},

		{
			Name:    "Музыка",
			PathImg: "/public/images/music/06.png",
			Url:     "/",
		},

		{
			Name:    "Технологии",
			PathImg: "/public/images/technology/03.png",
			Url:     "/",
		},

		{
			Name:    "Прочее",
			PathImg: "/public/images/abstract/07.png",
			Url:     "/",
		},
	}
	component := views.Registration(tags)

	return templeadapter.Render(c, component)
}

func (h *PagesHandler) home(c *fiber.Ctx) error {
	tags := []views.TagData{
		{
			Name:    "Еда",
			PathImg: "./public/images/food/01.png",
			Url:     "/",
		},
		{
			Name:    "Животные",
			PathImg: "/public/images/animal/10.png",
			Url:     "/",
		},

		{
			Name:    "Машины",
			PathImg: "/public/images/car/04.png",

			Url: "/",
		},

		{
			Name:    "Спорт",
			PathImg: "/public/images/sport/06.png",
			Url:     "/",
		},

		{
			Name:    "Музыка",
			PathImg: "/public/images/music/06.png",
			Url:     "/",
		},

		{
			Name:    "Технологии",
			PathImg: "/public/images/technology/03.png",
			Url:     "/",
		},

		{
			Name:    "Прочее",
			PathImg: "/public/images/abstract/07.png",
			Url:     "/",
		},
	}
	component := views.Main(tags)

	return templeadapter.Render(c, component)
}
