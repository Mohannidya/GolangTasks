package models

import (
	"time"

	"gorm.io/datatypes"
)

type OrderEvent struct {
	ID        uint           `gorm:"primaryKey;autoIncrement;type:bigserial"`
	OrderID   string         `gorm:"not null;index"`            // FK to orders.order_id
	Event     string         `gorm:"type:varchar(50);not null"` // e.g., CREATED, PLACED, etc.
	Timestamp time.Time      `gorm:"autoCreateTime"`            // timestamptz
	Meta      datatypes.JSON `gorm:"type:jsonb;default:null"`   // optional JSON metadata
}

// Optionally, if you want to enforce foreign key constraint (Postgres example):
func (OrderEvent) TableName() string {
	return "order_events"
}
