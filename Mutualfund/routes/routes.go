package routes

import (
	"MUTUALFUND_DEMO/handlers"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/login", handlers.Login)

	api := app.Group("/api", jwtware.New(jwtware.Config{
		SigningKey: []byte("your-jwt-signing-key"),
		ContextKey: "user",
	}))

	api.Post("/orders", handlers.PlaceOrder)
	api.Get("/orders", handlers.GetOrders)
}
