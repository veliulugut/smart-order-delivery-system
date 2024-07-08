package service

import (
	"order-service/internal/dto"

	"github.com/gofiber/fiber/v2"
)

type OrderService interface {
	CreateOrder(ctx *fiber.Ctx, request dto.OrderRequest) (*dto.CreateNewOrderResponse, int, error)
	DeleteOrder(ctx *fiber.Ctx, id int) (int, error)
	UpdateOrder(ctx *fiber.Ctx, id int, request dto.UpdateOrderRequest) (*dto.UpdateOrderResponse, int, error)
	ListOrders(ctx *fiber.Ctx) ([]dto.OrderDTO, int, error)
}
