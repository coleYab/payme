package users

type RegisterUserDto struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
	Role string `json:"role"`
	Password string `json:"Password"`
}

type UpdateUserDto struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}

type UpdateUserPasswordDto struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

type UpdateUserRoleDto struct {
	Role string `json:"role"`
}
