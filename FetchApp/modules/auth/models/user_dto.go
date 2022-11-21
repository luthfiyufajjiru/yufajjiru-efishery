package models

type UserDTO struct {
	PhoneNumber string `json:"phone_number"`
	Name        string `json:"name"`
	Role        string `json:"role"`
	Password    string `json:"password"`
}
