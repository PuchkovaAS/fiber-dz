package news

import (
	"fiber-dz/pkg/validator"
	"fiber-dz/views/components"
	"log/slog"
	"path/filepath"

	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/google/uuid"

	templeadapter "fiber-dz/pkg/templ_adapter"
)

type NewsHandler struct {
	router       fiber.Router
	customLogger *slog.Logger
	store        *session.Store
	repository   NewsRepository
}

func NewHandler(
	router fiber.Router,
	customLogger *slog.Logger,
	store *session.Store,
	repository NewsRepository,
) {
	h := &NewsHandler{
		router:       router,
		customLogger: customLogger,
		store:        store,
		repository:   repository,
	}

	authGroup := router.Group("/api")
	authGroup.Post("/news/create", h.createNews)
}

func (h *NewsHandler) createNews(c *fiber.Ctx) error {
	// Получаем файл превью
	previewFile, err := c.FormFile("preview_image")
	if err != nil {
		return templeadapter.Render(c,
			components.Notification(
				"Необходимо загрузить изображение для превью",
				components.NotificationFail,
			),
			fiber.StatusBadRequest,
		)
	}

	// Получаем данные формы
	form := newsCreateForm{
		Title:   c.FormValue("title"),
		Preview: previewFile,
		Text:    c.FormValue("text"),
	}

	// Валидация данных
	validationErr := validate.Validate(
		&validators.StringIsPresent{
			Name:    "Title",
			Field:   form.Title,
			Message: "Заголовок не задан",
		},
		&validators.StringIsPresent{
			Name:    "Text",
			Field:   form.Text,
			Message: "Текст не задан",
		},
	)

	if len(validationErr.Errors) > 0 {
		return templeadapter.Render(c,
			components.Notification(
				validator.FormatErrors(validationErr),
				components.NotificationFail,
			),
			fiber.StatusBadRequest,
		)
	}

	// Генерируем уникальное имя файла
	filename := uuid.New().String() + filepath.Ext(form.Preview.Filename)
	savePath := filepath.Join("public", "images", filename)

	// Сохраняем файл на сервере
	if err := c.SaveFile(form.Preview, savePath); err != nil {
		return templeadapter.Render(c,
			components.Notification(
				"Ошибка при сохранении изображения",
				components.NotificationFail,
			),
			fiber.StatusInternalServerError,
		)
	}

	// Получаем email пользователя из сессии
	sess, err := h.store.Get(c)
	if err != nil {
		return templeadapter.Render(c,
			components.Notification(
				"Ошибка сессии",
				components.NotificationFail,
			),
			fiber.StatusInternalServerError,
		)
	}

	userEmail, ok := sess.Get("email").(string)
	if !ok || userEmail == "" {
		return templeadapter.Render(c,
			components.Notification(
				"Требуется авторизация",
				components.NotificationFail,
			),
			fiber.StatusUnauthorized,
		)
	}

	// Создаем параметры для добавления новости
	newsParam := AddNewsParam{
		UserEmail: userEmail,
		Title:     form.Title,
		Text:      form.Text,
		Preview:   "/public/images/" + filename, // Сохраняем относительный путь
	}

	// Добавляем новость в БД
	if err := h.repository.AddNews(&newsParam); err != nil {
		return templeadapter.Render(c,
			components.Notification(
				"Ошибка при записи новости в БД, попробуйте позже",
				components.NotificationFail,
			),
			fiber.StatusInternalServerError,
		)
	}

	// Возвращаем успешный результат
	return templeadapter.Render(c,
		components.Notification(
			"Новость успешно создана",
			components.NotificationSuccess,
		),
		fiber.StatusCreated,
	)
}
