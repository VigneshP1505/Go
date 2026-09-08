package repository

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/vignesh/online-ordering/internal/events"
	"github.com/vignesh/online-ordering/internal/models"
)

type orderRepository struct {
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

//CreateByOutbox implements [OrderRepository]

func (o *orderRepository) CreateOutbox(ctx context.Context, order *models.Order, event events.OrderCreatedEvent) error {
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

	payload, err := json.Marshal(event)

	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, `
		INSERT into outbox_events(id, event_type,aggregate_id,payload,published) values($1,$2,$3,$4, FALSE)
	`, uuid.New(), "OrderCreated", order.ID, payload)

	if err != nil {
		return err
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

func (o *orderRepository) GetUnpublishedEvents(ctx context.Context) ([]events.OrderCreatedEvent, error) {
	query := `Select id,event_type,payload from outbox_events where published=FALSE order by created_at limit 100`
	rows, err := o.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	unpublishedEvents := make([]events.OrderCreatedEvent, 0)

	for rows.Next() {
		var event events.OrderCreatedEvent
		err = rows.Scan(&event.OrderId, &event.CustomerId, &event.RestaurantId, &event.TotalAmount)
		unpublishedEvents = append(unpublishedEvents, event)
	}
	return unpublishedEvents, nil
}

func (o *orderRepository) MarkPublished(ctx context.Context, id uuid.UUID) error {
	query := `update outbox_events set published=TRUE, published_at=$1 where OrderId=$2`
	_, err := o.db.Exec(ctx, query, time.Now(), id)
	if err != nil {
		log.Println("Failed to update event as published")
		return err
	}
	return nil
}

func NewOrderRepository(db *pgxpool.Pool) OrderRepository {
	return &orderRepository{
		db: db,
	}
}
