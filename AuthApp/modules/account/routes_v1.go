package account

import (
	"authapp/adapters"
	"authapp/modules/account/handler"

	"github.com/gofiber/fiber/v2"
)

func GetRoutesV1(x *adapters.Adapters) (routes *fiber.App) {
	routes = fiber.New()
	routes.Post("/register", handler.SignUp(x))
	return
}
