package tests

import (
	"authapp/modules/account/dal"
	"authapp/modules/account/models"
	"authapp/modules/token/services"
	"fmt"
	"testing"
)

func TestSignIn(t *testing.T) {
	fmt.Println("Test create then fetch and assert the password between them")

	x := getDb()

	u := models.UserModel{
		PhoneNumber: phone_number,
		Name:        name,
		Role:        user_role,
		Password:    services.HashAndSaltString(pass),
	}

	defer func() {
		_ = x.Delete(&u).Error
	}()

	err := dal.CreateUser(x, u)

	if err != nil {
		t.Fatalf("failed to create user: %s", err.Error())
	}

	user, err := dal.GetUser(x, models.UserModel{PhoneNumber: phone_number})

	if v := services.CompareHashAndString(user.Password, pass); !v {
		t.Fatalf("login failed: %s", err.Error())
	}

}
