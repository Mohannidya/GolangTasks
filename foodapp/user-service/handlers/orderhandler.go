package handlers

import (
	"foodapp/user-service/database"
	"foodapp/user-service/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	database.OrderRepository // promoted field
}
type IUserHandler interface {
	CreateUser(c *fiber.Ctx) error
	GetUserBy(c *fiber.Ctx) error
	GetUsersByLimit(c *fiber.Ctx) error
	CreateOrder(c *fiber.Ctx) error
}

func NewUserHandler(iuserdb database.OrderRepository) IUserHandler {
	return &UserHandler{OrderRepository}
}

func (uh *UserHandler) CreateUser(c *fiber.Ctx) error {
	user := new(models.Order)
	if err := c.BodyParser(user); err != nil {
		return err
	}

	user.Status = "active"
	user.CreatedAt = time.Now().Unix()

	createdUser, err := uh.Create(user)
	if err != nil {
		return err
	}

	return c.JSON(createdUser)

}
