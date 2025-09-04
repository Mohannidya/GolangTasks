package handlers

import (
	"MUTUALFUND_DEMO/database"
	"MUTUALFUND_DEMO/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func PlaceOrder(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string) // extracted from JWT middleware

	var order models.Order
	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	order.UserID = userID
	order.Status = "PLACED"
	order.PlacedAt = time.Now()

	func (udb *UserDb) CreateOrder(order *models.Order) (*models.Order, error) {
	_, err := udb.GetBy(order.UserID)
	if err != nil {
		return nil, errors.New("invalid user or user not found")
	}
	tx := udb.DB.Create(order)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return order, nil
}

	return c.JSON(order)
}

func GetOrders(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)

	var orders []models.Order
	if err := database.DB.Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not fetch orders"})
	}

	return c.JSON(orders)
}
