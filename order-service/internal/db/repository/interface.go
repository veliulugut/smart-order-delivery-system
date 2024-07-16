package repository

import (
	"context"
	"order-service/internal/db/models"
)

type OrderRepository interface {
	ListOrders(ctx context.Context) ([]*models.Order, error)
	CreateOrder(ctx context.Context, order *models.Order) (*models.Order, error)
	UpdateOrder(ctx context.Context, order *models.Order) error
	GetOrderById(id int) (*models.Order, error)
	SoftDelete(id int) error
}
