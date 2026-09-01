package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/vignesh/online-ordering/internal/models"
)

type OrderRepository interface {
	Create(ctx context.Context, order *models.Order) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.Order, error)
}

type CustomerRepository interface {
	Create(ctx context.Context, customer *models.Customer) error
	GetByID(ctx context.Context, id uuid.UUID) error
}
