package pages

import (
	"fiber-dz/internal/news"
	"fiber-dz/views"
	"log/slog"
	"math"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"

	templeadapter "fiber-dz/pkg/templ_adapter"
)

type PagesHandler struct {
	router         fiber.Router
	customLogger   *slog.Logger
	newsRepository *news.NewsRepository
	store          *session.Store
}

func NewHandler(
	router fiber.Router,
	customLogger *slog.Logger,
	newsRepository *news.NewsRepository,
	store *session.Store,
) {
	h := &PagesHandler{
		router:         router,
		customLogger:   customLogger,
		newsRepository: newsRepository,
		store:          store,
	}
	router.Get("/", h.home)
	router.Get("/register", h.register)
	router.Get("/login", h.login)
	router.Get("/searching", h.searching)
	router.Get("/article/:alias", h.article)
	router.Get("/create", h.createNews)
}

func GetTags() []views.TagData {
	return []views.TagData{
		{
			Name:    "Еда",
			PathImg: "./public/images/food/01.png",
			Url:     "/searching?category=еда",
		},
		{
			Name:    "Животные",
			PathImg: "/public/images/animal/10.png",
			Url:     "/searching?category=животные",
		},
		{
			Name:    "Машины",
			PathImg: "/public/images/car/04.png",
			Url:     "/searching?category=машины",
		},
		{
			Name:    "Спорт",
			PathImg: "/public/images/sport/06.png",
			Url:     "/searching?category=спорт",
		},
		{
			Name:    "Музыка",
			PathImg: "/public/images/music/06.png",
			Url:     "/searching?category=музыка",
		},
		{
			Name:    "Технологии",
			PathImg: "/public/images/technology/03.png",
			Url:     "/searching?category=технологии",
		},
		{
			Name:    "Прочее",
			PathImg: "/public/images/abstract/07.png",
			Url:     "/searching?category=прочее",
		},
	}
}

func (h *PagesHandler) createNews(c *fiber.Ctx) error {
	email, ok := c.Locals("email").(string)
	if !ok || email == "" {
		// Для HTMX-запросов
		if c.Get("HX-Request") == "true" {
			c.Response().Header.Add("HX-Redirect", "/login")
			return c.SendStatus(fiber.StatusOK)
		}
		// Для обычных запросов
		return c.Redirect("/login")
	}
	tags := GetTags()
	component := views.AddArticle(tags)

	return templeadapter.Render(c, component, http.StatusOK)
}

func (h *PagesHandler) article(c *fiber.Ctx) error {
	tags := GetTags()
	alias := c.Params("alias")

	news, err := h.newsRepository.GetArticleByAlias(alias)
	if err != nil {
		h.customLogger.Error(err.Error())
		return c.SendStatus(500)
	}
	component := views.Article(news, tags)

	return templeadapter.Render(c, component, http.StatusOK)
}

func (h *PagesHandler) login(c *fiber.Ctx) error {
	tags := GetTags()
	component := views.Login(tags)

	return templeadapter.Render(c, component, http.StatusOK)
}

func (h *PagesHandler) register(c *fiber.Ctx) error {
	tags := GetTags()
	component := views.Registration(tags)

	return templeadapter.Render(c, component, http.StatusOK)
}

func (h *PagesHandler) home(c *fiber.Ctx) error {
	PAGE_ITEMS := 4
	page := c.QueryInt("page", 1)

	count := h.newsRepository.CountAll()
	news, err := h.newsRepository.GetAll(PAGE_ITEMS, (page-1)*PAGE_ITEMS)
	if err != nil {
		h.customLogger.Error(err.Error())
		return c.SendStatus(500)
	}

	tags := GetTags()
	component := views.Main(
		news,
		int(math.Ceil(float64(count)/float64(PAGE_ITEMS))),
		page,
		tags,
	)

	return templeadapter.Render(c, component, http.StatusOK)
}

type SearchQuery struct {
	Category string `query:"category"`
	Keyword  string `query:"keyword"`
	Page     int    `query:"page"`
}

func (h *PagesHandler) searching(c *fiber.Ctx) error {
	PAGE_ITEMS := 8
	query := new(SearchQuery)
	if query.Page < 1 {
		query.Page = 1
	}
	if err := c.QueryParser(query); err != nil {
		h.customLogger.Error(err.Error())
		return c.SendStatus(http.StatusBadRequest)
	}
	params := news.SearchParam{
		Limit:    PAGE_ITEMS,
		Offset:   (query.Page - 1) * PAGE_ITEMS,
		Category: query.Category,
		Keyword:  query.Keyword,
	}
	count := h.newsRepository.CountByParam(params)
	news, err := h.newsRepository.GetByParam(params)
	if err != nil {
		h.customLogger.Error(err.Error())
		return c.SendStatus(500)
	}

	var titlePage string
	if query.Category != "" {
		titlePage = query.Category
	} else {
		titlePage = query.Keyword
	}

	tags := GetTags()
	component := views.Searching(
		titlePage,
		news,
		int(math.Ceil(float64(count)/float64(PAGE_ITEMS))),
		query.Page,
		tags,
	)

	return templeadapter.Render(c, component, http.StatusOK)
}
