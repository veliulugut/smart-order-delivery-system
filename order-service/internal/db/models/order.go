package models

import "time"

type Order struct {
	ID           int         `gorm:"primaryKey;autoIncrement"`
	UserID       int         `gorm:"not null"`
	RestaurantID int         `gorm:"not null"`
	Status       string      `gorm:"type:varchar(255);not null"`
	TotalPrice   float64     `gorm:"type:decimal(10,2);not null"`
	CreatedAt    time.Time   `gorm:"autoCreateTime"`
	UpdatedAt    time.Time   `gorm:"autoUpdateTime"`
	OrderItems   []OrderItem `gorm:"foreignKey:OrderID"`
}

func (o *Order) TableName() string {
	return "public.orders"
}
