package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/vignesh/online-ordering/internal/models"
	"github.com/vignesh/online-ordering/internal/repository"
)

type CustomerService struct {
	repo repository.CustomerRepository
}

func (c *CustomerService) Create(ctx context.Context, customer *models.Customer) error {
	customer.ID = uuid.New()
	err := c.repo.Create(ctx, customer)
	if err != nil {
		return err
	}
	return nil
}

func (c *CustomerService) GetByID(ctx context.Context, id uuid.UUID) error {
	return c.repo.GetByID(ctx, id)
}

func NewCustomerService(customerRepository repository.CustomerRepository) *CustomerService {
	return &CustomerService{
		repo: customerRepository,
	}
}
