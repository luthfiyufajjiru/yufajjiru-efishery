package main

import (
	"fetchapp/modules/resource"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func RunServer(address string) error {
	app := fiber.New()

	api := app.Group("/api")

	v1collection := api.Group("/v1")
	v1collection.Mount("/items", resource.GetRoutesV1())
	// v1collection.Get("/currency", currency.CurrencyConvert())

	fmt.Println(v1collection)

	return app.Listen(address)
}
