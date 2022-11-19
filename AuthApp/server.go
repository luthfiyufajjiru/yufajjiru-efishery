package main

import (
	"authapp/adapters"
	"authapp/modules/account"
	"authapp/modules/account/models"

	"github.com/gofiber/fiber/v2"
)

func RunServer(address string) error {

	x := new(adapters.Adapters)
	x.ActivateSql("authapp.sqlite")

	if table := x.SQL.DB.Migrator().HasTable(&models.UserModel{}); !table {
		x.SQL.DB.AutoMigrate(&models.UserModel{})
	}

	app := fiber.New()

	api := app.Group("/api")

	v1collection := api.Group("/v1")
	v1collection.Mount("/", account.GetRoutesV1(x))

	return app.Listen(address)
}
