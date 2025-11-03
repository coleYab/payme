package users

import "payme/internal/models"

type UserDto struct {
	ID string `json:"id"`
	Email string `json:"email"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Role string `json:"role"`
}

func FromUserModel(user *models.User) UserDto {
	return UserDto{
		ID: user.ID,
		Email: user.Email,
		FirstName: user.FirstName,
		LastName: user.LastName,
		Role: user.Role,
	}
}
