package auth

import (
	"errors"
	"fiber-dz/pkg/di"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepository di.IUserRepository
}

func NewAuthService(userRepository di.IUserRepository) *AuthService {
	return &AuthService{userRepository}
}

func (service *AuthService) Login(
	loginUserForm userLoginForm,
) error {
	existedUser, _ := service.UserRepository.FindByEmail(loginUserForm.Email)

	if existedUser == nil {
		return errors.New(ErrWrongCredentials)
	}

	err := bcrypt.CompareHashAndPassword(
		[]byte(existedUser.Password),
		[]byte(loginUserForm.Password),
	)
	if err != nil {
		return errors.New(ErrWrongCredentials)
	}

	return nil
}

func (service *AuthService) Register(
	createUserForm userCreateForm,
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
