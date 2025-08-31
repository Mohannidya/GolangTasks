package models

import (
	"time"
)

// Define status constants (optional but recommended)
const (
	StatusPlaced         = "PLACED"
	StatusPreparing      = "PREPARING"
	StatusCooking        = "COOKING"
	StatusOutForDelivery = "OUT_FOR_DELIVERY"
	StatusDelivered      = "DELIVERED"
)

type Order struct {
	ID           uint      `gorm:"primaryKey"`
	OrderID      string    `gorm:"uniqueIndex;not null"`
	CustomerName string    `gorm:"not null"`
	Address      string    `gorm:"not null"`
	Item         string    `gorm:"not null"`
	Size         string    `gorm:"not null"`
	Status       string    `gorm:"type:varchar(20);default:'PLACED';not null"`
	CreatedAt    int64     `json:"created" gorm:"index"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}
