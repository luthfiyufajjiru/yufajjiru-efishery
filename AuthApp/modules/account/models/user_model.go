package models

type UserModel struct {
	PhoneNumber string `json:"phone_number" gorm:"primaryKey", "column:phone_number", "index:idx_users_ids"`
	Name        string `json:"name" gorm:"column:name", "index:idx_users_ids"`
	Role        string `json:"role" gorm:"column:user_role"`
	Password    []byte `json:"password" gorm:"column:password"`
}

func (user *UserModel) TableName() string {
	return "users"
}
