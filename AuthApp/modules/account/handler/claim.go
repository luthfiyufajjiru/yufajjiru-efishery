package handler

import (
	"authapp/modules/token"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

func TokenClaim() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := string(c.Request().Header.Peek("Token"))

		claim, err := token.Claims(tokenString)

		date, _ := claim["timestamp"].(float64)

		if err != nil || (date < float64(time.Now().Unix())) {
			return fiber.NewError(http.StatusBadRequest, "token not valid")
		}

		c.Status(http.StatusOK).JSON(claim)

		return nil

	}
}
