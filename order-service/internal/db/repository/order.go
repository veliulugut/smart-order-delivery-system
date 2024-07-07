package repository

import (
	"order-service/internal/db/models"

	"gorm.io/gorm"
)

var _ OrderRepository = (*Order)(nil)

func New(db *gorm.DB) *Order {
	var order models.Order
	return &Order{
		dbClient:  db,
		tableName: order.TableName(),
	}
}

type Order struct {
	dbClient  *gorm.DB
	tableName string
}

func (o *Order) CreateOrder(order *models.Order) error {
	result := o.dbClient.Table(o.tableName).Create(order)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (o *Order) SoftDelete(id int) error {
	result := o.dbClient.Table(o.tableName).Where("id = ?", id).Update("is_active", false)
	if result.Error != nil {
		return result.Error
	}

	return nil

}

func (o *Order) GetOrderById(id int) (*models.Order, error) {
	var (
		order models.Order
	)

	result := o.dbClient.Table(o.tableName).Where("id = ?", id).First(&order)
	if result.Error != nil {
		return nil, result.Error
	}

	return &order, nil
}

func (o *Order) ListOrders() ([]*models.Order, error) {
	var orders []*models.Order

	result := o.dbClient.Table(o.tableName).Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}

	return orders, nil
}

func (o *Order) UpdateOrder(order *models.Order) error {
	result := o.dbClient.Table(o.tableName).Where("id = ?", order.ID).Updates(order)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
