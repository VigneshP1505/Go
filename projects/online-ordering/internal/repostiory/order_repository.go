package repostiory

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/vignesh/online-ordering/internal/models"
)

type orderRepository struct {
	db *sql.DB
}

// Create implements [OrderRepository].
func (o *orderRepository) Create(ctx context.Context, order *models.Order) error {
	tx, err := o.db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()
	_, err = tx.Exec(`
	INSERT into orders(id,customer_id,restaurant_id,total_amount,status)
	values($1,$2,$3,$4,$5)
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
		_, err := tx.Exec(`INSERT into order_items(id,order_id,item_name,quantity, price) values($1,$2,$3,$4,$5)`, item.ID, order.ID, item.ItemName, item.Quantity, item.Price)
		if err != nil {
			return err
		}
	}
	return tx.Commit()
}

// GetByID implements [OrderRepository].
func (o *orderRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Order, error) {
	query := `SELECT id,customer_id,restaurant_id,status,total_amount,createdAt from orders where id=$1`
	order := &models.Order{}
	err := o.db.QueryRowContext(ctx, query, id).Scan(&order.ID, &order.CustomerID, &order.RestaurantID, &order.Status, &order.TotalAmount, &order.CreatedAt)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func NewOrderRepository(db *sql.DB) OrderRepository {
	return &orderRepository{}
}
