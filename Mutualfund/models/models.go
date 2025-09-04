package models

import "time"

type User struct {
	ID    string `gorm:"primaryKey"`
	Email string `gorm:"unique"`
}

type Scheme struct {
	SchemeCode string `gorm:"primaryKey"`
	SchemeName string
}

type Order struct {
	ID          uint `gorm:"primaryKey"`
	UserID      string
	SchemeCode  string
	Side        string
	Amount      float64
	Status      string
	NavUsed     float64
	Units       float64
	ContractURL string
	PlacedAt    time.Time
	ConfirmedAt *time.Time
}

type Holding struct {
	UserID     string `gorm:"primaryKey"`
	SchemeCode string `gorm:"primaryKey"`
	Units      float64
}
