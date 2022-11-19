package dal

import (
	"authapp/modules/account/models"

	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user models.UserModel) error {
	err := db.Create(&user).Error
	return err
}
