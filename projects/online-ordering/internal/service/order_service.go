package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/vignesh/online-ordering/internal/models"
	repostiory "github.com/vignesh/online-ordering/internal/repository"
)

type OrderService struct {
	repo repostiory.OrderRepository
}

func NewOrderService(repo repostiory.OrderRepository) *OrderService {
	return &OrderService{
		repo: repo,
	}
}

func (o *OrderService) Create(ctx context.Context, order *models.Order) error {
	order.ID = uuid.New()

	order.Status = models.Created

	total := 0.0

	for i := range order.Items {

		order.Items[i].ID = uuid.New()

		total += float64(order.Items[i].Quantity) * order.Items[i].Price
	}

	order.TotalAmount = total

	return o.repo.Create(ctx, order)
}

func (o *OrderService) GetByID(ctx context.Context, id uuid.UUID) (*models.Order, error) {
	return o.repo.GetByID(ctx, id)
}
