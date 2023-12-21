package middlewares

import (
	"encoding/base64"
	"strings"
	"jubelio.com/chat/config"

	"github.com/gofiber/fiber/v2"
)

func VerifyBasicAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		auth := c.Get(fiber.HeaderAuthorization)
		if auth == "" {
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
		}

		authParts := strings.SplitN(auth, " ", 2)
		if len(authParts) != 2 || authParts[0] != "Basic" {
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
		}

		payload, err := base64.StdEncoding.DecodeString(authParts[1])
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
		}

		credentials := strings.SplitN(string(payload), ":", 2)
		if len(credentials) != 2 {
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
		}

		username := credentials[0]
		password := credentials[1]

		if username != config.GlobalEnv.BasicAuthUsername || password != config.GlobalEnv.BasicAuthPassword {
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
		}

		return c.Next()
	}
}
