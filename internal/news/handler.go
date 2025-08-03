package news

import (
	"fiber-dz/pkg/validator"
	"fiber-dz/views/components"
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofiber/fiber/v2"

	templeadapter "fiber-dz/pkg/templ_adapter"
)

type NewsHandler struct {
	router       fiber.Router
	customLogger *slog.Logger
}

func NewHandler(
	router fiber.Router,
	customLogger *slog.Logger,
) {
	h := &NewsHandler{
		router:       router,
		customLogger: customLogger,
	}

	authGroup := router.Group("/api")
	authGroup.Post("/create", h.createNews)
}

func (h *NewsHandler) createNews(c *fiber.Ctx) error {
	form := newsCreateForm{
		Title:   c.FormValue("title"),
		Preview: c.FormValue("preview"),
		Text:    c.FormValue("text"),
	}

	error := validate.Validate(
		&validators.StringIsPresent{
			Name:    "Title",
			Field:   form.Title,
			Message: "Tile не задан",
		},
		&validators.StringIsPresent{
			Name:    "Preview",
			Field:   form.Preview,
			Message: "preview не задано",
		},
		&validators.StringIsPresent{
			Name:    "Text",
			Field:   form.Text,
			Message: "Текст не задан",
		},
	)
	var component templ.Component
	if len(error.Errors) > 0 {
		component = components.Notification(
			validator.FormatErrors(error),
			components.NotificationFail,
		)
		return templeadapter.Render(c, component, http.StatusBadRequest)
	}

	if err := h.service.Register(form); err != nil {
		component = components.Notification(
			err.Error(),
			components.NotificationFail,
		)
		return templeadapter.Render(c, component, http.StatusBadRequest)
	} else {
		component = components.Notification(
			"Регистрация прошла успешно"+":  "+form.Email,
			components.NotificationSuccess,
		)
	}

	sess, err := h.store.Get(c)
	if err != nil {
		panic(err)
	}
	sess.Set("email", form.Email)
	if err := sess.Save(); err != nil {
		panic(err)
	}
	return templeadapter.Render(c, component, http.StatusCreated)
}
