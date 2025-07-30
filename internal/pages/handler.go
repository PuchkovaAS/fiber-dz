package pages

import (
	"fiber-dz/internal/news"
	"fiber-dz/views"
	"log/slog"
	"math"
	"net/http"

	"github.com/gofiber/fiber/v2"

	templeadapter "fiber-dz/pkg/templ_adapter"
)

type PagesHandler struct {
	router         fiber.Router
	customLogger   *slog.Logger
	newsRepository *news.NewsRepository
}

func NewHandler(
	router fiber.Router,
	customLogger *slog.Logger,
	newsRepository *news.NewsRepository,
) {
	h := &PagesHandler{
		router:         router,
		customLogger:   customLogger,
		newsRepository: newsRepository,
	}
	router.Get("/", h.home)
	router.Get("/register", h.register)
	router.Get("/login", h.login)
}

func (h *PagesHandler) login(c *fiber.Ctx) error {
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
	component := views.Login(tags)

	return templeadapter.Render(c, component, http.StatusOK)
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

	return templeadapter.Render(c, component, http.StatusOK)
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

	PAGE_ITEMS := 4
	page := c.QueryInt("page", 1)

	count := h.newsRepository.CountAll()
	news, err := h.newsRepository.GetAll(PAGE_ITEMS, (page-1)*PAGE_ITEMS)
	if err != nil {
		h.customLogger.Error(err.Error())
		return c.SendStatus(500)
	}

	component := views.Main(
		news,
		int(math.Ceil(float64(count)/float64(PAGE_ITEMS))),
		page,
		tags,
	)

	return templeadapter.Render(c, component, http.StatusOK)
}
