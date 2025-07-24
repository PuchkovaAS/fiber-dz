package auth

import (
	"fiber-dz/internal/users"
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

type AuthHandler struct {
	router       fiber.Router
	customLogger *slog.Logger
	service      AuthService
}

func NewHandler(
	router fiber.Router,
	customLogger *slog.Logger,
	service AuthService,
) {
	h := &AuthHandler{
		router:       router,
		customLogger: customLogger,
		service:      service,
	}

	authGroup := router.Group("/api")
	authGroup.Post("/register", h.createUser)
}

func (h *AuthHandler) createUser(c *fiber.Ctx) error {
	form := users.UserCreateForm{
		Email:    c.FormValue("email"),
		Name:     c.FormValue("name"),
		Password: c.FormValue("password"),
	}

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

	return templeadapter.Render(c, component, http.StatusCreated)
}
