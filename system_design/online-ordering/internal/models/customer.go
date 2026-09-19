package models

import "github.com/google/uuid"

type Customer struct {
	ID            uuid.UUID `json:"id,omitempty"`
	CustomerName  string    `json:"customer_name" validate:"required"`
	CustomerEmail string    `json:"customer_email" validate:"required,email"`
	Unit          string    `json:"unit" validate:"required"`
	Street        string    `json:"street" validate:"required"`
	City          string    `json:"city" validate:"required"`
	PostalCode    string    `json:"postal_code" validate:"required"`
}
