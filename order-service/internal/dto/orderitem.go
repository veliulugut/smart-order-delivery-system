package dto

type OrderItemDTO struct {
	ID        int     `json:"id"`
	OrderID   int     `json:"order_id"`
	ProductID int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

type OrderItemRequest struct {
	ProductID int     `json:"product_id" binding:"required"`
	Quantity  int     `json:"quantity" binding:"required"`
	Price     float64 `json:"price" binding:"required"`
}

type UpdateOrderItem struct {
	Quantity *int     `json:"quantity"`
	Price    *float64 `json:"price"`
}
