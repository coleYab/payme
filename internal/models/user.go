package models

import "fmt"

type User struct {
	ID string `bson:"id"`
	Email string `bson:"email"`
	FirstName string `bson:"firstName"`
	LastName string `bson:"lastName"`
	Password string `bson:"password"`
	Verified string `bson:"verified"`
	Role string `bson:"role"`
}

func NewUser(id, email, firstName, lastName, password, role string) (*User, error) {
	if role == "admin" {
		return nil, fmt.Errorf("user role cannot be admin")
	}

	return &User{
		ID: id,
		Email: email,
		FirstName: firstName,
		LastName: lastName,
		Password: password,
		Verified: "VERIFIED",
		Role: role,
	}, nil
}

func (u *User)UpdateRole(role string) {
	u.Role = role;
}


func (u *User)UpdatePassword(password string) {
	u.Password = password;
}

func (u *User)UpdateUser(firstName, lastName string) {
	u.FirstName = firstName
	u.LastName = lastName
}
