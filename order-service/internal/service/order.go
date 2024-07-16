package service

import (
	"context"
	"errors"
	"order-service/internal/db/models"
	"order-service/internal/db/repository"
	"order-service/internal/dto"
	"order-service/internal/i18n"
	"order-service/internal/i18n/messages"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

var _ OrderService = (*Order)(nil)

func NewOrderService(orderRepository repository.OrderRepository) *Order {
	return &Order{
		orderRepository: orderRepository,
	}
}

type Order struct {
	orderRepository repository.OrderRepository
}

func (o *Order) CreateOrder(ctx *fiber.Ctx, request dto.OrderRequest) (*dto.CreateNewOrderResponse, int, error) {
	order := models.Order{
		UserID:       request.UserID,
		RestaurantID: request.RestaurantID,
		Status:       request.Status,
		TotalPrice:   request.TotalPrice,
		OrderItems:   []models.OrderItem{},
	}

	for _, item := range request.OrderItems {
		order.OrderItems = append(order.OrderItems, models.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Price,
		})
	}

	createOrder, err := o.orderRepository.CreateOrder(context.Background(), &order)
	if err != nil {
		return nil, fiber.StatusInternalServerError, errors.New(i18n.CreateMsg(ctx, messages.UnexpectedError))
	}

	response := &dto.CreateNewOrderResponse{
		ID:           createOrder.ID,
		UserID:       createOrder.UserID,
		RestaurantID: createOrder.RestaurantID,
		Status:       createOrder.Status,
		TotalPrice:   createOrder.TotalPrice,
		CreatedAt:    createOrder.CreatedAt,
		UpdatedAt:    createOrder.UpdatedAt,
		OrderItems:   []dto.OrderItemDTO{},
	}

	for _, item := range createOrder.OrderItems {
		response.OrderItems = append(response.OrderItems, dto.OrderItemDTO{
			ID:        item.ID,
			OrderID:   item.OrderID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Price,
		})
	}

	return response, fiber.StatusCreated, nil
}

func (o *Order) DeleteOrder(ctx *fiber.Ctx, id int) (int, error) {
	err := o.orderRepository.SoftDelete(id)
	if err != nil {
		log.Error(err)
		return fiber.StatusBadGateway, errors.New(i18n.CreateMsg(ctx, messages.UnexpectedError))
	}

	return fiber.StatusNoContent, nil
}

func (o *Order) ListOrders(ctx *fiber.Ctx) ([]dto.OrderDTO, int, error) {
	orders, err := o.orderRepository.ListOrders(context.Background())
	if err != nil {
		return nil, fiber.StatusInternalServerError, errors.New(i18n.CreateMsg(ctx, messages.UnexpectedError))
	}

	var response []dto.OrderDTO
	for _, order := range orders {
		orderDTO := dto.OrderDTO{
			ID:           order.ID,
			UserID:       order.UserID,
			RestaurantID: order.RestaurantID,
			Status:       order.Status,
			TotalPrice:   order.TotalPrice,
			OrderItems:   []dto.OrderItemDTO{},
		}
		for _, item := range order.OrderItems {
			orderDTO.OrderItems = append(orderDTO.OrderItems, dto.OrderItemDTO{
				ID:        item.ID,
				OrderID:   item.OrderID,
				ProductID: item.ProductID,
				Quantity:  item.Quantity,
				Price:     item.Price,
			})
		}
		response = append(response, orderDTO)
	}

	return response, fiber.StatusOK, nil
}

func (o *Order) UpdateOrder(ctx *fiber.Ctx, id int, request dto.UpdateOrderRequest) (*dto.UpdateOrderResponse, int, error) {
	panic("implement me")
}
