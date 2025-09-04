package handlers

import (
	"context"

	"github.com/Nerzal/gocloak/v13"
	"github.com/gofiber/fiber/v2"
)

var client = gocloak.NewClient("http://localhost:8080") // Keycloak URL

func Login(c *fiber.Ctx) error {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	token, err := client.Login(context.Background(), "myapi-client" ,"","myrealm", body.Username, body.Password)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}

	return c.JSON(token)
}
