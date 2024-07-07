package repository

import "order-service/internal/db/models"

type OrderRepository interface {
	ListOrders() ([]*models.Order, error)
	CreateOrder(order *models.Order) error
	UpdateOrder(order *models.Order) error
	GetOrderById(id int) (*models.Order, error)
	SoftDelete(id int) error
}
