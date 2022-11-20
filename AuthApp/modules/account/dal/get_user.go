package dal

import (
	"authapp/modules/account/models"
	"authapp/modules/token/services"

	"gorm.io/gorm"
)

func GetUser(db *gorm.DB, user models.UserModel) (models.UserModel, error) {
	err := db.First(&user).Error
	return user, err
}

func ValidateUser(db *gorm.DB, user models.UserModel, pass string) (models.UserModel, bool) {
	err := db.First(&user).Error
	if err != nil {
		return models.UserModel{}, false
	}
	return user, services.CompareHashAndString(user.Password, pass)
}
