package users

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UsersRepository struct {
	Dbpool       *pgxpool.Pool
	CustomLogger *slog.Logger
}

func NewUsersRepository(
	dbpool *pgxpool.Pool,
	customLogger *slog.Logger,
) *UsersRepository {
	return &UsersRepository{
		Dbpool:       dbpool,
		CustomLogger: customLogger,
	}
}

func (r *UsersRepository) Create(email, name, password string) error {
	query := `INSERT INTO users (email, name, password) VALUES (@email, @name, @password)`

	args := pgx.NamedArgs{
		"email":    email,
		"name":     name,
		"password": password,
	}
	_, err := r.Dbpool.Exec(context.Background(), query, args)
	if err != nil {
		return fmt.Errorf("Невозможно создать вакансию: %s", err)
	}
	return nil
}

func (r *UsersRepository) FindByEmail(email string) (*User, error) {
	query := "SELECT email, name, password FROM users WHERE email = $1"

	var user User
	err := r.Dbpool.QueryRow(context.Background(), query, email).Scan(
		&user.Email,
		&user.Name,
		&user.Password,
	)
	if err != nil {
		return nil, fmt.Errorf("Ошибка получения данных: %w", err)
	}

	return &user, nil
}
