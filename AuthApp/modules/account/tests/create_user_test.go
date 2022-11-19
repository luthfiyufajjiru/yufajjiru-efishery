package tests

import (
	"authapp/adapters"
	"authapp/modules/account/dal"
	"authapp/modules/account/models"
	"authapp/modules/token/services"
	"testing"

	"gorm.io/gorm"
)

var (
	name         string = "Luthfi"
	phone_number        = "+6289999999999"
	user_role           = "user"
	pass                = services.DefaultPassword()
)

func getDb() *gorm.DB {
	x := new(adapters.Adapters)
	x.ActivateSql("../../../authapp.sqlite")
	return x.SQL.DB
}

func TestCreateUser(t *testing.T) {
	x := getDb()

	u := models.UserModel{
		PhoneNumber: phone_number,
		Name:        name,
		Role:        user_role,
		Password:    services.HashAndSaltString(pass),
	}

	defer func() {
		x.Delete(&u)
	}()

	err := dal.CreateUser(x, u)

	if err != nil {
		t.Log("failed to create user")
		t.Fail()
	}

}
