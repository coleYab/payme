package services

import (
	"fmt"
	"payme/internal/models"
	"payme/internal/repository"

	"github.com/google/uuid"
)

type AuthServices struct {
	userRepository repository.UserRepository
}

func NewAuthServices(userRepository repository.UserRepository) *AuthServices {
	return &AuthServices{userRepository}
}

func (us *AuthServices) RegisterUser(email, firstName, lastName, password, role string) (*models.User, error) {
	passwordHash := password
	user, err := models.NewUser(uuid.NewString(), email, firstName, lastName, passwordHash, role);
	if err != nil {
		return  nil, err
	}

	if _, err := us.userRepository.FindUserByEmail(email); err == nil {
		return nil, fmt.Errorf("email already taken")
	}

	return user, nil
}

func (us *AuthServices) LoginUser(email, password string) (*models.User, error) {
	user, err := us.userRepository.FindUserByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	// TODO: compare the password here
	if user.Password != password {
		return nil, fmt.Errorf("invalid credentials")
	}

	return user, err
}
