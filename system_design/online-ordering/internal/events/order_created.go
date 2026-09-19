package events

import "github.com/google/uuid"

type OrderCreatedEvent struct {
	OrderId      uuid.UUID `json:"order_id"`
	CustomerId   uuid.UUID `json:"customer_id"`
	RestaurantId uuid.UUID `json:"restaurant_id"`
	TotalAmount  float64   `json:"total_amount"`
}
