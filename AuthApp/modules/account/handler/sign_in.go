package handler

import (
	"authapp/adapters"
	"net/http"

	"authapp/modules/account/dal"
	am "authapp/modules/account/models"
	"authapp/modules/account/services"
	"authapp/modules/token"

	"github.com/gofiber/fiber/v2"
)

func SignIn(x *adapters.Adapters) fiber.Handler {
	return func(c *fiber.Ctx) (err error) {

		var u am.UserDTO
		err = c.BodyParser(&u)
		if err != nil {
			return fiber.NewError(http.StatusBadRequest, err.Error())
		} else if pn := services.PhoneNumberValidation(u.PhoneNumber); !pn && u.Password == "" {
			return fiber.NewError(http.StatusBadRequest, "invalid input")
		}

		user, valid := dal.ValidateUser(x.SQL.DB, am.UserModel{PhoneNumber: u.PhoneNumber}, u.Password)

		if !valid {
			return fiber.NewError(http.StatusUnauthorized, "not authorized")
		}

		tokenString, err := token.GenerateJWT(user.PhoneNumber, user.Name, user.Role)

		c.Status(http.StatusOK).SendString(tokenString)

		return
	}
}
