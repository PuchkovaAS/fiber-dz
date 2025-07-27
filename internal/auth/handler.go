package auth

import (
	"fiber-dz/pkg/validator"
	"fiber-dz/views/components"
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"

	templeadapter "fiber-dz/pkg/templ_adapter"
)

type AuthHandler struct {
	router       fiber.Router
	customLogger *slog.Logger
	service      AuthService
	store        *session.Store
}

func NewHandler(
	router fiber.Router,
	customLogger *slog.Logger,
	service AuthService,
	store *session.Store,
) {
	h := &AuthHandler{
		router:       router,
		customLogger: customLogger,
		service:      service,
		store:        store,
	}

	authGroup := router.Group("/api")
	authGroup.Post("/register", h.createUser)
	authGroup.Post("/login", h.login)
	authGroup.Get("/logout", h.logout)
}

func (h *AuthHandler) logout(c *fiber.Ctx) error {
	sess, err := h.store.Get(c)
	if err != nil {
		panic(err)
	}
	sess.Delete("email")
	if err := sess.Save(); err != nil {
		panic(err)
	}
	c.Response().Header.Add("Hx-Redirect", "/")
	return c.Redirect("/", http.StatusOK)
}

func (h *AuthHandler) login(c *fiber.Ctx) error {
	form := userLoginForm{
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}
	error := validate.Validate(
		&validators.EmailIsPresent{
			Name:    "Email",
			Field:   form.Email,
			Message: "Email не задан или не верный",
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
		c.Set("Content-Type", "text/html")
		return templeadapter.Render(c, component, http.StatusBadRequest)
	}

	if err := h.service.Login(form); err != nil {
		component = components.Notification(
			err.Error(),
			components.NotificationFail,
		)
		c.Set("Content-Type", "text/html")
		return templeadapter.Render(c, component, http.StatusBadRequest)
	}

	sess, err := h.store.Get(c)
	if err != nil {
		panic(err)
	}
	sess.Set("email", form.Email)
	if err := sess.Save(); err != nil {
		panic(err)
	}
	c.Response().Header.Add("Hx-Redirect", "/")
	return c.Redirect("/", http.StatusOK)
}

func (h *AuthHandler) createUser(c *fiber.Ctx) error {
	form := userCreateForm{
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
