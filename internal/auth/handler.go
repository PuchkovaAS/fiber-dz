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

	authGroup := router.Group("/api")
	authGroup.Post("/register", h.createUser)
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
