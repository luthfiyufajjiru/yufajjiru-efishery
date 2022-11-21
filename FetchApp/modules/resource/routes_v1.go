package resource

import (
	"fetchapp/modules/resource/handlers"

	"github.com/gofiber/fiber/v2"
)

func GetRoutesV1() (routes *fiber.App) {
	routes = fiber.New()
	routes.Get("", handlers.GetItems())
	return
}
