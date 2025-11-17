package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"

	c	"ex/controller"
	valid	"ex/validator"
)

func main() {
	valid.Init()

	app := fiber.New()

	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"gofiber": "21022566",
		},
	}))

	apiv1 := app.Group("/api/v1")
	apiv1.Get("/fact/:num", c.Factorial)
	apiv1.Post("/register", c.CreateUser)

	apiv3 := app.Group("/api/v3")
	apiv3.Post("/tan", c.ToAscii)
	// apiv3.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Get method")
	// })

	app.Listen(":3000")
}
