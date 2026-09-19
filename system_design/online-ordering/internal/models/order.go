package models

import (
	"time"

	"github.com/google/uuid"
)

type OrderStatus string

const (
	Created        OrderStatus = "CREATED"
	PaymentPending OrderStatus = "PAYMENT_PENDING"
)

type OrderItem struct {
	ID       uuid.UUID `json:"id"`
	ItemName string    `json:"item_name" validate:"required"`
	Quantity int       `json:"quantity" validate:"gte=1"`
	Price    float64   `json:"price" validate:"gt=0"`
}

type Order struct {
	ID           uuid.UUID   `json:"id"`
	CustomerID   uuid.UUID   `json:"customer_id" validate:"required"`
	RestaurantID uuid.UUID   `json:"resturant_id" validate:"required"`
	Items        []OrderItem `json:"items" validate:"required,dive"`
	TotalAmount  float64     `json:"total_amount"`
	Status       OrderStatus `json:"status"`
	CreatedAt    time.Time   `json:"created_at"`
}
