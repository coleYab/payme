package services

import (
	"fmt"
	"payme/internal/models"
	"payme/internal/repository"
)

type UserServices struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserServices {
	return &UserServices{userRepository}
}

func (us *UserServices) SaveUser(user *models.User) error {
	return us.userRepository.Save(user)
}

func (us *UserServices) FindAllUsers(skip, limit int) ([]*models.User, error) {
	return us.userRepository.FindUsers()
}

func (us *UserServices) FindUserByID(id string) (*models.User, error) {
	return us.userRepository.FindUserByID(id)
}

func (us *UserServices) FindUserByEmail(email string) (*models.User, error) {
	return us.userRepository.FindUserByEmail(email)
}


func (us *UserServices) UpdateRole(id, newRole string) (*models.User, error) {
	user, err := us.userRepository.FindUserByID(id)
	if err != nil {
		return user, err
	}

	user.UpdateRole(newRole)
	if err := us.userRepository.Update(user); err != nil {
		return user, err
	}

	return user, err
}

func (us *UserServices) UpdateUser(id, firstName, lastName string) (*models.User, error) {
	user, err := us.userRepository.FindUserByID(id)
	if err != nil {
		return nil, err
	}

	user.UpdateUser(firstName, lastName)
	if err := us.userRepository.Update(user); err != nil {
		return nil, err
	}

	return user, err
}

func (us *UserServices) UpdateUserPassword(id, oldPassword, newPassword string) (*models.User, error) {
	user, err := us.userRepository.FindUserByID(id)
	if err != nil {
		return nil, err
	}

	// TODO: change it to compare password
	if user.Password != oldPassword {
		return nil, fmt.Errorf("old password is not valid")
	}
	user.UpdatePassword(newPassword)
	if err := us.userRepository.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (us *UserServices) DeleteUser(id string) error {
	return us.userRepository.Delete(id)
}
