package main

import (
	"github.com/gofiber/fiber/v2"
	"jubelio.com/chat/modules/chats/handlers"
	"jubelio.com/chat/modules/chats/middlewares"
)

func main() {
	app := fiber.New()
	app.Use(middlewares.VerifyBasicAuth())

	handler := handlers.New()
	handler.Mount(app)

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
