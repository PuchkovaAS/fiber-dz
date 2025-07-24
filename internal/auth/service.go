package auth

import (
	"errors"
	"fiber-dz/internal/users"
	"fiber-dz/pkg/di"

	"github.com/jordan-wright/email"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepository di.IUserRepository
}

func NewAuthService(userRepository di.IUserRepository) *AuthService {
	return &AuthService{userRepository}
}

func (service *AuthService) Register(
	createUserForm users.UserCreateForm,
) error {
	existedUser, _ := service.UserRepository.FindByEmail(createUserForm.Email)

	if existedUser != nil {
		return errors.New(ErrUserExists)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(createUserForm.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	err = service.UserRepository.Create(
		createUserForm.Email,
		createUserForm.Name,
		string(hashedPassword),
	)
	if err != nil {
		return err
	}
	return nil
}
