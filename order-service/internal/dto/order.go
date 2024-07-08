package dto

import "time"

type OrderDTO struct {
	ID           int            `json:"id"`
	UserID       int            `json:"user_id"`
	RestaurantID int            `json:"restaurant_id"`
	Status       string         `json:"status"`
	TotalPrice   float64        `json:"total_price"`
	OrderItems   []OrderItemDTO `json:"order_items"`
}

type OrderRequest struct {
	UserID       int                `json:"user_id" binding:"required"`
	RestaurantID int                `json:"restaurant_id" binding:"required"`
	Status       string             `json:"status" binding:"required"`
	TotalPrice   float64            `json:"total_price" binding:"required"`
	OrderItems   []OrderItemRequest `json:"order_items" binding:"required"`
}

type CreateNewOrderResponse struct {
	ID           int            `json:"id"`
	UserID       int            `json:"user_id"`
	RestaurantID int            `json:"restaurant_id"`
	Status       string         `json:"status"`
	TotalPrice   float64        `json:"total_price"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	OrderItems   []OrderItemDTO `json:"order_items"`
}

type UpdateOrderRequest struct {
	UserID       *int               `json:"user_id"`
	RestaurantID *int               `json:"restaurant_id"`
	Status       *string            `json:"status"`
	TotalPrice   *float64           `json:"total_price"`
	OrderItems   *[]UpdateOrderItem `json:"order_items"`
}

type UpdateOrderResponse struct {
	ID           int            `json:"id"`
	UserID       int            `json:"user_id"`
	RestaurantID int            `json:"restaurant_id"`
	Status       string         `json:"status"`
	TotalPrice   float64        `json:"total_price"`
	UpdatedAt    time.Time      `json:"updated_at"`
	OrderItems   []OrderItemDTO `json:"order_items"`
}
