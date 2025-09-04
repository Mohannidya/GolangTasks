package main

import (
	"MUTUALFUND_DEMO/database"
	"MUTUALFUND_DEMO/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.ConnectDB()
	routes.SetupRoutes(app)

	app.Listen(":3000")
}
