package database

import (
	"foodapp/user-service/models"

	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(order *models.Order) (*models.Order, error)
	Create(user *models.Order) (*models.Order, error)
}

type orderRepositoryImpl struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepositoryImpl{db}
}

func (repo *orderRepositoryImpl) Create(user *models.Order) (*models.Order, error) {
	// Assuming same logic as CreateOrder
	if err := repo.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
func (repo *orderRepositoryImpl) CreateOrder(order *models.Order) (*models.Order, error) {
	if err := repo.DB.Create(order).Error; err != nil {
		return nil, err
	}
	return order, nil
}
