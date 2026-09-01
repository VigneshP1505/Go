package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/vignesh/online-ordering/internal/models"
)

type orderRepository struct {
	db *pgxpool.Pool
}

type customerRepository struct {
	db *pgxpool.Pool
}

// Create implements [OrderRepository].
func (o *orderRepository) Create(
	ctx context.Context,
	order *models.Order,
) error {

	tx, err := o.db.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, `
		INSERT INTO orders (
			id,
			customer_id,
			restaurant_id,
			total_amount,
			order_status
		)
		VALUES ($1, $2, $3, $4, $5)
	`,
		order.ID,
		order.CustomerID,
		order.RestaurantID,
		order.TotalAmount,
		order.Status,
	)

	if err != nil {
		return err
	}

	for _, item := range order.Items {

		_, err = tx.Exec(ctx, `
			INSERT INTO order_items (
				id,
				order_id,
				item_name,
				quantity,
				price
			)
			VALUES ($1, $2, $3, $4, $5)
		`,
			item.ID,
			order.ID,
			item.ItemName,
			item.Quantity,
			item.Price,
		)

		if err != nil {
			return err
		}
	}

	return tx.Commit(ctx)
}

// GetByID implements [OrderRepository].
func (o *orderRepository) GetByID(
	ctx context.Context,
	id uuid.UUID,
) (*models.Order, error) {

	query := `
		SELECT
			id,
			customer_id,
			restaurant_id,
			order_status,
			total_amount,
			created_at
		FROM orders
		WHERE id = $1
	`

	order := &models.Order{}

	err := o.db.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&order.ID,
		&order.CustomerID,
		&order.RestaurantID,
		&order.Status,
		&order.TotalAmount,
		&order.CreatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, err
		}

		return nil, err
	}

	return order, nil
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

func NewOrderRepository(db *pgxpool.Pool) OrderRepository {
	return &orderRepository{
		db: db,
	}
}
