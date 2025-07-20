package auth

import (
	"fiber-dz/pkg/validator"
	"fiber-dz/views"
	"fiber-dz/views/components"
	"log/slog"

	"github.com/a-h/templ"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofiber/fiber/v2"

	templeadapter "fiber-dz/pkg/templ_adapter"
)

type AuthHandler struct {
	router       fiber.Router
	customLogger *slog.Logger
}

func NewHandler(router fiber.Router, customLogger *slog.Logger) {
	h := &AuthHandler{
		router:       router,
		customLogger: customLogger,
	}

	authGroup := router.Group("/api/register")
	authGroup.Get("/", h.register)
	authGroup.Post("/", h.createUser)
}

func (h *AuthHandler) createUser(c *fiber.Ctx) error {
	form := AuthRegisterForm{
		Email:    c.FormValue("email"),
		Name:     c.FormValue("name"),
		Password: c.FormValue("password"),
	}
	h.customLogger.Info("%s", form.Email)

	error := validate.Validate(
		&validators.EmailIsPresent{
			Name:    "Email",
			Field:   form.Email,
			Message: "Email не задан или не верный",
		},
		&validators.StringIsPresent{
			Name:    "Name",
			Field:   form.Name,
			Message: "Имя не задано",
		},
		&validators.StringIsPresent{
			Name:    "Password",
			Field:   form.Password,
			Message: "Пароль не задан",
		},
	)
	var component templ.Component
	if len(error.Errors) > 0 {
		component = components.Notification(
			validator.FormatErrors(error),
			components.NotificationFail,
		)
	} else {
		component = components.Notification(
			"Регистрация прошла успешно",
			components.NotificationSuccess,
		)
	}
	return templeadapter.Render(c, component)
}

func (h *AuthHandler) register(c *fiber.Ctx) error {
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
