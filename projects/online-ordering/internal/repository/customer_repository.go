package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vignesh/online-ordering/internal/models"
)

type customerRepository struct {
	db *pgxpool.Pool
}

func (c *customerRepository) Create(ctx context.Context, customer models.Customer) error {
	tx, err := c.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	_, err = tx.Exec(ctx, `
	INSERT into customers(
	id,
	customer_name,
	customer_email,
	unit,
	street,
	city,
	postal
	)
	`, customer.ID, customer.CustomerName, customer.CustomerEmail, customer.Unit, customer.Street, customer.City, customer.PostalCode)

	if err != nil {
		return nil
	}

	return tx.Commit(ctx)
}

func (c *customerRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Customer, error) {
	query := `Select * from customers where id=$1`
	customer := &models.Customer{}
	err := c.db.QueryRow(ctx, query).Scan(
		&customer.ID,
		&customer.CustomerName,
		&customer.CustomerEmail,
		&customer.City,
		&customer.CustomerEmail,
		&customer.Unit,
		&customer.Street,
		&customer.City,
		&customer.PostalCode,
	)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func NewCustomerRepository(db *pgxpool.Pool) OrderRepository {
	return &orderRepository{
		db: db,
	}
}
