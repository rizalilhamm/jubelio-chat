package main

import (
	"github.com/gofiber/fiber/v2"
	"jubelio.com/chat/modules/chats/handlers"
)

func main() {
	// Create a new Fiber instance
	app := fiber.New()

	// Define a route handling GET requests to "/"
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	handler := handlers.New()

	// Mount routes from your HTTPHandler to the Fiber app
	handler.Mount(app)

	// Start the Fiber server on port 3000
	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
