package domain

import (
	"time"
)

// Product is an entity that represents a product
type Product struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Stock     int64     `json:"stock"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
