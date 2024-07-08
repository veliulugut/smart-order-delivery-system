package service

import (
	"order-service/internal/dto"

	"github.com/gofiber/fiber/v2"
)

var _ OrderService = (*Order)(nil)

type Order struct {
}

func (o *Order) CreateOrder(ctx *fiber.Ctx, request dto.OrderRequest) (*dto.CreateNewOrderResponse, int, error) {
	panic("unimplemented")
}

func (o *Order) DeleteOrder(ctx *fiber.Ctx, id int) (int, error) {
	panic("unimplemented")
}

func (o *Order) ListOrders(ctx *fiber.Ctx) ([]dto.OrderDTO, int, error) {
	panic("unimplemented")
}

func (o *Order) UpdateOrder(ctx *fiber.Ctx, id int, request dto.UpdateOrderRequest) (*dto.UpdateOrderResponse, int, error) {
	panic("unimplemented")
}
