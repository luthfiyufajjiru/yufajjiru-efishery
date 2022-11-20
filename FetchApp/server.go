package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func RunServer(address string) error {
	app := fiber.New()

	api := app.Group("/api")

	v1collection := api.Group("/v1")
	// v1collection.Mount("/account", account.GetRoutesV1(x))
	// v1collection.Get("/currency", currency.CurrencyConvert())

	fmt.Println(v1collection)

	return app.Listen(address)
}
