package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/vignesh/online-ordering/internal/events"
	"github.com/vignesh/online-ordering/internal/models"
)

type OrderRepository interface {
	Create(ctx context.Context, order *models.Order) error
	CreateOutbox(ctx context.Context, order *models.Order, event events.OrderCreatedEvent) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.Order, error)
	GetUnpublishedEvents(ctx context.Context) ([]events.OrderCreatedEvent, error)
	MarkPublished(ctx context.Context, id uuid.UUID) error
}

type CustomerRepository interface {
	Create(ctx context.Context, customer *models.Customer) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.Customer, error)
}
