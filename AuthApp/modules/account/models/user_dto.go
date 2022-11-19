package models

import (
	"authapp/modules/token/services"
)

type UserDTO struct {
	PhoneNumber string `json:"phone_number"`
	Name        string `json:"name"`
	Role        string `json:"role"`
	Password    string `json:"password"`
}

func (u *UserDTO) MapToUserModel() UserModel {
	if u.Password == "" {
		u.Password = services.DefaultPassword()
	}

	return UserModel{
		PhoneNumber: u.PhoneNumber,
		Name:        u.Name,
		Role:        u.Role,
		Password:    services.HashAndSaltString(u.Password)}
}
