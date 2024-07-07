package models

import "time"

type OrderItem struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	OrderID   int       `gorm:"not null"`
	ProductID int       `gorm:"not null"`
	Quantity  int       `gorm:"not null"`
	Price     float64   `gorm:"type:decimal(10,2);not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (o *OrderItem) TableName() string {
	return "public.order_items"
}
