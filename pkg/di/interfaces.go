package di

import "fiber-dz/internal/users"

type IUserRepository interface {
	Create(email, name, password string) error
	FindByEmail(email string) (*users.User, error)
}
