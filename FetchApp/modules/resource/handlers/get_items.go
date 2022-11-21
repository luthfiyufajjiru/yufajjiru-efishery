package handlers

import (
	"fetchapp/modules/auth"
	"fetchapp/modules/auth/models"
	"fetchapp/modules/resource/services"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetItems() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := string(c.Request().Header.Peek("Token"))

		if tokenString == "" {
			return fiber.NewError(http.StatusForbidden, "not allowed")
		}

		user, err := auth.TokenClaim(tokenString)

		if err != nil || (err == nil && user == models.UserDTO{}) {
			return fiber.NewError(http.StatusForbidden, "not allowed")
		}

		data := services.GetResources()

		c.Status(http.StatusOK).JSON(data)

		return err
	}
}
